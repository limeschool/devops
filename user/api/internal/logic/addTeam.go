package logic

import (
	"context"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeamLogic {
	return &AddTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTeamLogic) AddTeam(req *types.AddTeamRequest) error {
	team := models.Team{}
	tb := l.svcCtx.Orm.Table(team.Table())
	return tb.Create(req).Error
}
