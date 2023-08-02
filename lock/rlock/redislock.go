package rlock

// redis分布式锁
// https://github.com/bsm/redislock
// https://github.com/go-redsync/redsync

//import (
//	"fmt"
//	"log"
//	"time"
//
//	"github.com/bsm/redislock"
//	"github.com/redis/go-redis/v9"
//)
//
//func main() {
//	// Connect to redis.
//	client := redis.NewClient(&redis.Options{
//		Network: "tcp",
//		Addr:    "127.0.0.1:6379",
//	})
//	defer client.Close()
//
//	// Create a new lock client.
//	locker := redislock.New(client)
//
//	// Try to obtain lock.
//	lock, err := locker.Obtain("my-key", 100*time.Millisecond, nil)
//	if err == redislock.ErrNotObtained {
//		fmt.Println("Could not obtain lock!")
//	} else if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Don't forget to defer Release.
//	defer lock.Release()
//	fmt.Println("I have a lock!")
//
//	// Sleep and check the remaining TTL.
//	time.Sleep(50 * time.Millisecond)
//	if ttl, err := lock.TTL(); err != nil {
//		log.Fatalln(err)
//	} else if ttl > 0 {
//		fmt.Println("Yay, I still have my lock!")
//	}
//
//	// Extend my lock.
//	if err := lock.Refresh(100*time.Millisecond, nil); err != nil {
//		log.Fatalln(err)
//	}
//
//	// Sleep a little longer, then check.
//	time.Sleep(100 * time.Millisecond)
//	if ttl, err := lock.TTL(); err != nil {
//		log.Fatalln(err)
//	} else if ttl == 0 {
//		fmt.Println("Now, my lock has expired!")
//	}
//
//}
