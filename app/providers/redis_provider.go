package providers

import (
	"github.com/raiworks/rapidgo/v2/core/cache"
	"github.com/raiworks/rapidgo/v2/core/container"
)

// RedisProvider registers a shared *redis.Client singleton in the container.
// Uses cache.NewRedisClient() which reads REDIS_HOST, REDIS_PORT, REDIS_PASSWORD,
// REDIS_DB, REDIS_POOL_SIZE, REDIS_DIAL_TIMEOUT, REDIS_READ_TIMEOUT, REDIS_WRITE_TIMEOUT
// from the environment.
type RedisProvider struct{}

// Register binds a *redis.Client singleton. Connection is lazy.
func (p *RedisProvider) Register(c *container.Container) {
	// Default client — reads REDIS_DB from env (default 0).
	c.Singleton("redis", func(_ *container.Container) interface{} {
		return cache.NewRedisClient(nil)
	})

	// Named clients for multi-DB isolation — uncomment and adjust as needed:
	// c.Singleton("redis.cache",  func(_ *container.Container) interface{} { db := 2; return cache.NewRedisClient(&db) })
	// c.Singleton("redis.ssr",    func(_ *container.Container) interface{} { db := 3; return cache.NewRedisClient(&db) })
	// c.Singleton("redis.pubsub", func(_ *container.Container) interface{} { db := 5; return cache.NewRedisClient(&db) })
}

// Boot is a no-op.
func (p *RedisProvider) Boot(c *container.Container) {}