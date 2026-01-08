package concurrence_test

import (
	"sync"
	"testing"

	"github.com/chenyang-zz/go-learn/basic/concurrence"
)

func TestSimpleGoroutine(t *testing.T) {
	concurrence.SimpleGoroutine()
}

func TestWaitGroup(t *testing.T) {
	// concurrence.SubRoutine()
	concurrence.WaitGroup()
}

func TestLock(t *testing.T) {
	// concurrence.Atomic()
	// concurrence.Lock()
	// concurrence.ReentranceRLock(3)
	// concurrence.ReentranceWLock(3)
	// concurrence.RLockExclusion()
	concurrence.WLockExclusion()
}

func TestConcurrentMap(t *testing.T) {
	cm1 := concurrence.NewConcurrentMap[string, int](50)
	cm1.Store("张三", 18)
	if v, exists := cm1.Load("张三"); !exists {
		t.Fail()
	} else {
		if v != 18 {
			t.Fail()
		}
	}
	if _, exists := cm1.Load("李四"); exists {
		t.Fail()
	}

	cm2 := concurrence.NewConcurrentMap[int, bool](50)
	cm2.Store(18, true)
	if v, exists := cm2.Load(18); !exists {
		t.Fail()
	} else {
		if v != true {
			t.Fail()
		}
	}
	if _, exists := cm2.Load(19); exists {
		t.Fail()
	}

	// 测试在高并发情况下是否安全
	const P = 10
	var wg sync.WaitGroup
	for i := 0; i < P; i++ {
		wg.Go(func() {
			for j := 0; j < 1000; j++ {
				cm2.Store(j, true)
				cm2.Load(j)
			}
		})
	}
	wg.Wait()
}

func TestBroadcast(t *testing.T) {
	concurrence.Broadcast()
}

func TestCountDownLatch(t *testing.T) {
	concurrence.CountDownLatch()
}

func TestCondSignal(t *testing.T) {
	concurrence.CondSignal()
}

func TestChannelSignal(t *testing.T) {
	concurrence.ChannelSignal()
}

func TestCondBroadcast(t *testing.T) {
	concurrence.CondBroadcast()
}

func TestChannelBroadcast(t *testing.T) {
	concurrence.ChannelBroadcast()
}

func TestDealMassFile(t *testing.T) {
	concurrence.DealMassFile("../data/biz_log")
}

func TestQpsLimit(t *testing.T) {
	concurrence.QpsLimit()
}

func TestRoutineLimit(t *testing.T) {
	concurrence.RoutineLimit()
}

func TestListenMultiWay(t *testing.T) {
	concurrence.ListenMultiWay()
}

func TestSelectBlock(t *testing.T) {
	concurrence.SelectBlock()
}
