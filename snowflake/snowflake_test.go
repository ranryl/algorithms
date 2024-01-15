package snowflake

import (
	"fmt"
	"testing"
	"time"
)

func TestSnowFlake(t *testing.T) {
	s := NewSnowFlake(999)
	for i := 0; i < 100; i++ {
		result, err := s.NextId()
		if err != nil {
			t.Log(err)
			continue
		}
		fmt.Printf("%d\n", result)
	}
	fmt.Println()
	result, _ := s.NextId()
	fmt.Println(result)
	time.Sleep(1 * time.Millisecond)
	result, _ = s.NextId()
	fmt.Println(result)
}
