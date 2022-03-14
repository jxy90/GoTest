package cyclicBarrier

import (
	"context"
	"github.com/marusama/cyclicbarrier"
	"log"
	"time"
)

//https://www.cnblogs.com/softlin/p/14133635.html
func cyclicBarrierDemo() {
	barrier := cyclicbarrier.New(3)
	for i := 0; i < 3; i++ {
		go func(id int) {
			log.Printf("start: %v", id)
			barrier.Await(context.Background())
			log.Printf("finish: %v", id)
		}(i)
	}

	time.Sleep(3 * time.Second)
	log.Printf("完成")
}
