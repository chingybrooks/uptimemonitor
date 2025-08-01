package uptimemonitor

import "time"

type Session struct {
	ID        int64
	UserID    int64
	Uuid      string
	CreatedAt time.Time
	ExpiresAt time.Time
	User      User
}
