package logic

import (
	"context"
	"devops/common/tools"
	"devops/configrue/models"

	"devops/configrue/api/internal/svc"
	"devops/configrue/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnvironmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEnvironmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnvironmentLogic {
	return &GetEnvironmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEnvironmentLogic) GetEnvironment() (resp *types.GetEnvironmentResponse, err error) {
	env := models.Environment{}
	resp = new(types.GetEnvironmentResponse)
	list, _, err := env.All(nil)
	tools.Transform(list, &resp.List)
	return resp, err
}
