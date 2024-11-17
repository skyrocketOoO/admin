package rpc

import (
	"admin/internal/service/Session"
	"admin/proto/protoconnect"
)

type MainServer struct {
	protoconnect.UnimplementedMainServiceHandler
	SessionSvc *Session.SessionSvc
}

func NewMainServer(sessionSvc *Session.SessionSvc) (sv *MainServer) {
	return &MainServer{
		SessionSvc: sessionSvc,
	}
}

type SideProjectServer struct {
	protoconnect.UnimplementedSideProjectServiceHandler
}

func NewSideProjectServer() (sv *SideProjectServer) {
	return &SideProjectServer{}
}
