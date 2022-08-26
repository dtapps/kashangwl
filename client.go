package kashangwl

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		customerId  int    // 商家编号
		customerKey string // 商家密钥
	}
	log struct {
		gormClient     *dorm.GormClient  // 日志数据库
		gorm           bool              // 日志开关
		logGormClient  *golog.ApiClient  // 日志服务
		mongoClient    *dorm.MongoClient // 日志数据库
		mongo          bool              // 日志开关
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// NewClient 创建实例化
// customerId 商家编号
// customerKey 商家密钥
func NewClient(customerId int, customerKey string, gormClientFun gormClientFun, mongoClientFun mongoClientFun, debug bool) (*Client, error) {

	var err error
	c := &Client{}

	c.config.customerId = customerId
	c.config.customerKey = customerKey

	c.requestClient = gorequest.NewHttp()

	gormClient := gormClientFun()
	if gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(func() (*dorm.GormClient, string) {
			return gormClient, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
	}
	c.log.gormClient = gormClient

	mongoClient, databaseName := mongoClientFun()
	if mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(func() (*dorm.MongoClient, string, string) {
			return mongoClient, databaseName, logTable
		}, debug)
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
	}
	c.log.mongoClient = mongoClient

	return c, nil
}
