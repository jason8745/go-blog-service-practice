package v1

import "github.com/gin-gonic/gin"

//對應路由的處理方法
type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context)    {}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
