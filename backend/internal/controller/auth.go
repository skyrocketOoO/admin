package controller

import (
	"context"
	"fmt"

	"admin/gen/column"
	"admin/internal/domain/Error"
	"admin/internal/model"
	"admin/internal/service/Session"
	"admin/internal/service/orm"
	"admin/internal/utils/password"
	"admin/proto"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/skyrocketOoO/GoUtils/Struct"
	"github.com/skyrocketOoO/gorm-enhance-plugin/operator"
	"github.com/skyrocketOoO/gorm-enhance-plugin/query"
	"gorm.io/gorm"
)

func (s *Server) Login(
	ctx context.Context,
	connReq *connect.Request[proto.LoginReq],
) (connResp *connect.Response[proto.LoginResp], err error) {
	req := connReq.Msg
	account := model.Account{}
	if err = orm.GetDb().
		Where(query.Build(column.Accounts.UserName, operator.Equal), req.GetUserName()).
		Take(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, Error.NotFound.WithTrace(err)
		}
		return nil, Error.Internal.WithTrace(err)
	}

	if !password.Equal(req.GetPassword(), account.Salt, account.HashPass) {
		return nil, Error.Unauthenticated.WithTrace(fmt.Errorf("invalid password"))
	}

	resp := Struct.DeepNew[proto.LoginResp]()
	resp.SessionID = uuid.UUID(s.SessionSvc.GetSession(account.ID)).String()

	if err = Struct.Scan(account.Role, resp.Role); err != nil {
		return nil, Error.Internal.WithTrace(err)
	}

	return connect.NewResponse(resp), nil
}

func (s *Server) Logout(
	ctx context.Context,
	connReq *connect.Request[proto.LogoutReq],
) (connResp *connect.Response[proto.Empty], err error) {
	req := connReq.Msg
	sessionID, err := uuid.Parse(req.GetSessionID())
	if err != nil {
		return nil, Error.InvalidArgument.WithTrace(err)
	}

	s.SessionSvc.DeleteSession(Session.SessionID(sessionID))
	return
}
