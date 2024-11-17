package rpc

import (
	"context"
	"fmt"

	"admin/gen/column"
	"admin/internal/model"
	"admin/internal/service/orm"
	"admin/internal/utils/dao"
	"admin/proto"

	"admin/internal/domain/Error"

	"connectrpc.com/connect"
	"github.com/skyrocketOoO/GoUtils/Struct"
	"github.com/skyrocketOoO/gorm-enhance-plugin/operator"
	"github.com/skyrocketOoO/gorm-enhance-plugin/query"
	"gorm.io/gorm"
)

func (s *MainServer) CreateRole(
	ctx context.Context,
	connReq *connect.Request[proto.CreateRoleReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	role := Struct.DeepNew[model.Role]()
	if err = Struct.Scan(req, req); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	if err = orm.GetDb().Create(role).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *MainServer) ListRole(
	ctx context.Context,
	connReq *connect.Request[proto.ListRoleReq],
) (connResp *connect.Response[proto.ListRoleResp], err error) {
	req := connReq.Msg

	db := orm.GetDb().Model(&model.Role{}).Select(column.Roles.ID, column.Roles.Name)
	resp := Struct.DeepNew[proto.ListRoleResp]()
	Roles := []model.Role{}
	if resp.Total, err = dao.ListWithPager(db, req.GetOption(), &Roles); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp.List = make([]*proto.RoleData, len(Roles))
	for i, Role := range Roles {
		resp.List[i] = Struct.DeepNew[proto.RoleData]()
		if err = Struct.Scan(&Role, resp.List[i]); err != nil {
			return nil, Error.Internal.WithTrace(err)
		}
	}

	return connect.NewResponse(resp), nil
}

func (s *MainServer) GetRoleAuth(
	ctx context.Context,
	connReq *connect.Request[proto.GetRoleAuthReq],
) (connResp *connect.Response[proto.GetRoleAuthResp], err error) {
	req := connReq.Msg
	role := model.Role{}
	if err = orm.GetDb().
		Where(query.Build(column.Roles.ID, operator.Equal), req.GetID()).
		Take(&role).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp := Struct.DeepNew[proto.GetRoleAuthResp]()
	if err = Struct.Scan(role.Page, resp.Page); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return connect.NewResponse(resp), nil
}

func (s *MainServer) UpdateRole(
	ctx context.Context,
	connReq *connect.Request[proto.UpdateRoleReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	update := Struct.DeepNew[model.Role]()
	if err = Struct.Scan(req, update); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	tx := orm.GetDb().
		Model(&model.Role{}).
		Where(query.Build(column.Roles.ID, operator.Equal), req.GetID()).
		Updates(update)
	if tx.Error != nil {
		return nil, Error.Internal.WithTrace(tx.Error)
	}
	if tx.RowsAffected == 0 {
		return nil, Error.NotFound.WithTrace(
			fmt.Errorf("record not found: ID: %s", req.GetID()))
	}

	return
}

func (s *MainServer) DeleteRole(
	ctx context.Context,
	connReq *connect.Request[proto.DeleteRoleReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	if err = orm.GetDb().Delete(&model.Role{}, req.GetID()).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	return
}

func (s *MainServer) BindRole(
	ctx context.Context,
	connReq *connect.Request[proto.BindRoleReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	db := orm.GetDb()

	if err = db.Transaction(func(tx *gorm.DB) error {
		role := model.Role{}
		if err = tx.Take(
			&role,
			query.Build(column.Roles.ID, operator.Equal),
			req.GetRoleID(),
		).Error; err != nil {
			return Error.Internal.WithTrace(tx.Error)
		}

		tx = tx.Model(&model.Account{}).
			Where(query.Build(column.Accounts.ID, operator.Equal), req.GetAccountID()).
			Update(column.Accounts.RoleID, req.GetRoleID())
		if tx.Error != nil {
			return Error.Internal.WithTrace(tx.Error)
		}
		if tx.RowsAffected == 0 {
			return Error.NotFound.WithTrace(
				fmt.Errorf("record not found: AccountID: %s", req.GetAccountID()))
		}

		return nil
	}); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *MainServer) UnBindRole(
	ctx context.Context,
	connReq *connect.Request[proto.UnBindRoleReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	db := orm.GetDb()

	if err = db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&model.Account{}).
			Where(query.Build(column.Accounts.ID, operator.Equal), req.GetAccountID()).
			Update(column.Accounts.RoleID, 0)
		if tx.Error != nil {
			return Error.Internal.WithTrace(tx.Error)
		}
		if tx.RowsAffected == 0 {
			return Error.NotFound.WithTrace(
				fmt.Errorf("record not found: AccountID: %s", req.GetAccountID()))
		}

		return nil
	}); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}
