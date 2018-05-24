package worker

import (
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	_ "github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/internal/cluster"
	"github.com/ubiqueworks/joat/internal/rpc"
	"google.golang.org/grpc"
)

type Config struct {
	AdvertiseIP *net.IP
	ClusterAddr *net.TCPAddr
	DataDir     string
	VerboseMode bool
	JoinAddr    *net.TCPAddr
	RpcAddr     *net.TCPAddr
}

func StartWorker(conf *Config) error {
	worker := &worker{
		config:     conf,
		shutdownCh: make(chan struct{}, 1),
	}
	return worker.start()
}

type worker struct {
	config       *Config
	debugMode    bool
	shutdownLock sync.Mutex
	shutdownCh   chan struct{}
	shutdown     bool
}

func (w *worker) start() error {
	errCh := make(chan error)
	shutdownCh := w.shutdownCh

	var wg sync.WaitGroup
	wg.Add(2)

	go w.startRpcServer(&wg, shutdownCh, errCh)
	go w.startClusterManager(&wg, shutdownCh, errCh)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errCh:
		log.Error().Err(err).Msg("caught service error")
		w.stop()
	case <-quit:
		w.stop()
	}

	log.Debug().Msg("waiting for shutdown to complete...")
	wg.Wait()
	log.Debug().Msg("shutdown completed")

	return nil
}

func (w *worker) stop() {
	w.shutdownLock.Lock()
	defer w.shutdownLock.Unlock()

	if w.shutdown {
		return
	}
	w.shutdown = true
	close(w.shutdownCh)
}

func (w *worker) startRpcServer(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	grpcServer := grpc.NewServer()
	rpc.RegisterWorkerServiceServer(grpcServer, newRpcServer(w))

	startFunc := func() {
		serverAddr := w.config.RpcAddr.String()
		log.Info().Msgf("Starting RPC server on %v", serverAddr)
		listener, err := net.Listen("tcp", serverAddr)
		if err != nil {
			startErrCh <- err
			return
		}
		grpcServer.Serve(listener)
	}

	stopFunc := func() {
		log.Debug().Msg("Stopping RPC server...")
		grpcServer.GracefulStop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		close(doneCh)
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	log.Info().Msg("RPC server stopped")
}

func (w *worker) startClusterManager(wg *sync.WaitGroup, shutdownCh <-chan struct{}, errCh chan<- error) {
	defer wg.Done()

	startErrCh := make(chan error)
	doneCh := make(chan struct{}, 1)

	manager := cluster.NewClusterManager(&cluster.Config{
		AdvertiseIP:   w.config.AdvertiseIP,
		ClusterAddr:   w.config.ClusterAddr,
		Role:          cluster.RoleWorker,
		DataDir:       w.config.DataDir,
		EnableSerfLog: w.config.VerboseMode,
		RpcPort:       w.config.RpcAddr.Port,
	})

	startFunc := func() {
		serverAddr := w.config.ClusterAddr.String()
		log.Info().Msgf("Starting cluster manager on %v", serverAddr)
		if err := manager.Start(); err != nil {
			startErrCh <- err
		}

		if err := manager.Join(w.config.JoinAddr); err != nil {
			startErrCh <- err
		}
	}

	stopFunc := func() {
		log.Debug().Msg("Stopping cluster manager...")
		manager.Stop()
		close(doneCh)
	}

	go startFunc()

	select {
	case err := <-startErrCh:
		errCh <- err
		stopFunc()
	case <-shutdownCh:
		stopFunc()
	}

	<-doneCh
	log.Info().Msg("Cluster manager stopped")
}
