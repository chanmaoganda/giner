package giner

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const DATABASE_URL string = "DATABASE_URL"
const REDIS_URL string = "REDIS_URL"

var (
	envs map[string]string
	/// make sure envOnce only be loaded once
	envOnce sync.Once
)

func LoadEnv() {
	envOnce.Do(func() {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Cannot find .env file")
		}
	
		databaseUrl := os.Getenv(DATABASE_URL)
	
		if databaseUrl == "" {
			log.Fatal("Cannot load databaseUrl from .env")
		}
	
		redisUrl := os.Getenv(REDIS_URL)
		
		if redisUrl == "" {
			log.Fatal("Cannot load redisUrl from .env")
		}
	
		envs[DATABASE_URL] = databaseUrl
		envs[REDIS_URL] = redisUrl
	})
}