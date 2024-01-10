package bloom

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type BloomDistributed struct {
	key      string
	Len      uint64
	hashList []BloomFilterHash
	rdb      *redis.Client
}

func NewBloomDistributed(key string, lens uint64, hashFuncs []BloomFilterHash) (*BloomDistributed, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	val, err := rdb.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	// fmt.Printf("%b\n", []byte(val))
	if val == "" {
		initVal := make([]byte, lens/8+1)
		fmt.Println(initVal)
		err := rdb.Set(ctx, key, string(initVal), 0).Err()
		if err != nil {
			return nil, err
		}
	}
	return &BloomDistributed{
		key:      key,
		hashList: hashFuncs,
		Len:      lens,
		rdb:      rdb,
	}, nil
}

func (b *BloomDistributed) Add(data []byte) {
	for _, hashFunc := range b.hashList {
		hashValue := hashFunc(data) % b.Len
		// fmt.Printf("data=%s, index=%d\n", string(data), hashValue)
		err := b.rdb.SetBit(context.Background(), b.key, int64(hashValue), 1).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (b *BloomDistributed) IsExsit(data []byte) bool {
	for _, hashFunc := range b.hashList {
		hashValue := hashFunc(data) % b.Len
		if val, _ := b.rdb.GetBit(context.Background(), b.key, int64(hashValue)).Result(); val == 0 {
			fmt.Println(val)
			return false
		}
	}
	return true
}
