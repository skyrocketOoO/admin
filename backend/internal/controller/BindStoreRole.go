package Middleware

import (
	"errors"

	"gorm.io/gorm"
	"taihe.asia/TeaSwap/api/V0"
	"taihe.asia/TeaSwap/internal/Common/Bytes"
	"taihe.asia/TeaSwap/internal/Common/Error"
	"taihe.asia/TeaSwap/internal/Common/PrepareBeforeAuth"
	"taihe.asia/TeaSwap/internal/Common/Type"
	"taihe.asia/TeaSwap/internal/Connect"
	"taihe.asia/TeaSwap/internal/model/entity/Core/Account"
	"taihe.asia/TeaSwap/internal/model/entity/Core/StoreRole"
)

func (it *Handler) BindStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (response *Connect.InnerResponse, err error) {
	reqData, err := Bytes.ProtoDecode[V0.BindStoreRoleRequest](req.Msg.Data)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	dbSet, operAccount, operRole, err := PrepareBeforeAuth.StoreGet(context)
	if nil != err {
		err = Error.Throw(err)
		return
	}
	if !operRole.Page.Role.Update {
		err = Error.New("not permitted")
		err = Error.Throw(err)
		return
	}

	if err = dbSet.Core.Transaction(func(tx *gorm.DB) (err error) {
		tgAccount := Account.Account{}
		if err = dbSet.Core.Take(&tgAccount, "ID = ?", reqData.AccountID).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		if tgAccount.EntryType != Account.Store || tgAccount.EntryID != operAccount.EntryID {
			err = Error.New("not permitted")
			err = Error.Throw(err)
			return
		}

		oldRole := StoreRole.StoreRole{}
		if err = dbSet.Core.Take(&oldRole, "ID = ?", tgAccount.Store.Role.ID).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		if oldRole.Name == StoreRole.WebMasterName {
			err = Error.Throw(errors.New("cannot update webmaster account"))
			return
		}

		tgRole := StoreRole.StoreRole{}
		if err = dbSet.Core.Take(&tgRole, "ID = ?", reqData.RoleID).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		if tgRole.StoreID != operAccount.EntryID {
			err = Error.New("not permitted")
			err = Error.Throw(err)
			return
		}
		if tgRole.Name == StoreRole.WebMasterName {
			err = Error.New("cannot bind webmaster role")
			err = Error.Throw(err)
			return
		}

		tgAccount.Store.Role.ID = tgRole.ID
		if err = dbSet.Core.Save(&tgAccount).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		return nil
	}); err != nil {
		err = Error.Throw(err)
		return
	}

	response, err = Connect.NewProtoResponse(&V0.BindStoreRoleResponse{})
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
