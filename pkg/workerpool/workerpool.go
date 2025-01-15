package workerpool

import "sync"

type WorkerPool struct {
	workerCount int
	jobChannel  chan func()
	wg          sync.WaitGroup
}

// NewWorkerPool создает новый пул воркеров
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		jobChannel:  make(chan func()),
	}
}

// Start запускает пул воркеров
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		go wp.worker(i)
	}
}

// worker выполняет задания из канала
func (wp *WorkerPool) worker(id int) {
	for job := range wp.jobChannel {
		job()
		wp.wg.Done()
	}
}

// AddJob добавляет задание в пул
func (wp *WorkerPool) AddJob(job func()) {
	wp.wg.Add(1)
	wp.jobChannel <- job
}

// Wait ожидает завершения всех заданий
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

// Stop закрывает канал
func (wp *WorkerPool) Stop() {
	close(wp.jobChannel)
}
