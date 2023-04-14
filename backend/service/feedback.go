package service

import (
	"ai-backend/common/global"
	"ai-backend/model/dto"
	"ai-backend/model/po"
	"time"
)

type FeedBackService struct{}

func (b *FeedBackService) SaveFeedback(dto *dto.SaveFeedbackDTO) error {
	feedback := po.FeedbackPO{}
	feedback.UserId = dto.UserId
	feedback.CreateDate = time.Now()
	feedback.Content = dto.Content
	tx := global.GDb.Model(&feedback).Save(&feedback)
	return tx.Error
}
