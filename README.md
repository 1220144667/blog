# Blog Service 介绍

## 1.目录结构
- configs：配置文件。
- docs：文档集合。
- global：全局变量。
- internal：内部模块。
  - dao：数据访问层（Database Access Object），所有与数据相关的操作都会在 dao 层进行，例如 MySQL、ElasticSearch 等。
  - middleware：HTTP 中间件。
  - model：模型层，用于存放 model 对象。
  - routers：路由相关逻辑处理。
  - service：项目核心业务逻辑。
- pkg：项目相关的模块包。
- storage：项目生成的临时文件。
- scripts：各类构建，安装，分析等操作的脚本。
- third_party：第三方的资源工具，例如 Swagger UI。

## 2.数据库设计
- 项目使用MySQL数据库，数据库设计如下：
  - 标签管理：文章所归属的分类，也就是标签。我们平时都会针对文章的内容打上好几个标签，用于标识文章内容的要点要素，这样子便于读者的识别和 SEO 的收录等。
  - 文章管理：整个文章内容的管理，并且需要将文章和标签进行关联。

## 3.公共组件
- 项目中公共组件的设计，我们将公共组件分为以下几类：
  - 公共错误码
  - 配置管理
  - 数据库连接
  - 日志管理
  - 响应处理
- 1.公共错误码
  - Success                   = NewError(0, "成功")
  - ServerError               = NewError(10000000, "服务内部错误")
  - InvalidParams             = NewError(10000001, "入参错误")
  - NotFound                  = NewError(10000002, "找不到")
  - UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 - AppSecret")
  - UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token 错误")
  - UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token 超时")
  - UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token 生成失败")
  - TooManyRequests           = NewError(10000007, "请求过多")
- 2.配置管理
  - 项目中，我们使用 YAML 格式的配置文件进行配置，并且使用 viper 进行配置管理，配置文件存放在 configs 目录下。
    - Server：服务配置，设置 gin 的运行模式、默认的 HTTP 监听端口、允许读取和写入的最大持续时间。
    - App：应用配置，设置默认每页数量、所允许的最大每页数量以及默认的应用日志存储路径。
    - Database：数据库配置，主要是连接实例所必需的基础参数。
