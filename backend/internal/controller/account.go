package controller

import (
	"context"

	"admin/api"
	"admin/internal/service/orm"
	"admin/internal/service/orm/model"
	"admin/internal/utils/dao"
	"admin/internal/utils/password"

	"admin/internal/domain/Error"

	"github.com/skyrocketOoO/GoUtils/Struct"
)

type Server struct {
	api.UnsafeAdminServer
}

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
	db := orm.GetDb()
	if req.GetOption() != nil {
		opt := req.GetOption()
		if opt.GetSorter() != nil {
			expr := opt.GetSorter().GetField()
			if !opt.GetSorter().GetAsc() {
				expr += " desc"
			}
			db = db.Order(expr)
		}
		if opt.GetQuery() != nil {
			for _, q := range opt.GetQuery() {
				if q.GetFuzzy() {
					db = db.Where(q.GetField()+" LIKE ?", "%"+q.GetValue()+"%")
				} else {
					db = db.Where(q.GetField()+" = ?", q.GetValue())
				}
			}
		}
	}

	resp = Struct.DeepNew[api.ListAccountResponse]()
	accounts := []model.Account{}
	if resp.Total, err = dao.ListWithPager(
		db,
		dao.Pager{
			Number: int(req.GetOption().GetPager().GetNumber()),
			Size:   int(req.GetOption().GetPager().GetSize()),
		},
		&accounts,
	); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	resp.List = make([]*api.AccountData, len(accounts))
	for i, account := range accounts {
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
	return
}

func (s *Server) DeleteAccount(
	ctx context.Context,
	req *api.DeleteAccountRequest,
) (resp *api.Empty, err error) {
	return
}
