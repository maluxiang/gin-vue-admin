package global

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mark3labs/mcp-go/server"
	"github.com/qiniu/qmgo"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB        *gorm.DB
	GVA_DBList    map[string]*gorm.DB
	GVA_RDS       *redis.Client
	GVA_REDIS     redis.UniversalClient
	GVA_REDISList map[string]redis.UniversalClient
	GVA_MONGO     *qmgo.QmgoClient
	GVA_CONFIG    config.Server
	GVA_VP        *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG                 *zap.Logger
	GVA_Timer               timer.Timer = timer.NewTimerTask()
	GVA_Concurrency_Control             = &singleflight.Group{}
	GVA_ROUTERS             gin.RoutesInfo
	GVA_ACTIVE_DBNAME       *string
	GVA_MCP_SERVER          *server.MCPServer
	BlackCache              local_cache.Cache
	lock                    sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func GetRedis(name string) redis.UniversalClient {
	redis, ok := GVA_REDISList[name]
	if !ok || redis == nil {
		panic(fmt.Sprintf("redis `%s` no init", name))
	}
	return redis
}

func InitRedis() {
	redisCfg := GVA_CONFIG.Redis

	GVA_RDS = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	// 测试连接
	_, err := GVA_RDS.Ping(context.Background()).Result()
	if err != nil {
		GVA_LOG.Fatal("Redis连接失败", zap.Error(err))
	}
	GVA_LOG.Info("Redis连接成功")
}
