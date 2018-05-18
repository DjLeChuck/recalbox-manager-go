package store

import (
	"github.com/kataras/iris/sessions"
)

var (
	cookieNameForSessionID = "t|-|eC4keIsA|_i3"
	// Sessions store the iris session.
	Sessions = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)
