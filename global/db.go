/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-29 18:05:48
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-29 18:07:16
 * @FilePath: /blog-service/global/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package global

import "github.com/jinzhu/gorm"

var (
	DBEngine *gorm.DB
)
