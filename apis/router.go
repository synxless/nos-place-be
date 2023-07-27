package apis

import (
	"nosplace/storage"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Router struct {
	Port       string
	Domain     string
	Storage    *storage.Storage
	RedisCfg   *redis.Options
	DefaultRPC string
}

func (rt *Router) Start() error {
	var cacheStore persist.CacheStore

	cacheStore = persist.NewMemoryStore(1 * time.Minute)

	if rt.RedisCfg != nil {
		cacheStore = persist.NewRedisStore(redis.NewClient(rt.RedisCfg))
	}

	r := gin.Default()

	r.Use(cors.Default())

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	apiv1 := r.Group("/api/v1")
	_ = cacheStore
	_ = apiv1
	apiv1.GET("/pixels", cache.CacheByRequestURI(cacheStore, 1*time.Second), rt.GetPixels)
	apiv1.GET("/latest", cache.CacheByRequestURI(cacheStore, 1*time.Second), rt.GetLatest)
	// apiv1.GET("/tokens", cache.CacheByRequestURI(cacheStore, 10*time.Second), rt.GetTokens)
	// apiv1.GET("/token/:address", cache.CacheByRequestURI(cacheStore, 10*time.Second), rt.GetToken)
	// // apiv1.GET("/token/:address/holders", cache.CacheByRequestURI(cacheStore, 2*time.Second), rt.GetTokenHolders)
	// apiv1.GET("/:address/tokens", cache.CacheByRequestURI(cacheStore, 5*time.Second), rt.GetHolderTokens)
	// apiv1.GET("/:address/:token/balance", cache.CacheByRequestURI(cacheStore, 5*time.Second), rt.GetHolderTokenBalance)
	// apiv1.GET("/search", rt.Search)

	err := r.Run("0.0.0.0:" + rt.Port)
	if err != nil {
		return err
	}
	return nil
}
