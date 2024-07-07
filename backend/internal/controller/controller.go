package controller

import (
	"admin/api"
	"admin/internal/service/Session"
)

type (
	Server struct {
		api.UnimplementedMainServer
		SessionSvc *Session.SessionSvc
	}
)

func NewServer(sessionSvc *Session.SessionSvc) (sv *Server) {
	return &Server{
		SessionSvc: sessionSvc,
	}
}
