// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package aliyun

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gpencil/upload/internal/dal/model"
	"github.com/gpencil/upload/internal/svc"
	"github.com/gpencil/upload/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVoiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传/修改音线
func NewUploadVoiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVoiceLogic {
	return &UploadVoiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVoiceLogic) UploadVoice(req *types.UploadVoiceReq) (resp *types.EmptyRequest, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic:", e)
		}
	}()

	// 将 VoiceId 转换为小写
	voiceId := strings.ToLower(req.VoiceId)

	// 查询是否存在该 voice_id 的声线
	vc := l.svcCtx.Query.Voices
	existingVoice, err := vc.WithContext(l.ctx).Where(vc.VoiceID.Eq(voiceId)).First()

	now := time.Now().Unix()
	status := req.Status == 1
	sortOrder := int32(req.SortOrder)

	if err != nil {
		// 如果不存在，则创建新的声线
		voice := &model.Voices{
			VoiceID:     voiceId,
			ScenarioID:  req.ScenarioId,
			Name:        req.Name,
			IconURL:     &req.IconUrl,
			Description: &req.Description,
			PreviewURL:  &req.PreviewUrl,
			Gender:      &req.Gender,
			AgeGroup:    &req.AgeGroup,
			Style:       &req.Style,
			Language:    &req.Language,
			SortOrder:   &sortOrder,
			Status:      &status,
			CreatedAt:   now,
			UpdatedAt:   now,
		}

		err = vc.WithContext(l.ctx).Create(voice)
		if err != nil {
			l.Errorf("创建声线失败: %v", err)
			return nil, fmt.Errorf("创建声线失败")
		}
	} else {
		// 如果存在，则更新声线
		updateData := map[string]interface{}{
			"scenario_id": req.ScenarioId,
			"name":        req.Name,
			"icon_url":    req.IconUrl,
			"description": req.Description,
			"preview_url": req.PreviewUrl,
			"gender":      req.Gender,
			"age_group":   req.AgeGroup,
			"style":       req.Style,
			"language":    req.Language,
			"sort_order":  sortOrder,
			"status":      status,
			"updated_at":  now,
		}

		// 检查 preview_url 或 description 是否发生变化，如果变化则清空 voice_hash
		if (existingVoice.PreviewURL != nil && *existingVoice.PreviewURL != req.PreviewUrl) ||
			(existingVoice.Description != nil && *existingVoice.Description != req.Description) ||
			(existingVoice.PreviewURL == nil && req.PreviewUrl != "") ||
			(existingVoice.Description == nil && req.Description != "") {
			updateData["voice_hash"] = ""
			l.Infof("preview_url 或 description 发生变化，清空 voice_hash")
		}

		_, err = vc.WithContext(l.ctx).Where(vc.VoiceID.Eq(voiceId)).Updates(updateData)

		if err != nil {
			l.Errorf("更新声线失败: %v", err)
			return nil, fmt.Errorf("更新声线失败")
		}
	}

	l.Infof("声线 %s 操作成功, 现有ID: %d", voiceId, existingVoice.ID)

	return &types.EmptyRequest{}, nil
}
