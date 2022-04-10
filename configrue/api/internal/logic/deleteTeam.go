package logic

import (
	"configure/models"
	"context"

	"configure/api/internal/svc"
	"configure/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTeamLogic {
	return &DeleteTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTeamLogic) DeleteTeam(req *types.DeleteTeamRequest) error {
	team := models.Team{}
	tb := l.svcCtx.Orm.Table(team.Table())
	return tb.Delete(&team, req.ID).Error
}