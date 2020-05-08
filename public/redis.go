package public

import (
	"errors"
	"tbwisk/common/lib"
	"time"

	"github.com/TBWISK/goconf"
	"github.com/garyburd/redigo/redis"
)

var redisPoolMap map[string]*redis.Pool

//InitRedis 初始化redis
func InitRedis() error {
	redisPool := goconf.InitRedis("demo", 0)
	redisPoolMap["demo"] = redisPool
	return nil
}

//通过配置 执行redis
func RedisConfDo(trace *lib.TraceContext, name string, commandName string, args ...interface{}) (interface{}, error) {
	pool, ok := redisPoolMap[name]
	if ok != false {
		// Log.TagError(trace, "_com_redis_failure", map[string]interface{}{
		// 	"method": commandName,
		// 	"err":    errors.New("RedisConnFactory_error:" + name),
		// 	"bind":   args,
		// })
		return nil, errors.New("no redis pool")
	}
	c := pool.Get()
	defer c.Close()

	startExecTime := time.Now()
	reply, err := c.Do(commandName, args...)
	endExecTime := time.Now()
	if err != nil {
		// Log.TagError(trace, "_com_redis_failure", map[string]interface{}{
		// 	"method":    commandName,
		// 	"err":       err,
		// 	"bind":      args,
		// 	"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		// })
	} else {
		replyStr, _ := redis.String(reply, nil)
		// Log.TagInfo(trace, "_com_redis_success", map[string]interface{}{
		// 	"method":    commandName,
		// 	"bind":      args,
		// 	"reply":     replyStr,
		// 	"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		// })
	}
	return reply, err
}
