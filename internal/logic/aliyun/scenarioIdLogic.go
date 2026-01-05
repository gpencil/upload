// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package aliyun

import (
	"context"

	"github.com/gpencil/upload/internal/svc"
	"github.com/gpencil/upload/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScenarioIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询列ID
func NewScenarioIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScenarioIdLogic {
	return &ScenarioIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ScenarioIdLogic) ScenarioId(req *types.EmptyRequest) (resp *types.ScenarioIdResp, err error) {
	sos := l.svcCtx.Query.Scenarios
	scenarios, err := sos.WithContext(l.ctx).Where(sos.Status.Is(true)).Order(sos.ID).Find()
	if err != nil {
		l.Errorf("查询场景列表失败: %v", err)
		return nil, err
	}

	// 转换为响应格式
	var result []types.ScenarioIdData
	for _, scenario := range scenarios {
		result = append(result, types.ScenarioIdData{
			Id:   scenario.ID,
			Name: scenario.Name,
		})
	}

	return &types.ScenarioIdResp{ScenarioIds: result}, nil
}
