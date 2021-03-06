package logic

import (
	"context"
	"devops/common/tools"
	"devops/common/tools/tree"
	"devops/user/api/internal/svc"
	"devops/user/api/internal/types"
	"devops/user/models"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeamLogic {
	return &GetTeamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeamLogic) GetTeam() (resp *types.GetTeamResponse, err error) {
	var list []models.Team
	resp = new(types.GetTeamResponse)

	team := models.Team{}
	list, _, err = team.All(nil)
	var teamTree []TeamTree
	tools.Transform(list, &teamTree)
	nodeArray := make([]tree.Tree, len(teamTree))
	for i := 0; i < len(teamTree); i++ {
		nodeArray[i] = &teamTree[i]
	}
	//进行转菜单树
	root := tree.BuildTree(nodeArray)
	tools.Transform(root, resp)
	return
}

type TeamTree struct {
	types.GetTeamResponse
}

func (t *TeamTree) ID() int64 {
	return t.GetTeamResponse.ID
}

func (t *TeamTree) ParentID() int64 {
	return t.GetTeamResponse.ParentID
}

func (t *TeamTree) AppendChildren(node interface{}) {
	n := node.(*TeamTree)
	t.Children = append(t.Children, &n.GetTeamResponse)
}
