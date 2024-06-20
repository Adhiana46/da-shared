package constants

const (
	CacheInvalidateTokenKey = "token-blacklist:%s" // %s => md5(token)
	CacheUserKey            = "user:%s"            // %s => id|email
)
