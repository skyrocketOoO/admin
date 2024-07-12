package Middleware

import (
	"taihe.asia/TeaSwap/api/V0"
	V0Account "taihe.asia/TeaSwap/api/V0/Model/Account"
	"taihe.asia/TeaSwap/internal/Common/Bytes"
	"taihe.asia/TeaSwap/internal/Common/Error"
	"taihe.asia/TeaSwap/internal/Common/PrepareBeforeAuth"
	"taihe.asia/TeaSwap/internal/Common/Struct"
	"taihe.asia/TeaSwap/internal/Common/Type"
	"taihe.asia/TeaSwap/internal/Common/orm"
	"taihe.asia/TeaSwap/internal/Connect"
	"taihe.asia/TeaSwap/internal/model/entity/Core/Account"
	"taihe.asia/TeaSwap/internal/model/entity/Core/StoreRole"
)

func (it *Handler) ListStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (response *Connect.InnerResponse, err error) {
	reqData, err := Bytes.ProtoDecode[V0.ListStoreRoleRequest](req.Msg.Data)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	dbSet, operAccount, operRole, err := PrepareBeforeAuth.StoreGet(context)
	if nil != err {
		err = Error.Throw(err)
		return
	}
	if !operRole.Page.Role.Read {
		err = Error.New("not permitted")
		return
	}

	out := new(V0.ListStoreRoleResponse)
	db := dbSet.Core.Unscoped().
		Model(&StoreRole.Entity{}).
		Where("StoreID = ?", operAccount.EntryID)
	if reqData.GetSorter() != nil {
		if !reqData.Sorter.Asc {
			db = db.Order(reqData.Sorter.Field + " DESC")
		} else {
			db = db.Order(reqData.Sorter.Field)
		}
	} else {
		db = db.Order("Name")
	}
	var storeRoles []StoreRole.Entity
	if out.Total, err = orm.ListWithCount(db,
		orm.Pager{
			Number: reqData.GetPager().GetNumber(),
			Size:   reqData.GetPager().GetSize(),
		},
		&storeRoles); nil != err {
		err = Error.Throw(err)
		return
	}

	out.List = make([]*V0.ListStoreRoleData, len(storeRoles))
	for i, storeRole := range storeRoles {
		out.List[i] = Struct.DeepNew[V0.ListStoreRoleData]()
		if err = Struct.Scan(&storeRole, out.List[i].Role); err != nil {
			err = Error.Throw(err)
			return
		}

		var accounts []Account.Account
		if err = dbSet.Core.Where("StoreRoleID = ?", storeRole.ID).Find(&accounts).
			Error; err != nil {
			err = Error.Throw(err)
			return
		}
		out.List[i].Accounts = make([]*V0Account.Entity, len(accounts))
		for j, account := range accounts {
			out.List[i].Accounts[j] = Struct.DeepNew[V0Account.Entity]()
			if err = Struct.Scan(&account, out.List[i].Accounts[j]); err != nil {
				err = Error.Throw(err)
				return
			}
		}
	}

	response, err = Connect.NewProtoResponse(out)
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
