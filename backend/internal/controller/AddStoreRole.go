package Middleware

import (
	"gorm.io/gorm"
	"taihe.asia/TeaSwap/api/V0"
	"taihe.asia/TeaSwap/internal/Common/Bytes"
	"taihe.asia/TeaSwap/internal/Common/Error"
	"taihe.asia/TeaSwap/internal/Common/PrepareBeforeAuth"
	"taihe.asia/TeaSwap/internal/Common/Struct"
	"taihe.asia/TeaSwap/internal/Common/Type"
	"taihe.asia/TeaSwap/internal/Connect"
	"taihe.asia/TeaSwap/internal/model/entity/Core/StoreRole"
)

func (it *Handler) AddStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (response *Connect.InnerResponse, err error) {
	reqData, err := Bytes.ProtoDecode[V0.AddStoreRoleRequest](req.Msg.Data)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	dbSet, operAccount, operRole, err := PrepareBeforeAuth.StoreGet(context)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	if !operRole.Page.Role.Create {
		err = Error.New("not permitted")
		err = Error.Throw(err)
		return
	}

	if err = dbSet.Core.Transaction(func(tx *gorm.DB) (err error) {
		res := dbSet.Core.
			Where("StoreID = ?", operAccount.EntryID).
			Where("Name = ?", reqData.GetName()).
			Take(&StoreRole.StoreRole{})
		if res.Error != nil {
			if res.Error != gorm.ErrRecordNotFound {
				err = Error.Throw(res.Error)
				return
			}
		}
		if res.RowsAffected > 0 {
			err = Error.New("role name already exists")
			err = Error.Throw(err)
			return
		}

		tgRole := Struct.DeepNew[StoreRole.StoreRole]()
		tgRole.StoreID = operAccount.EntryID
		if err = Struct.Scan(reqData, tgRole); err != nil {
			err = Error.Throw(err)
			return
		}

		if err = dbSet.Core.Create(tgRole).Error; err != nil {
			err = Error.Throw(err)
			return
		}
		return nil
	}); err != nil {
		err = Error.Throw(err)
		return
	}

	response, err = Connect.NewProtoResponse(&V0.AddStoreRoleResponse{})
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
