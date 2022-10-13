package v1

import (
	"go-blog-service/global"
	"go-blog-service/pkg/app"
	"go-blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Tag路由的Handler
type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=01"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
