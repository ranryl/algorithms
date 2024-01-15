package snowflake

import (
	"errors"
	"time"
)

const (
	STARTTIMESTR = "2024-01-15 00:00:00 000"
	CUSTOMRFC    = "2006-01-02 15:04:05 000"
	// TIMEBIT      = 41
	ZONENODEBIT = 10
	INSTANCEBIT = 12
)

type SnowFlake struct {
	lastTimeMill int64
	startTime    int64
	zoneNodeID   int64
	instanceSize int64
}

func NewSnowFlake(zoneNodeID int64) *SnowFlake {
	startTime, _ := time.Parse(CUSTOMRFC, STARTTIMESTR)
	return &SnowFlake{
		lastTimeMill: time.Now().UnixMilli(),
		startTime:    startTime.UnixMilli(),
		zoneNodeID:   zoneNodeID,
		instanceSize: 1,
	}
}
func (s *SnowFlake) NextId() (int64, error) {
	now := time.Now().UnixMilli()
	if now < s.lastTimeMill {
		return 0, errors.New("time is changed, now timestamp < last timestamp")
	}
	if now == s.lastTimeMill {
		if s.instanceSize%(1<<INSTANCEBIT) == 0 {
			time.Sleep(1 * time.Millisecond)
			now += 1
			s.instanceSize = 1
		}
	}
	result := (now-s.startTime)<<(ZONENODEBIT+INSTANCEBIT) + s.zoneNodeID<<ZONENODEBIT + s.instanceSize
	s.instanceSize++
	return result, nil
}
