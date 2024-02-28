package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	schema "github.com/Forget-C/http-structer/pkg/schema/base"
)

type Collection struct {
	alwaysCode int
	setAlways  bool
}

func NewCollection(alwaysCode int) *Collection {
	return &Collection{alwaysCode: alwaysCode, setAlways: alwaysCode != -1}
}

func (c *Collection) AcceptedResponse(ctx *gin.Context) {
	resp := schema.DefaultResp{}
	resp.SetAccepted()
	if c.setAlways {
		ctx.AbortWithStatusJSON(c.alwaysCode, resp)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusAccepted, resp)
}

func (c *Collection) BadRequestResponse(ctx *gin.Context, message string) {
	resp := schema.DefaultResp{}
	resp.SetBadRequest(message)
	if c.setAlways {
		ctx.AbortWithStatusJSON(c.alwaysCode, resp)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
}

func (c *Collection) NotFoundResponse(ctx *gin.Context, message string) {
	resp := schema.DefaultResp{}
	resp.SetNotFound(message)
	if c.setAlways {
		ctx.AbortWithStatusJSON(c.alwaysCode, resp)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusNotFound, resp)
}

func (c *Collection) ServerErrorResponse(ctx *gin.Context, message string) {
	resp := schema.DefaultResp{}
	resp.SetServerError(message)
	if c.setAlways {
		ctx.AbortWithStatusJSON(c.alwaysCode, resp)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}

func (c *Collection) AutoResponse(ctx *gin.Context, resp schema.Response, err error) {
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.NotFoundResponse(ctx, err.Error())
			return
		}
		c.BadRequestResponse(ctx, err.Error())
		return
	}
	resp.SetSuccess()
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
