package controller

import (
	"admin/internal/service/Session"
	"admin/proto/protoconnect"
)

type (
	Server struct {
		protoconnect.UnimplementedMainServiceHandler
		SessionSvc *Session.SessionSvc
	}
)

func NewServer(sessionSvc *Session.SessionSvc) (sv *Server) {
	return &Server{
		SessionSvc: sessionSvc,
	}
}
