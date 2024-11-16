package controller

import (
	"context"

	"admin/proto"

	"connectrpc.com/connect"
	"github.com/skyrocketOoO/Guess-number/usecase"
)

func (s *SideProjectServer) GuessNumber(
	ctx context.Context,
	connReq *connect.Request[proto.GuessNumberReq],
) (connResp *connect.Response[proto.GuessNubmerResp], err error) {
	req := connReq.Msg

	props := make([][]float64, len(req.GetProportions()))
	for i, p := range req.GetProportions() {
		props[i] = make([]float64, len(p.Values))
		for j, v := range p.Values {
			props[i][j] = float64(v)
		}
	}

	usecase.GuessNumbers(int(req.GetLowerBound()), int(req.GetUpperBound()), props)
	return
}
