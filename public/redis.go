package public

import (
	"github.com/TBWISK/goconf"
	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

//InitRedis 初始化redis
func InitRedis() error {
	redisPool = goconf.InitRedis("demo", 1)
	return nil
}
func RedisConfDo() {

}
