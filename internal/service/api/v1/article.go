/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-23 17:00:32
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-29 21:36:20
 * @FilePath: /blog-service/internal/routers/api/v1/article.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

// @Summary 获取多个文章
// @Produce  json
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (a Article) List(c *gin.Context) {}

func (a Article) Create(c *gin.Context) {}

func (a Article) Update(c *gin.Context) {}

func (a Article) Delete(c *gin.Context) {}
