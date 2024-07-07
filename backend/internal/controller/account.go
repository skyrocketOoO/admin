package controller

import (
	"context"
	"fmt"

	"admin/api"
	"admin/internal/model"
	"admin/internal/service/orm"
	"admin/internal/utils/dao"
	"admin/internal/utils/password"

	"admin/internal/domain/Error"

	"github.com/skyrocketOoO/GoUtils/Struct"
)

func (s *Server) CreateAccount(
	ctx context.Context,
	req *api.CreateAccountRequest,
) (resp *api.Empty, err error) {
	salt, err := password.GenSalt()
	if err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	if err = orm.GetDb().Create(&model.Account{
		UserName:    req.UserName,
		HashPass:    password.Hash(req.Password, salt),
		Salt:        salt,
		DisplayName: req.DisplayName,
	}).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *Server) ListAccount(
	ctx context.Context,
	req *api.ListAccountRequest,
) (resp *api.ListAccountResponse, err error) {
	option := Struct.DeepNew[dao.ListOption]()
	if req.GetOption() != nil {
		if err = Struct.Scan(req.Option, option); err != nil {
			return nil, Error.Internal.WithTrace(err)
		}
	}

	db := orm.GetDb().Model(&model.Account{})
	resp = Struct.DeepNew[api.ListAccountResponse]()
	accounts := []model.Account{}
	if resp.Total, err = dao.ListWithPager(db, *option, &accounts); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp.List = make([]*api.AccountData, len(accounts))
	for i, account := range accounts {
		resp.List[i] = Struct.DeepNew[api.AccountData]()
		if err = Struct.Scan(&account, resp.List[i]); err != nil {
			return nil, Error.Internal.WithTrace(err)
		}
	}

	return
}

func (s *Server) UpdateAccount(
	ctx context.Context,
	req *api.UpdateAccountRequest,
) (resp *api.Empty, err error) {
	tx := orm.GetDb().
		Model(&model.Account{}).
		Where("ID = ?", req.GetID()).
		Updates(model.Account{
			DisplayName: req.GetDisplayName(),
			State:       req.GetState(),
		})
	if tx.Error != nil {
		return nil, Error.Internal.WithTrace(tx.Error)
	}
	if tx.RowsAffected == 0 {
		return nil, Error.NotFound.WithTrace(
			fmt.Errorf("record not found: ID: %s", req.GetID()))
	}

	return
}

func (s *Server) DeleteAccount(
	ctx context.Context,
	req *api.DeleteAccountRequest,
) (resp *api.Empty, err error) {
	if err = orm.GetDb().Delete(&model.Account{}, req.GetID()).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	return
}
