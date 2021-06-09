package sde

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/patrickmn/go-cache"
	"log"
	"runtime"
	"time"
)

// Define the local LRU cache, default caching is 30 minutes, with scans every 15
var c = cache.New(30*time.Minute, 15*time.Minute)
var pool *pgxpool.Pool

func getItemFromCache(cacheExtension string) (interface{}, bool) {
	pc, _, _, _ := runtime.Caller(1)
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v_%v", runtime.FuncForPC(pc).Name(), cacheExtension)))
	return c.Get(string(h.Sum(nil)))
}

func setItemInCache(cacheExtension string, cacheData interface{}) {
	pc, _, _, _ := runtime.Caller(1)
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v_%v", runtime.FuncForPC(pc).Name(), cacheExtension)))
	c.Set(string(h.Sum(nil)), cacheData, cache.DefaultExpiration)
}

func InitSDE(pg *pgxpool.Pool) {
	pool = pg
}

func getPoolConn() (*pgxpool.Conn, error) {
	return pool.Acquire(context.Background())
}

func mustGetPoolConn() *pgxpool.Conn {
	conn, err := getPoolConn()
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
