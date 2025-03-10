package organs

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type V2getUnitIdMapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewV2getUnitIdMapLogic(ctx context.Context, svcCtx *svc.ServiceContext) V2getUnitIdMapLogic {
	return V2getUnitIdMapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *V2getUnitIdMapLogic) V2getUnitIdMap(req types.V2GetUnitIdMapReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
