# go-blog-backend

## 使用的框架和库
- `Gin`
- `GORM v2`
- `Viper`
- `logrus`
- `jwt-go`

日志采用简单的`Logrus`库，只是为了调试用，后续考虑使用`Zap`。

数据库方面，`Article`表和`Tag`表是多对多关系，但两表未设置外键，采用增加第三张表`ArticleTag`的形式，记录`articleID`和`tagID`。

关于`GORM`框架的使用，这里只是为了学习一下，实际最好是自己写`SQL`语句创建数据库，对数据库的数据能自己完全把控。

## Docker部署
`Dockerfile`已经更新并测试过，数据库需要再另开一个`MySQL`镜像搭配使用。

## Future
未来会考虑加上`Redis`缓存机制，文章评论，也可能会尝试`Web`端展示文章。