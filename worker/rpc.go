package worker

import (
	"context"

	"github.com/ubiqueworks/joat/internal/rpc"
)

func newRpcServer(worker *worker) *rpcServer {
	return &rpcServer{
		worker: worker,
	}
}

type rpcServer struct {
	worker *worker
}

func (s *rpcServer) Join(context.Context, *rpc.JoinRequest) (*rpc.JoinResponse, error) {
	return nil, nil
}

func (s *rpcServer) Leave(context.Context, *rpc.LeaveRequest) (*rpc.LeaveResponse, error) {
	return nil, nil
}
