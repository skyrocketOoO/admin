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

func (it *Handler) DeleteStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (resp *Connect.InnerResponse, err error) {
	reqData, err := Bytes.ProtoDecode[V0.DeleteStoreRoleRequest](req.Msg.Data)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	dbSet, _, operRole, err := PrepareBeforeAuth.StoreGet(context)
	if nil != err {
		err = Error.Throw(err)
		return
	}
	if !operRole.Page.Role.Delete {
		err = Error.New("not permitted")
		return
	}

	tgRole := StoreRole.StoreRole{}
	if err = dbSet.Core.Take(&tgRole, "ID = ?", reqData.GetID()).Error; err != nil {
		err = Error.Throw(err)
		return
	}
	if tgRole.Name == StoreRole.MinimumAuthorityName ||
		tgRole.Name == StoreRole.WebMasterName {
		err = Error.New("cannot delete default role")
		return
	}

	if err = dbSet.Core.Transaction(func(tx *gorm.DB) (err error) {
		// check whether has binding account
		if err = dbSet.Core.Where("StoreRoleID =?", tgRole.ID).
			Take(&Account.Account{}).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				err = Error.Throw(err)
				return
			}
		}

		if err = dbSet.Core.Unscoped().Delete(&tgRole).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		return nil
	}); err != nil {
		err = Error.Throw(err)
		return
	}

	resp, err = Connect.NewProtoResponse(&V0.DeleteStoreRoleResponse{})
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
