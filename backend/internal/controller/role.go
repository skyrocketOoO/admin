package controller

import (
	"context"
	"fmt"

	"admin/api"
	"admin/internal/model"
	"admin/internal/service/orm"
	"admin/internal/utils/dao"

	"admin/internal/domain/Error"

	"github.com/skyrocketOoO/GoUtils/Struct"
	"gorm.io/gorm"
)

func (s *Server) CreateRole(
	ctx context.Context,
	req *api.CreateRoleRequest,
) (resp *api.Empty, err error) {
	role := Struct.DeepNew[model.Role]()
	if err = Struct.Scan(req, req); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	if err = orm.GetDb().Create(role).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *Server) ListRole(
	ctx context.Context,
	req *api.ListRoleRequest,
) (resp *api.ListRoleResponse, err error) {
	option := Struct.DeepNew[dao.ListOption]()
	if err = Struct.Scan(req.Option, option); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	db := orm.GetDb().Model(&model.Role{}).Select("ID", "Name")
	resp = Struct.DeepNew[api.ListRoleResponse]()
	Roles := []model.Role{}
	if resp.Total, err = dao.ListWithPager(db, *option, &Roles); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp.List = make([]*api.RoleData, len(Roles))
	for i, Role := range Roles {
		if err = Struct.Scan(&Role, resp.List[i]); err != nil {
			return nil, Error.Internal.WithTrace(err)
		}
	}

	return
}

func (s *Server) GetRoleAuth(
	ctx context.Context,
	req *api.GetRoleAuthRequest,
) (resp *api.GetRoleAuthResponse, err error) {
	role := model.Role{}
	if err = orm.GetDb().Where("ID =?", req.GetID()).Take(&role).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp = Struct.DeepNew[api.GetRoleAuthResponse]()
	if err = Struct.Scan(role.Page, resp.Page); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return
}

func (s *Server) UpdateRole(
	ctx context.Context,
	req *api.UpdateRoleRequest,
) (resp *api.Empty, err error) {
	update := Struct.DeepNew[model.Role]()
	if err = Struct.Scan(req, update); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	tx := orm.GetDb().
		Model(&model.Role{}).
		Where("ID = ?", req.GetID()).
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

func (s *Server) DeleteRole(
	ctx context.Context,
	req *api.DeleteRoleRequest,
) (resp *api.Empty, err error) {
	if err = orm.GetDb().Delete(&model.Role{}, req.GetID()).Error; err != nil {
		return nil, Error.Internal.WithTrace(err)
	}
	return
}

func (s *Server) BindRole(
	ctx context.Context,
	req *api.BindRoleRequest,
) (resp *api.Empty, err error) {
	db := orm.GetDb()

	if err = db.Transaction(func(tx *gorm.DB) error {
		role := model.Role{}
		if err = tx.Take(&role, "ID = ?", req.GetRoleID()).Error; err != nil {
			return Error.Internal.WithTrace(tx.Error)
		}

		tx = tx.Model(&model.Account{}).
			Where("ID = ?", req.GetAccountID()).
			Update("RoleID", req.GetRoleID())
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

func (s *Server) BindRole(
	ctx context.Context,
	req *api.BindRoleRequest,
) (resp *api.Empty, err error) {
	db := orm.GetDb()

	if err = db.Transaction(func(tx *gorm.DB) error {
		role := model.Role{}
		if err = tx.Take(&role, "ID = ?", req.GetRoleID()).Error; err != nil {
			return Error.Internal.WithTrace(tx.Error)
		}

		tx = tx.Model(&model.Account{}).
			Where("ID = ?", req.GetAccountID()).
			Update("RoleID", req.GetRoleID())
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
