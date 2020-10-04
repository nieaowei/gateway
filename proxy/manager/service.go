package manager

import (
	"context"
	"fmt"
	"gateway/lib"
	"github.com/go-redis/redis/v8"
	"log"
	"sync/atomic"
	"time"
)

type RedisFlowCountService struct {
	ticker      *time.Ticker
	AppID       string
	QPS         int64
	Unix        int64
	TickerCount int64
	TotalCount  int64
	notify      chan int64
}

func (o *RedisFlowCountService) ServiceName() string {
	return o.AppID
}

func (o *RedisFlowCountService) Stop() {
	o.notify <- 2
}

func (o *RedisFlowCountService) Exec() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%v", err)
			}
		}()
		atomic.AddInt64(&o.TickerCount, 1)
		if o.ticker == nil {
			o.notify <- 1
		}
		//fmt.Println(o)
		//data, _ := lib.DefaultRedisCluster().Get(context.Background(), o.GetDayKey(time.Now())).Int64()
	}()
}

// 定时上传任务
func (o *RedisFlowCountService) Start() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		// New ticker
		//ticker := time.NewTicker(o.Interval)
		for true {
			// wait for ticker channel
			//<-o.ticker.C
			if o.ticker != nil {
				<-o.ticker.C
			} else {
				data := <-o.notify
				if data == 2 {
					break
				}
			}
			// start statistic
			// read data
			tickerCount := atomic.LoadInt64(&o.TickerCount)
			// reset data
			atomic.StoreInt64(&o.TickerCount, 0)
			currentTime := time.Now() // 当前时间

			dayKey := o.GetDayKey(currentTime)   // 日Key
			hourKey := o.GetHourKey(currentTime) // 时Key
			// redis 事务
			_, err := lib.DefaultRedisCluster().Pipelined(context.Background(), func(p redis.Pipeliner) error {
				_, err := p.IncrBy(context.Background(), dayKey, tickerCount).Result()
				if err != nil {
					return err
				}
				_, err = p.Expire(context.Background(), dayKey, 86400*2*time.Millisecond).Result()
				if err != nil {
					return err
				}
				_, err = p.IncrBy(context.Background(), hourKey, tickerCount).Result()
				if err != nil {
					return err
				}
				_, err = p.Expire(context.Background(), hourKey, 86400*2*time.Millisecond).Result()
				return err
			})
			if err != nil {
				log.Printf("Redis write error %v", err.Error())
				continue
			}
			total, err := o.GetDayData(currentTime)
			if err != nil {
				log.Printf("Redis write error %v", err.Error())
				continue
			}
			nowUnix := time.Now().Unix()
			if o.Unix == 0 {
				o.Unix = time.Now().Unix()
				continue
			}
			// Computing the qps
			tickerCount = total - o.TotalCount
			if nowUnix > o.Unix {
				o.TotalCount = total
				o.QPS = tickerCount / (nowUnix - o.Unix)
				o.Unix = time.Now().Unix()
			}
			//log.Printf(" [INFO] Service: %v , Count: %v ,QPS: %v \n", o.AppID, o.TotalCount, o.QPS)
		}
	}()
}

var TimeLocation *time.Location

func init() {
	var err error
	TimeLocation, err = time.LoadLocation(lib.GetDefaultConfProxy().Base.TimeLocation)
	if err != nil {
		log.Fatal(err)
	}
}

const (
	RedisFlowDayKey  = "flow_day_count"
	RedisFlowHourKey = "flow_hour_count"
	HourFormat       = "2006010215"
	DayFormat        = "20060102"
)

func (o *RedisFlowCountService) GetDayKey(t time.Time) string {
	dayStr := t.In(TimeLocation).Format(DayFormat)
	return fmt.Sprintf("%s_%s_%s", RedisFlowDayKey, dayStr, o.ServiceName())
}

func (o *RedisFlowCountService) GetHourKey(t time.Time) string {
	hourStr := t.In(TimeLocation).Format(HourFormat)
	return fmt.Sprintf("%s_%s_%s", RedisFlowHourKey, hourStr, o.ServiceName())
}

func (o *RedisFlowCountService) GetHourData(t time.Time) (int64, error) {
	return lib.DefaultRedisCluster().Get(context.Background(), o.GetHourKey(t)).Int64()
}

func (o *RedisFlowCountService) GetDayData(t time.Time) (int64, error) {
	return lib.DefaultRedisCluster().Get(context.Background(), o.GetDayKey(t)).Int64()
}

// NewRedisFlowCountService is used by statistic.
// If interval == 0 , used by computing query number per second.
// Else it used by recording query number every day.
func NewRedisFlowCountService(appID string, interval time.Duration) *RedisFlowCountService {
	reqCounter := &RedisFlowCountService{
		AppID: appID,
		//ticker: time.NewTicker(interval),
		QPS:    0,
		Unix:   0,
		notify: make(chan int64),
	}
	if interval != 0 {
		reqCounter.ticker = time.NewTicker(interval)
	}
	return reqCounter
}