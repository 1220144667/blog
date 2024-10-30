/*
 * @Author: 洪陪 hp2022a@163.com
 * @Date: 2024-10-29 16:20:57
 * @LastEditors: 洪陪 hp2022a@163.com
 * @LastEditTime: 2024-10-29 19:20:21
 * @FilePath: /blog-service/global/setting.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package global

import (
	"blog-service/pkg/setting"

	"blog-service/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
