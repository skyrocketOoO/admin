package controller

import (
	"context"

	"admin/api"
	"admin/internal/service/orm"
	"admin/internal/service/orm/model"
	"admin/internal/utils/password"

	"admin/internal/domain/Error"
)

type AccountServer struct {
	api.UnimplementedAccountServiceServer
}

func (s *AccountServer) CreateAccount(
	ctx context.Context,
	req *api.CreateAccountRequest,
) (empty *api.Empty, err error) {
	salt, err := password.GenSalt()
	if err != nil {
		return nil, Error.Internal.LogWithTrace(err)
	}

	if err = orm.GetDb().Create(&model.Account{
		UserName:    req.UserName,
		HashPass:    password.Hash(req.Password, salt),
		DisplayName: req.DisplayName,
	}).Error; err != nil {
		return nil, Error.Internal.LogWithTrace(err)
	}

	return
}
