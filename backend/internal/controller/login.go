package controller

import (
	"context"
	"fmt"

	"admin/api"
	"admin/internal/domain/Error"
	"admin/internal/model"
	"admin/internal/service/Session"
	"admin/internal/service/orm"
	"admin/internal/utils/password"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Server) Login(
	ctx context.Context,
	req *api.LoginRequest,
) (resp *api.LoginResponse, err error) {
	account := model.Account{}
	if err = orm.GetDb().Where("UserName = ?", req.GetUserName()).
		Take(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, Error.NotFound.WithTrace(err)
		}
		return nil, Error.Internal.WithTrace(err)
	}

	if !password.Equal(req.GetPassword(), account.Salt, account.HashPass) {
		return nil, Error.Unauthenticated.WithTrace(fmt.Errorf("invalid password"))
	}

	return &api.LoginResponse{
		SessionID: uuid.UUID(s.SessionSvc.GetSession(account.ID)).String(),
	}, nil
}

func (s *Server) Logout(
	ctx context.Context,
	req *api.LogoutRequest,
) (resp *api.Empty, err error) {
	sessionID, err := uuid.Parse(req.GetSessionID())
	if err != nil {
		return nil, Error.InvalidArgument.WithTrace(err)
	}

	s.SessionSvc.DeleteSession(Session.SessionID(sessionID))
	return
}
