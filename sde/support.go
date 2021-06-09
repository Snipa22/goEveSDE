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

type SDE struct {
	pool *pgxpool.Pool
}

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

func (s *SDE) InitSDE(pg *pgxpool.Pool) {
	s.pool = pg
}

func (s *SDE) getPoolConn() (*pgxpool.Conn, error) {
	return s.pool.Acquire(context.Background())
}

func (s *SDE) mustGetPoolConn() *pgxpool.Conn {
	conn, err := s.getPoolConn()
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
