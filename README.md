# go-blog-backend

## 使用的框架和库
- Gin
- GORM v2：目前使用不太熟练。
- Viper
- logrus：本来打算学学zap的。
- jwt-go


对于博客系统，如果只是为了展示需要就不需要user模块，所以user单独做一个表并切断和其他表联系，这里只是为了学习注册登录验证。目前只做了注册模块，后续再加上登录、JWT和session存储。

关于GORM框架的使用，这里只是为了学习一下，实际最好是自己写SQL语句创建数据库，对数据库的数据能自己完全把控。

## Docker部署
Dockerfile已经更新并测试过，数据库需要再另开一个MySQL镜像搭配使用。

## Future
未来会考虑加上Redis缓存机制，文章评论，也可能会尝试Web端展示文章。