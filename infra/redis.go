package infra
//
//import (
//	"fmt"
//	"github.com/go-redis/redis"
//	"github.com/golang/glog"
//)
//
//type RedisClient struct {
//	*redis.Client
//}
//
//func NewRedis() *RedisClient {
//	// RedisClient - 初始化参数
//	client := redis.NewFailoverClient(&redis.FailoverOptions{
//		MasterName:    Config.Redis.Sentinel.MasterName,
//		SentinelAddrs: Config.Redis.Sentinel.Addrs,
//		Password:      Config.Redis.Sentinel.Password,
//		DB:            Config.Redis.Sentinel.Db,
//		PoolSize:      Config.Redis.Sentinel.PoolSize,
//	})
//	p, err := client.Ping().Result()
//	if err != nil {
//		glog.Error(err)
//		panic(err)
//	}
//	glog.Infof("redis ping: %v", p)
//
//	return &RedisClient{client}
//}
//
//func (c *RedisClient) QRCodeKey(key string) string {
//	return fmt.Sprintf("%s:qrcode:%s", Config.AppName, key)
//}

//"github.com/bsm/redislock"
//locker := redislock.New(s.redis)
//lock, err := locker.Obtain(key, 1*time.Minute, nil)
//if err == redislock.ErrNotObtained {
//return nil, infra.FailedPrecondition(Module, SystemBusy, "The system is busy. Please try again later").Err()
//} else if err != nil {
//glog.Error(err)
//return nil, err
//}
//
//defer func() {
//	if err := lock.Release(); err != nil {
//		glog.Warning(err)
//	}
//}()
