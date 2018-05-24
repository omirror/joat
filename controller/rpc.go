package controller

import (
	"context"

	"github.com/ubiqueworks/joat/internal/rpc"
)

func newRpcServer(controller *controller) *rpcServer {
	return &rpcServer{
		controller: controller,
	}
}

type rpcServer struct {
	controller *controller
}

func (s *rpcServer) Join(context.Context, *rpc.JoinRequest) (*rpc.JoinResponse, error) {
	return nil, nil
}

func (s *rpcServer) Leave(context.Context, *rpc.LeaveRequest) (*rpc.LeaveResponse, error) {
	return nil, nil
}
