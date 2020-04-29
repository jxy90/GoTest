package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"sync"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(4)
	rdb := NewRedisClient(1)
	defer rdb.Close()
	wg := sync.WaitGroup{}
	shopUser := make([]string, 0)
	wg.Add(1)
	for i := 0; i < UserCount; i++ {
		shopUser = append(shopUser, fmt.Sprintf("用户-%v", i))
	}
	for i := 0; i < 1; i++ {
		go qiang(rdb, shopUser[i], wg)
	}
	wg.Wait()
	fmt.Println("end")
}

var nKuCuen = 1

const (
	//初始化库存
	//商品key名字
	shangpingKey = "sku001"
	//获取锁的超时时间 秒
	timeout   = 30 * 1000
	UserCount = 10
)

//successUser := make([]string, UserCount)

func qiangdan(rdb *redis.Client, wg sync.WaitGroup) []string {
	shopUser := make([]string, 0)
	for i := 0; i < UserCount; i++ {
		shopUser = append(shopUser, fmt.Sprintf("用户-%v", i))
	}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		qiang(rdb, shopUser[i], wg)
	}
	return shopUser
}

func qiang(rdb *redis.Client, user string, wg sync.WaitGroup) {
	timeStart := time.Now()
	timeStart.Add(timeout)
	for timeStart.Add(timeout).Before(time.Now()) {
		if nKuCuen < 1 {
			return
		}
		err := rdb.SetNX(shangpingKey, user, time.Minute).Err()
		if err != nil {
			fmt.Println(err)
		}
		dingdan(rdb, user)
	}
	wg.Done()
}

func dingdan(rdb *redis.Client, user string) {
	defer rdb.Del(shangpingKey)
	defer fmt.Printf("用户%v释放锁...", user)
	fmt.Printf("用户%v拿到锁...", user)
	if nKuCuen < 1 {
		return
	}
	//  //模拟生成订单耗时操作，方便查看 多次获取锁记录
	time.Sleep(time.Second)
	nKuCuen -= 1
	//抢单成功跳出
	fmt.Printf("用户%v抢单成功跳出...所剩库存：%v", user, nKuCuen)
	return
}

func NewRedisClient(db int) *redis.Client {

	return redis.NewClient(&redis.Options{
		//Network:            "",
		Addr: "localhost:6379",
		//Dialer:             nil,
		//OnConnect:          nil,
		Password: "",
		DB:       db,
		//MaxRetries:         0,
		//MinRetryBackoff:    0,
		//MaxRetryBackoff:    0,
		//DialTimeout:        0,
		//ReadTimeout:        0,
		//WriteTimeout:       0,
		//PoolSize:           0,
		//MinIdleConns:       0,
		//MaxConnAge:         0,
		//PoolTimeout:        0,
		//IdleTimeout:        0,
		//IdleCheckFrequency: 0,
		//TLSConfig:          nil,
		//Limiter:            nil,
	})
}
