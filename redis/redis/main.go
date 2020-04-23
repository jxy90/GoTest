package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//err := client.Set("k1", "v1", time.Second).Err()
	//if err != nil {
	//	fmt.Println(err)
	//}

	val, err := client.Get("k1").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
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

func main() {
	rdb := NewRedisClient(1)
	defer rdb.Close()

	//HyperLogLog(rdb)
	//ZSet(rdb)
	//PipeLine(rdb)

}

func HyperLogLog(rdb *redis.Client) {
	err := rdb.PFAdd("PF1", "Redis", "MongoDB", "MySQL").Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.PFCount("PF1").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("HyperLogLog:%v \r\n", val)
}

func ZSet(rdb *redis.Client) {
	err := rdb.ZAdd("zset:k1", &redis.Z{
		Score:  60,
		Member: "v1",
	}, &redis.Z{
		Score:  100,
		Member: "v5",
	}, &redis.Z{
		Score:  70,
		Member: "v2",
	}, &redis.Z{
		Score:  80,
		Member: "v3",
	}, &redis.Z{
		Score:  90,
		Member: "v4",
	}).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.Get("zset:k1").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("zset set:k1:%v \r\n", val)
	val, err = rdb.Get("zset:not exit").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("zset set:not exit:%v \r\n", val)
}

func Set(rdb *redis.Client) {
	err := rdb.SAdd("set:ABC", "A", "B", "C", "D", "E", "F").Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.SPop("set:ABC").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("set set:k1:%v \r\n", val)
	val, err = rdb.SPop("set:not exit").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("pop set:not exit:%v \r\n", val)
}

func String(rdb *redis.Client) {
	err := rdb.Set("string:k1", "v1", time.Hour).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := rdb.Get("string:k1").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("get string:k1:%v \r\n", val)
	val, err = rdb.Get("string:not exit").Result()
	if err != redis.Nil {
		fmt.Println(err)
	} else if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("set string:not exit:%v \r\n", val)
}

func PipeLine(rdb *redis.Client) {
	pipeline := rdb.TxPipeline()
	defer pipeline.Close()
	for i := 0; i < 1000; i++ {
		pipeline.Set(fmt.Sprintf("k%v", i), fmt.Sprintf("v%v", i), time.Second*30)
	}
	pipeline.Exec()

}

func Hash(rdb *redis.Client) {
	err := rdb.HMSet("website", "baidu", "www.baidu.com", "youtube", "www.youtube.com", "a", "www.a.com").Err()
	if err != nil {
		fmt.Println(err)
	}
	result, err := rdb.HGetAll("website").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
