package distributedlock

//import (
//	"context"
//	"fmt"
//	"github.com/bsm/redislock"
//	"github.com/redis/go-redis/v9"
//	"log"
//	"time"
//)
///**
//https://github.com/bsm/redislock
// */
//func Lock() {
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
//	ctx := context.Background()
//
//	// Try to obtain lock.
//	lock, err := locker.Obtain(ctx, "my-key", 100*time.Millisecond, nil)
//	if err == redislock.ErrNotObtained {
//		fmt.Println("Could not obtain lock!")
//	} else if err != nil {
//		log.Fatalln(err)
//	}
//
//	// Don't forget to defer Release.
//	defer lock.Release(ctx)
//	fmt.Println("I have a lock!")
//
//	// Sleep and check the remaining TTL.
//	time.Sleep(50 * time.Millisecond)
//	if ttl, err := lock.TTL(ctx); err != nil {
//		log.Fatalln(err)
//	} else if ttl > 0 {
//		fmt.Println("Yay, I still have my lock!")
//	}
//
//	// Extend my lock.
//	if err := lock.Refresh(ctx, 100*time.Millisecond, nil); err != nil {
//		log.Fatalln(err)
//	}
//
//	// Sleep a little longer, then check.
//	time.Sleep(100 * time.Millisecond)
//	if ttl, err := lock.TTL(ctx); err != nil {
//		log.Fatalln(err)
//	} else if ttl == 0 {
//		fmt.Println("Now, my lock has expired!")
//	}
//}
