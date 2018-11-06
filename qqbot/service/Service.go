package service

import (
	"github.com/mapleFU/QQBot/qqbot/data/group"
)

type Servicer interface {
	PutRequest(Request *group.ChatRequestData)
} 