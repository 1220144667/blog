/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-23 16:55:07
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-30 10:49:41
 * @FilePath: /blog-service/internal/model/article.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "blog-service/pkg/app"

// Article 文章模型
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

// TableName 表名
func (a *Article) TableName() string {
	return "blog_article"
}
