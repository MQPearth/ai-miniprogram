package utils

import (
	"ai-backend/common/global"
	"context"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

// GetAnswer 单次问答
func GetAnswer(question string) (string, error) {
	resp, err := global.GAi.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			//Model: openai.GPT3Dot5Turbo,
			Model: global.GConfig.Ai.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		global.GLog.Error("ChatCompletion error:", zap.Error(err))
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	global.GLog.Info("answer", zap.String("answer", answer))
	return answer, nil
}
