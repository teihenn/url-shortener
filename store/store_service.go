package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full
const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
	We want to be able to save the mapping between the originalURL

and the generated shortURL url
*/
func SaveURLMapping(shortURL string, originalURL string, userId string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf(
			"Failed saving key URL | Error: %v - shortURL: %s = originalURL: %s\n",
			err, shortURL, originalURL))
	}

}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
URL, so what we need to do here is to retrieve the long URL and
think about redirect.
*/
func RetrieveInitialURL(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf(
			"Failed RetrieveInitialURL url | Error: %v - shortURL: %s\n",
			err, shortURL))
	}
	return result
}
