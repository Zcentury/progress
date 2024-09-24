package main

import (
	"fmt"
	"github.com/Zcentury/progress"
)

func main() {
	// 创建主任务
	mainTask := progress.NewProgress("主任务", 0)

	// 创建子任务1并添加到主任务中
	subTask1 := progress.NewProgress("子任务1", 30)
	if err := mainTask.AddSubTask(subTask1); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask1.SetProgress(0)

	// 创建子任务2并添加到主任务中
	subTask2 := progress.NewProgress("子任务2", 70)
	if err := mainTask.AddSubTask(subTask2); err != nil {
		fmt.Println("错误:", err)
		return
	}

	// 创建子任务2-1并添加到子任务2中
	subTask21 := progress.NewProgress("子任务2-1", 0)
	if err := subTask2.AddSubTask(subTask21); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask21.SetProgress(50)

	// 创建子任务2-2并添加到子任务2中
	subTask22 := progress.NewProgress("子任务2-2", 0)
	if err := subTask2.AddSubTask(subTask22); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask22.SetProgress(50)

	// 创建子任务2-3并添加到子任务2中
	subTask23 := progress.NewProgress("子任务2-3", 0)
	if err := subTask2.AddSubTask(subTask23); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask23.SetProgress(50)

	// 打印可视化表格
	mainTask.DisplayProgress()

	// 计算总进度
	totalProgress, err := mainTask.CalculateTotalProgress()
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Printf("总进度: %.2f%%\n", totalProgress)
}
