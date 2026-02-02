package main

import (
    "fmt"
    "sync"
)

// Пул воркеров
type WorkerPool struct {
    tasks chan string
    wg sync.WaitGroup
}

//Создание нового пула воркеров
func NewWorkerPool(numWorkers int) *WorkerPool {
    return &WorkerPool{
        tasks: make(chan string),
		wg: sync.WaitGroup{},
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
    
    pool.Start(amount)
    
    for i := 1; i <= 100; i++ {
        pool.AddTask(fmt.Sprintf("Task data %d", i))
    }
    
    pool.Close()
}

/*
Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.
*/