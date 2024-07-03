package controller

import (
	"context"

	"admin/api"
	"admin/internal/service/orm"
	"admin/internal/service/orm/model"
	"admin/internal/utils/password"

	"admin/internal/domain/Error"

	"github.com/golang/protobuf/ptypes/empty"
)

type AccountServer struct {
	api.UnimplementedAccountServer
}

func (s *AccountServer) CreateAccount(
	ctx context.Context,
	req *api.CreateAccountRequest,
) (*empty.Empty, error) {
	salt, err := password.GenSalt()
	if err != nil {
		return nil, Error.Internal.LogWithTrace(err)
	}

	orm.GetDb().Create(&model.Account{
		UserName:    in.UserName,
		Password:    in.Password,
		DisplayName: in.DisplayName,
	})
	return nil
	return nil, nil
}
