package main

import (
	"fmt"
	"os"
	"time"

	obs "github.com/aprksy/bricks/base/pattern/observer"
)

type TimeProvider struct {
	obs.SimpleSubject[uint, time.Time]
}

type TimeDisplay struct {
	obs.SimpleObserver[uint, time.Time]
}

func print(key string, value time.Time) {
	fmt.Printf("%s: %v\n", key, value)
}

func main() {
	tsubjmgr := obs.NewSubjectManager[uint]()
	tprov := &TimeProvider{
		SimpleSubject: *obs.NewSimpleSubject[uint, time.Time](1, "time", time.Now()),
	}
	obs.AddSubjects[uint, time.Time](tsubjmgr, tprov)

	tdisp := &TimeDisplay{
		SimpleObserver: *obs.NewSimpleObserverWithSubjectManager[uint, time.Time](2, print, tsubjmgr),
	}
	tdisp.SubscribeByKey("time")

	ticker := time.NewTicker(time.Second)
	endTime := time.After(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			obs.Inject[uint, time.Time](tsubjmgr, "time", time.Now())
		case <-endTime:
			os.Exit(0)
		}
	}
}
