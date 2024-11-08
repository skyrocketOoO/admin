package controller

import (
	"context"
	"fmt"

	"admin/internal/model"
	"admin/internal/service/orm"
	"admin/internal/utils/dao"
	"admin/internal/utils/password"
	"admin/proto"

	"connectrpc.com/connect"

	"admin/internal/domain/Error"

	"github.com/skyrocketOoO/GoUtils/Struct"
)

func (s *Server) CreateAccount(
	ctx context.Context,
	connReq *connect.Request[proto.CreateAccountReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	salt, err := password.GenSalt()
	if err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	tarAccount := &model.Account{
		UserName:    req.GetUserName(),
		HashPass:    password.Hash(req.GetPassword(), salt),
		Salt:        salt,
		DisplayName: req.GetDisplayName(),
	}

	if req.GetRoleID() != 0 {
		tarAccount.RoleID = uint(req.GetRoleID())
	}

	if err = orm.GetDb().Create(tarAccount).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *Server) ListAccount(
	ctx context.Context,
	connReq *connect.Request[proto.ListAccountReq],
) (connResp *connect.Response[proto.ListAccountResp], err error) {
	req := connReq.Msg

	db := orm.GetDb().Model(&model.Account{})
	resp := Struct.DeepNew[proto.ListAccountResp]()
	accounts := []model.Account{}
	if resp.Total, err = dao.ListWithPager(db, req.GetOption(), &accounts); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp.List = make([]*proto.Account, len(accounts))
	for i, account := range accounts {
		resp.List[i] = Struct.DeepNew[proto.Account]()
		resp.List[i].ID = int32(account.ID)
		if err = Struct.Scan(&account, resp.List[i]); err != nil {
			return nil, Error.Internal.WithTrace(err)
		}
	}

	return connect.NewResponse(resp), nil
}

func (s *Server) UpdateAccount(
	ctx context.Context,
	connReq *connect.Request[proto.UpdateAccountReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
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
	connReq *connect.Request[proto.DeleteAccountReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	if err = orm.GetDb().Delete(&model.Account{}, req.GetID()).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	return
}
