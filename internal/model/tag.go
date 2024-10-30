/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-23 16:53:19
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-23 16:54:01
 * @FilePath: /blog-service/internal/model/tag.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

import "blog-service/pkg/app"

// Tag 标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// TableName 表名
func (t Tag) TableName() string {
	return "blog_tag"
}
