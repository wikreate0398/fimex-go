package lifecycle

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type LifecycleManager struct {
	instances []Lifecycle
}

func (lm *LifecycleManager) Run() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := new(sync.WaitGroup)

	for _, instance := range lm.instances {
		wg.Add(1)

		go func() {
			defer wg.Done()
			instance.append.OnStart(ctx)
		}()
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-shutdown:

		fmt.Println("Shutting down...")
		for _, instance := range lm.instances {
			instance.append.OnStop(ctx)
		}

		fmt.Println("Cancel ctx...")
		cancel()
		close(shutdown)
		//default:
		//	close(lm.shutdown)
	}

	// Ждём завершения горутин
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("All goroutines finished gracefully")
	case <-time.After(10 * time.Second):
		fmt.Println("Forcing shutdown...")
	}
}

type Lifecycle struct {
	append *AppendLifecycle
}

type AppendLifecycle struct {
	OnStart func(ctx context.Context) interface{}
	OnStop  func(ctx context.Context) interface{}
}

func (l *Lifecycle) Append(append AppendLifecycle) {
	l.append = &append
}

func Register(instances ...func(*Lifecycle)) *LifecycleManager {
	var items []Lifecycle
	for _, instance := range instances {
		lf := &Lifecycle{}
		instance(lf)
		items = append(items, *lf)
	}

	return &LifecycleManager{instances: items}
}
