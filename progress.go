package progress

import (
	"errors"
	"github.com/olekukonko/tablewriter"
	"os"
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

	if len(p.subTasks) == 0 {
		return p.value, nil
	}

	totalProgress := 0.0

	for _, subTask := range p.subTasks {
		subProgress, err := subTask.CalculateTotalProgress()
		if err != nil {
			return 0, err
		}
		if subTask.percentage != 0 {
			totalProgress += subProgress * (subTask.percentage / 100)
		} else {
			totalProgress += subProgress * ((100.0 / float64(len(p.subTasks))) / 100)
		}
	}

	return totalProgress, nil
}

// DisplayProgress 显示进度的表格
func (p *Progress) DisplayProgress() {
	data := CollectData(p)
	normalizedData := NormalizeRows(data)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(normalizedData)
	table.Render()
}
