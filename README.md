# progress
一个方便快捷的进度计算工具



# 函数

### func NewProgress(name string, percentage float64) *Progress

创建一个进度

- **name**：任务名称
- **percentage**：进度占比（0为自动分配），一个任务下的子任务占比不能超过100



# 方法

### func AddSubTask(subTask *Progress) error

添加子任务



### func SetProgress(value float64)

设置进度



# 演示

```go
func main() {
	// 创建主任务
	mainTask := progress.NewProgress("主任务", 0)

	// 创建子任务并添加到主任务中
	subTask1 := progress.NewProgress("子任务 1", 30)
	if err := mainTask.AddSubTask(subTask1); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask1.SetProgress(0)

	// 创建子任务并添加到主任务中
	subTask2 := progress.NewProgress("子任务 2", 70)
	if err := mainTask.AddSubTask(subTask2); err != nil {
		fmt.Println("错误:", err)
		return
	}

	// 创建子任务并添加到子任务 2中
	subTask3 := progress.NewProgress("子任务 2-1", 0)
	if err := subTask2.AddSubTask(subTask3); err != nil {
		fmt.Println("错误:", err)
		return
	}

	subTask3.SetProgress(50)

	// 创建子任务并添加到子任务 2中
	subTask4 := progress.NewProgress("子任务 2-2", 0)
	if err := subTask2.AddSubTask(subTask4); err != nil {
		fmt.Println("错误:", err)
		return
	}
	subTask4.SetProgress(50)

	// 创建子任务并添加到子任务 2中
	subTask5 := progress.NewProgress("子任务 2-3", 0)
	if err := subTask2.AddSubTask(subTask5); err != nil {
		fmt.Println("错误:", err)
		return
	}

	subTask5.SetProgress(50)

	// 计算总进度
	totalProgress, err := mainTask.CalculateTotalProgress()
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	fmt.Printf("总进度: %.2f%%\n", totalProgress)
}
```



> 图示

```
                    +-------35--------+
                    |      主任务      |
                    +-----------------+
                              |
        +---------------------+---------------------+
        |                                           |
+-------0---------+                         +-------50--------+
|  子任务1-占比30%  |                         |  子任务2-占比70%  |
+-----------------+                         +-----------------+
                                                     |
                            +------------------------+----------------------+
                            |                        |                      |
                    +-------50--------+     +--------50-------+     +-------50--------+
                    | 子任务2-1-占比33% |     | 子任务2-2-占比33% |     | 子任务2-3-占比33% |
                    +-----------------+     +-----------------+     +-----------------+
```

