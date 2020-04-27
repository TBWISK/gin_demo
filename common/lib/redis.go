package lib

import (
	"errors"
	"fmt"
	"time"

	"github.com/TBWISK/goconf"
	"github.com/garyburd/redigo/redis"
)

func RedisConnFactory(name string, db int) (*redis.Pool, error) {
	// for confName, cfg := range ConfRedisMap.List {
	// 	if name == confName {
	// 		randHost := cfg.ProxyList[rand.Intn(len(cfg.ProxyList))]
	// 		return redis.Dial(
	// 			"tcp",
	// 			randHost,
	// 			redis.DialConnectTimeout(50*time.Millisecond),
	// 			redis.DialReadTimeout(100*time.Millisecond),
	// 			redis.DialWriteTimeout(100*time.Millisecond))
	// 	}
	// }
	return goconf.InitRedis("demo", db), nil
	// return nil, errors.New("create redis conn fail")
}

func RedisLogDo(trace *TraceContext, c redis.Conn, commandName string, args ...interface{}) (interface{}, error) {
	startExecTime := time.Now()
	reply, err := c.Do(commandName, args...)
	endExecTime := time.Now()
	if err != nil {
		Log.TagError(trace, "_com_redis_failure", map[string]interface{}{
			"method":    commandName,
			"err":       err,
			"bind":      args,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		})
	} else {
		replyStr, _ := redis.String(reply, nil)
		Log.TagInfo(trace, "_com_redis_success", map[string]interface{}{
			"method":    commandName,
			"bind":      args,
			"reply":     replyStr,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		})
	}
	return reply, err
}

//通过配置 执行redis
func RedisConfDo(trace *TraceContext, name string, commandName string, db int, args ...interface{}) (interface{}, error) {
	pool, err := RedisConnFactory(name, db)
	c := pool.Get()
	defer c.Close()
	if err != nil {
		Log.TagError(trace, "_com_redis_failure", map[string]interface{}{
			"method": commandName,
			"err":    errors.New("RedisConnFactory_error:" + name),
			"bind":   args,
		})
		return nil, err
	}
	defer c.Close()

	startExecTime := time.Now()
	reply, err := c.Do(commandName, args...)
	endExecTime := time.Now()
	if err != nil {
		Log.TagError(trace, "_com_redis_failure", map[string]interface{}{
			"method":    commandName,
			"err":       err,
			"bind":      args,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		})
	} else {
		replyStr, _ := redis.String(reply, nil)
		Log.TagInfo(trace, "_com_redis_success", map[string]interface{}{
			"method":    commandName,
			"bind":      args,
			"reply":     replyStr,
			"proc_time": fmt.Sprintf("%fs", endExecTime.Sub(startExecTime).Seconds()),
		})
	}
	return reply, err
}
