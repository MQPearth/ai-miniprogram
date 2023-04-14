package service

import (
	"ai-backend/common/cons"
	"ai-backend/common/global"
	"ai-backend/common/utils"
	d "ai-backend/model/dto"
	"ai-backend/model/middleware"
	"ai-backend/model/po"
	"gorm.io/gorm"
	"time"
)

type WechatUserService struct{}

func (b *WechatUserService) Login(reqDto d.LoginReqDTO) (d.LoginDTO, error) {
	var dto d.LoginDTO

	now := time.Now()

	if err := global.GDb.Transaction(func(tx *gorm.DB) error {
		user := new(po.WechatUserPO)

		rows := tx.Where("openid = ? AND deleted = ?", reqDto.Openid, 0).Find(user).RowsAffected
		if rows <= 0 {
			user.Openid = reqDto.Openid
			user.CreateDate = now
			user.BasePO.Deleted = false
			user.LastLogin = now
		}
		user.NickName = reqDto.NickName
		user.AvatarUrl = reqDto.AvatarUrl
		user.LastLogin = now

		if err := tx.Save(user).Error; err != nil {
			return err
		}
		dto.NickName = user.NickName
		dto.AvatarUrl = user.AvatarUrl
		dto.Token = utils.Token.GenerateToken()
		loginCache := middleware.RedisLoginCache{}

		loginCache.Token = dto.Token
		loginCache.LoginTime = now
		loginCache.LastRequestTime = now
		loginCache.Id = user.Id

		if err := utils.Redis.Set(utils.Token.GetRedisTokenKey(dto.Token), loginCache, cons.TokenCacheTimeout); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return d.LoginDTO{}, err
	}

	return dto, nil
}
