package progress

import (
	"errors"
	"sync"
)

// Progress 结构体表示一个进度单元
type Progress struct {
	name       string      // 进度的名称
	value      float64     // 当前进度值
	percentage float64     // 占比，0表示未设置
	subTasks   []*Progress // 子任务进度
	mu         sync.RWMutex
}

// NewProgress 创建一个新的 Progress 实例
func NewProgress(name string, percentage float64) *Progress {
	return &Progress{
		name:       name,
		percentage: percentage,
	}
}

func (p *Progress) SetProgress(value float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if value > 100 {
		value = 100
	}
	p.value = value
}

// AddSubTask 添加子任务进度，并进行验证
func (p *Progress) AddSubTask(subTask *Progress) error {
	currentTotal := 0.0
	for _, task := range p.subTasks {
		if task.percentage != 0 {
			currentTotal += task.percentage
		}
	}
	if subTask.percentage != 0 {
		currentTotal += subTask.percentage
	}

	if currentTotal > 100 {
		return errors.New("子任务的占比总和超过了100%")
	}

	p.subTasks = append(p.subTasks, subTask)
	return nil
}

// CalculateTotalProgress 计算总进度
func (p *Progress) CalculateTotalProgress() (float64, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// 如果没有子任务
	if len(p.subTasks) == 0 {
		return p.value, nil
	}

	totalProgress := 0.0 // 初始化总进度为0

	// 遍历所有子任务
	for _, subTask := range p.subTasks {

		// 递归计算子任务的进度
		subProgress, err := subTask.CalculateTotalProgress()
		if err != nil {
			// 如果计算子任务进度时出错，返回错误
			return 0, err
		}
		if subTask.percentage != 0 {
			// 如果子任务设置了占比，将子任务的进度乘以该占比并加到总进度中
			totalProgress += subProgress * (subTask.percentage / 100)
		} else {
			// 如果子任务设置了占比，将子任务的进度乘以该占比并加到总进度中
			totalProgress += subProgress * ((100.0 / float64(len(p.subTasks))) / 100)
		}
	}

	// 返回计算得到的总进度
	return totalProgress, nil
}
