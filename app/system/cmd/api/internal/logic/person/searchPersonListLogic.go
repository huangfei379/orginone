package person

import (
	"context"

	"orginone/app/system/cmd/api/internal/svc"
	"orginone/app/system/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPersonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchPersonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchPersonListLogic {
	return SearchPersonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchPersonListLogic) SearchPersonList(req types.SearchPersonListReq) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
