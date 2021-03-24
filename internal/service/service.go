package service

import (
	"context"
	"github.com/ZRehtt/go-blog-backend/globals"
	"gorm.io/gorm"
)

type Service struct {
	ctx context.Context
	db  *gorm.DB
}

func NewService(ctx context.Context) *Service {
	stc := Service{ctx: ctx}
	stc.db = globals.GDB
	return &stc
}

