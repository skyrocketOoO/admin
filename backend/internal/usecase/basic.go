package usecase

import (
	"context"

	"web-server-template/internal/service/dao"
)

type BasicUsecase struct {
	dao *dao.Dao
}

func NewBasicUsecase(dao *dao.Dao) *BasicUsecase {
	return &BasicUsecase{
		dao: dao,
	}
}

func (u *BasicUsecase) Healthy(c context.Context) error {
	// do something check like db connection is established
	if err := u.dao.Ping(c); err != nil {
		return err
	}

	return nil
}
