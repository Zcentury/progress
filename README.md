### progress
一个方便快捷的进度计算工具



### API

### 函数



#### func NewProgress(name string, percentage float64) *Progress

创建一个进度

- **name**：任务名称
- **percentage**：进度占比（0为自动分配），一个任务下的子任务占比不能超过100



### 方法

#### func (*Progress)AddSubTask(subTask *Progress) error

添加子任务



#### func (*Progress)SetProgress(value float64)

设置进度



#### func (*Progress)CalculateTotalProgress()

计算进度



#### func (*Progress)DisplayProgress()

打印可视化图表

```
+-----------------+-------------------+--------------------+
| 主任务[0]35.00% | 子任务1[30]0.00%  |                    |
+                 +-------------------+--------------------+
|                 | 子任务2[70]50.00% | 子任务2-1[0]50.00% |
+                 +                   +--------------------+
|                 |                   | 子任务2-2[0]50.00% |
+                 +                   +--------------------+
|                 |                   | 子任务2-3[0]50.00% |
+-----------------+-------------------+--------------------+
```

### 演示

```go
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
```



> 图示

此图示非图表输出函数结果，而是为了方便大家理解计算

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
