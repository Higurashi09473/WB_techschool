package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Пул воркеров
type WorkerPool struct {
	tasks chan string
	wg    sync.WaitGroup
}

// Создание нового пула воркеров
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks: make(chan string),
		wg:    sync.WaitGroup{},
	}
}

// Запуск пула воркеров
func (wp *WorkerPool) Start(numWorkers int) {
	for i := 0; i <= numWorkers; i++ {
		wp.wg.Add(1)
		go func(workerID int) {
			defer wp.wg.Done()
			for task := range wp.tasks {
				fmt.Printf("Worker %d processing task: %s\n", workerID, task)
			}
		}(i)
	}
}

// Запуск добавление задачи воркерам
func (wp *WorkerPool) AddTask(task string) {
	wp.tasks <- task
}

// Закрытие канала с ожиданием завершения работы горутин воркеров
func (wp *WorkerPool) Close() {
	close(wp.tasks)
	wp.wg.Wait()
}

func main() {
	var amount int
	fmt.Scan(&amount)
	pool := NewWorkerPool(amount)

	//Инициализируем контест чтобы в будущем по нему завершить горутину
	ctx, cancel := context.WithCancel(context.Background())

	//Инициализируем канал signalCh для отлова syscall
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	pool.Start(amount)

	wgMain := &sync.WaitGroup{}
	wgMain.Add(1)
	go func(ctx context.Context, wgwgMain *sync.WaitGroup) {
		var i int
		for {
			select {
				case <-ctx.Done():
					wgMain.Done()
					return
				default:
					pool.AddTask(fmt.Sprintf("Task data %d", i))
					i++
			}
		}
	}(ctx, wgMain)

	<-signalCh
	
	//завершаем горутину по контексту
	cancel()

	wgMain.Wait()

	pool.Close()
}

/*
Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.
*/
