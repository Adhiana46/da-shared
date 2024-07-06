package constants

import "time"

const (
	ACCESS_TOKEN_DURATION  = time.Hour * time.Duration(2)  // 2 hours
	REFRESH_TOKEN_DURATION = time.Hour * time.Duration(24) // 24 hours
	USER_KEY_CTX           = "user"                        // use in middleware
)
