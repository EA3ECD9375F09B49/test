/**
 * @Author michael
 * @Description
 * @Date 23:22 2021/2/23
 * @Param
 * @return
 **/
package libs

import (
	"github.com/RussellLuo/timingwheel"
	"sync"
	"time"
)

var once sync.Once

func init() {
	once.Do(func() {
		TimingWheel.Start()
	})
}

var TimingWheel = timingwheel.NewTimingWheel(time.Millisecond, 20)

type RotateScheduler struct {
	Interval time.Duration
}

func (s *RotateScheduler) Next(prev time.Time) time.Time {
	return prev.Add(s.Interval)
}
