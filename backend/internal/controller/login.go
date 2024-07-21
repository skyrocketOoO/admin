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
	"github.com/skyrocketOoO/GoUtils/Struct"
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

	resp = &api.LoginResponse{
		SessionID: uuid.UUID(s.SessionSvc.GetSession(account.ID)).String(),
	}
	if err = Struct.Scan(account.Role, &resp.Role); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return resp, nil
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
