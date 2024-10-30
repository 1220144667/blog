/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-23 16:56:06
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-23 16:56:26
 * @FilePath: /blog-service/internal/model/article_tag.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

// ArticleTag 文章标签关系表
type ArticleTag struct {
	*Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

// TableName 表名
func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
