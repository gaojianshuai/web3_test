package main

import (
	"fmt"
	"sync"
	"time"
)

// 复杂了些
// Task 定义任务接口
type Task struct {
	Name     string
	Function func() error
}

// TaskResult 任务执行结果
type TaskResult struct {
	TaskName  string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks     []Task
	results   []TaskResult
	wg        sync.WaitGroup
	mu        sync.Mutex
	startTime time.Time
	endTime   time.Time
}

// NewScheduler 创建新的调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(name string, taskFunc func() error) {
	s.tasks = append(s.tasks, Task{
		Name:     name,
		Function: taskFunc,
	})
}

// Execute 并发执行所有任务
func (s *Scheduler) Execute() {
	s.startTime = time.Now()

	for _, task := range s.tasks {
		s.wg.Add(1)
		go s.executeTask(task)
	}

	s.wg.Wait()
	s.endTime = time.Now()
}

// executeTask 执行单个任务
func (s *Scheduler) executeTask(task Task) {
	defer s.wg.Done()

	start := time.Now()
	err := task.Function()
	end := time.Now()

	result := TaskResult{
		TaskName:  task.Name,
		StartTime: start,
		EndTime:   end,
		Duration:  end.Sub(start),
		Error:     err,
	}

	s.mu.Lock()
	s.results = append(s.results, result)
	s.mu.Unlock()
}

// PrintResults 打印执行结果
func (s *Scheduler) PrintResults() {
	fmt.Printf("\n=== 任务执行结果统计 ===\n")
	fmt.Printf("总任务数: %d\n", len(s.tasks))
	fmt.Printf("总执行时间: %v\n", s.endTime.Sub(s.startTime))
	fmt.Printf("开始时间: %v\n", s.startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("结束时间: %v\n", s.endTime.Format("2006-01-02 15:04:05"))

	fmt.Printf("\n=== 详细任务执行时间 ===\n")
	for _, result := range s.results {
		status := "成功"
		if result.Error != nil {
			status = fmt.Sprintf("失败: %v", result.Error)
		}
		fmt.Printf("任务: %-15s | 耗时: %-12v | 状态: %s\n",
			result.TaskName, result.Duration, status)
	}
}

// GetResults 获取所有任务结果
func (s *Scheduler) GetResults() []TaskResult {
	return s.results
}

// 示例任务函数
func sampleTask1() error {
	time.Sleep(2 * time.Second)
	fmt.Println("任务1执行完成")
	return nil
}

func sampleTask2() error {
	time.Sleep(1 * time.Second)
	fmt.Println("任务2执行完成")
	return nil
}

func sampleTask3() error {
	time.Sleep(3 * time.Second)
	fmt.Println("任务3执行完成")
	return nil
}

func failingTask() error {
	time.Sleep(500 * time.Millisecond)
	return fmt.Errorf("模拟任务失败")
}

func main() {
	// 创建调度器实例
	scheduler := NewScheduler()

	// 添加示例任务
	scheduler.AddTask("数据处理任务", sampleTask1)
	scheduler.AddTask("文件读取任务", sampleTask2)
	scheduler.AddTask("网络请求任务", sampleTask3)
	scheduler.AddTask("失败测试任务", failingTask)

	fmt.Println("开始并发执行任务...")

	// 执行所有任务
	scheduler.Execute()

	// 打印执行结果
	scheduler.PrintResults()

	// 计算统计信息
	results := scheduler.GetResults()
	var totalDuration time.Duration
	var successCount int

	for _, result := range results {
		totalDuration += result.Duration
		if result.Error == nil {
			successCount++
		}
	}

	fmt.Printf("\n=== 性能统计 ===\n")
	fmt.Printf("任务成功率: %.1f%%\n", float64(successCount)/float64(len(results))*100)
	fmt.Printf("平均任务耗时: %v\n", totalDuration/time.Duration(len(results)))
}
