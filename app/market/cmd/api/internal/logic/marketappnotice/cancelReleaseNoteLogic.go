package marketappnotice

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"orginone/app/market/cmd/api/internal/svc"
	"orginone/app/market/cmd/api/internal/types"
	"orginone/common/rpcs/market"
	"orginone/common/schema"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelReleaseNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelReleaseNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelReleaseNoteLogic {
	return CancelReleaseNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelReleaseNoteLogic) CancelReleaseNote(req types.CancelReleaseNoteReq) (resp *types.CommonResponse, err error) {
	rpcRes, err := l.svcCtx.MarketRpc.FindMarkAppNoticeByIds(l.ctx, &market.MarkAppNoticeIdsReq{
		NotoceIds: []int64{req.Id},
	})
	if err != nil {
		return &types.CommonResponse{}, err
	}
	var noticeEntitys []*schema.AsMarketAppNotice
	err = json.Unmarshal([]byte(rpcRes.Json), &noticeEntitys)
	if err != nil {
		return &types.CommonResponse{}, err
	}
	if len(noticeEntitys) > 0 {
		if noticeEntitys[0].NoticeReleaseStatus == 0 {
			return types.Failed(http.StatusInternalServerError, errors.New("该通知已下架!"))
		} else {
			count, err := l.svcCtx.MarketRpc.UpdateMarkAppNoticesStatus(l.ctx, &market.UpdateMarkAppNoticesStatusReq{
				NotoceId:     req.Id,
				NoticeStatus: 0,
				IsDeleted:    noticeEntitys[0].IsDeleted,
			})
			return types.JsonResult(count.Json == "1", err)
		}
	} else {
		return types.Failed(http.StatusInternalServerError, errors.New("未找到该条通知!"))
	}
}
