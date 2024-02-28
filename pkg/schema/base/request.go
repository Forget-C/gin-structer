package base

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

var (
	CtxUserNameKey = "userName"
	CtxUserIDKey   = "userID"
)

type RequestMeta struct {
	RequestUserID   string          `json:"-" form:"-" uri:"-"` // 请求用户ID
	RequestUserName string          `json:"-" form:"-" uri:"-"` // 请求用户名
	Ctx             context.Context `json:"-" form:"-" uri:"-"` // 上下文
	GinCtx          *gin.Context    `json:"-" form:"-" uri:"-"` // gin上下文
}

func (b *RequestMeta) GetCtx() context.Context {
	if b.Ctx == nil {
		return context.Background()
	}
	return b.Ctx
}

func (b *RequestMeta) Modify(c *gin.Context) {
	b.Ctx = c.Request.Context()
	b.GinCtx = c
	username, _ := c.Get(CtxUserNameKey)
	userID, _ := c.Get(CtxUserIDKey)
	b.RequestUserName = cast.ToString(username)
	b.RequestUserID = cast.ToString(userID)
}

type PaginationReq struct {
	Page     uint32 `json:"page" form:"page"`
	PageSize uint32 `json:"page_size" form:"page_size"`
}

const (
	maxPageSize = 300
)

func (c *PaginationReq) Modify(ctx *gin.Context) {
	if c.PageSize > maxPageSize {
		c.PageSize = maxPageSize
	}
	if c.PageSize == 0 {
		c.PageSize = 50
	}
	if c.Page == 0 {
		c.Page = 1
	}
}

func (c *PaginationReq) Offset() int64 {
	return int64((c.Page - 1) * c.PageSize)
}

func (c *PaginationReq) Limit() int64 {
	return int64(c.PageSize)
}

type DefaultListReq struct {
	PaginationReq
	RequestMeta
	SQLSearchReq
	SQLSortReq
}

func (c *DefaultListReq) Modify(ctx *gin.Context) {
	c.PaginationReq.Modify(ctx)
	c.RequestMeta.Modify(ctx)
}
