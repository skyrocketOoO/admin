package Middleware

import (
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

func (it *Handler) UnBindStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (response *Connect.InnerResponse, err error) {
	reqData, err := Bytes.ProtoDecode[V0.UnBindStoreRoleRequest](req.Msg.Data)
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
			return
		}
		tgRole := StoreRole.StoreRole{}
		if err = dbSet.Core.Take(&tgRole, "ID = ?", tgAccount.Store.Role.ID).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		if tgRole.Name == StoreRole.WebMasterName {
			err = Error.New("cannot Unbind webmaster account")
			err = Error.Throw(err)
			return
		}

		minAuthRole := StoreRole.StoreRole{}
		if err = dbSet.Core.
			Where("Name = ?", StoreRole.MinimumAuthorityName).
			Where("StoreID = ?", operAccount.EntryID).
			Take(&minAuthRole).Error; err != nil {
			err = Error.Throw(err)
			return
		}

		tgAccount.Store.Role.ID = minAuthRole.ID
		if err = dbSet.Core.Save(&tgAccount).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		return nil
	}); err != nil {
		err = Error.Throw(err)
		return
	}

	response, err = Connect.NewProtoResponse(&V0.UnBindStoreRoleResponse{})
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
