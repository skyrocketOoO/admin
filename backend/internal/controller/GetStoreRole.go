package Middleware

import (
	"taihe.asia/TeaSwap/api/V0"
	"taihe.asia/TeaSwap/internal/Common/Error"
	"taihe.asia/TeaSwap/internal/Common/PrepareBeforeAuth"
	"taihe.asia/TeaSwap/internal/Common/Struct"
	"taihe.asia/TeaSwap/internal/Common/Type"
	"taihe.asia/TeaSwap/internal/Connect"
)

func (it *Handler) GetStoreRole_Gob(
	context Type.Context,
	req *Connect.InnerRequest,
) (response *Connect.InnerResponse, err error) {
	_, _, operRole, err := PrepareBeforeAuth.StoreGet(context)
	if nil != err {
		err = Error.Throw(err)
		return
	}

	out := Struct.DeepNew[V0.GetStoreRoleResponse]()
	if err = Struct.Scan(operRole, out.Role); err != nil {
		err = Error.Throw(err)
		return
	}

	response, err = Connect.NewProtoResponse(out)
	if nil != err {
		err = Error.Throw(err)
		return
	}
	return
}
