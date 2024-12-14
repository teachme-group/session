package redis

import (
	_ "embed"
)

var (
	//go:embed scripts/fetch_user_session.lua
	fetchUserSessionScript string
)
