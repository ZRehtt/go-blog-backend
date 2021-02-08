package settings

import "github.com/sirupsen/logrus"

//NewLogger 日志配置初始化
func NewLogger() {
	log := logrus.New()
	//这里用logrus简单记录日志，设置日志记录级别为Debug，只调试使用
	log.SetLevel(logrus.DebugLevel)
	//将调用的函数名添加为字段"method"，但会增加性能开销
	logrus.SetReportCaller(true)
	//日志输出格式
	log.SetFormatter(&logrus.JSONFormatter{})
}
