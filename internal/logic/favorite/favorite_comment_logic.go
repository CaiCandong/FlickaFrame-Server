package favorite

import (
	"context"
	favorite_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteCommentLogic {
	return &FavoriteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteCommentLogic) FavoriteComment(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.FavoriteModel.Create(l.ctx,
		util.MustString2Int64(req.TargetId),
		doerId, favorite_model.CommentFavoriteType)
	return
}