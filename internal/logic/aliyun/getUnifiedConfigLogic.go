// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package aliyun

import (
	"context"
	"fmt"
	"time"

	"github.com/gpencil/upload/internal/svc"
	"github.com/gpencil/upload/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUnifiedConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取统一配置
func NewGetUnifiedConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUnifiedConfigLogic {
	return &GetUnifiedConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUnifiedConfigLogic) GetUnifiedConfig(req *types.GetUnifiedConfigReq) (resp *types.UnifiedConfigResponse, err error) {
	vc := l.svcCtx.Query.Voices
	// 构建查询条件
	query := vc.WithContext(l.ctx)

	// 如果指定了 voice_id，则查询该声线
	if req.VoiceId != "" {
		query = query.Where(vc.VoiceID.Eq(req.VoiceId))
	}

	// 如果指定了 scenario_id，则查询该场景下的声线
	if req.ScenarioId != 0 {
		query = query.Where(vc.ScenarioID.Eq(req.ScenarioId))
	}

	// 只查询启用状态的声线
	query = query.Where(vc.Status.Is(true))

	// 按排序权重排序
	voices, err := query.Order(vc.SortOrder).Find()
	if err != nil {
		l.Errorf("查询声线失败: %v", err)
		return nil, fmt.Errorf("查询声线失败")
	}

	// 转换为响应格式
	var voiceResponses []types.VoiceResponse
	for _, voice := range voices {
		voiceResp := types.VoiceResponse{
			Id:          voice.ID,
			VoiceId:     voice.VoiceID,
			ScenarioId:  voice.ScenarioID,
			Name:        voice.Name,
			IconUrl:     l.getStringValue(voice.IconURL),
			Description: l.getStringValue(voice.Description),
			PreviewUrl:  l.getStringValue(voice.PreviewURL),
			PreviewDesc: l.getStringValue(voice.PreviewDesc),
			Gender:      l.getStringValue(voice.Gender),
			AgeGroup:    l.getStringValue(voice.AgeGroup),
			Style:       l.getStringValue(voice.Style),
			Language:    l.getStringValue(voice.Language),
			SortOrder:   l.getInt32Value(voice.SortOrder),
			Status:      l.getBoolValue(voice.Status),
			CreatedAt:   time.Unix(voice.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(voice.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
		}
		voiceResponses = append(voiceResponses, voiceResp)
	}

	// 构建响应（简化版，只返回声线数据）
	resp = &types.UnifiedConfigResponse{
		FullConfigs: []types.ScenarioConfigResponse{
			{
				Scenario: types.ScenarioResponse{
					Id:          req.ScenarioId,
					Name:        "默认场景",
					Description: "默认场景描述",
					ImageUrl:    "",
					Status:      1,
					CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
					UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
				},
				Voices:      voiceResponses,
				SampleFiles: []types.SampleFileResponse{},
			},
		},
	}

	return resp, nil
}

func (l *GetUnifiedConfigLogic) getStringValue(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (l *GetUnifiedConfigLogic) getInt32Value(i *int32) int32 {
	if i == nil {
		return 0
	}
	return *i
}

func (l *GetUnifiedConfigLogic) getBoolValue(b *bool) int32 {
	if b == nil || !*b {
		return 0
	}
	return 1
}
