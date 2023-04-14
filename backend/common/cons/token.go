package cons

import "time"

const (
	// RequestHeaderTokenKey 请求头 token key
	RequestHeaderTokenKey = "token"
	// TokenCacheKeyPrefix redis 缓存前缀
	TokenCacheKeyPrefix = "wechat_user:token:"
	// TokenCacheTimeout redis 缓存过期时间
	TokenCacheTimeout = 2 * time.Hour
	RequestInfoCache  = "request_info_cache"
)
