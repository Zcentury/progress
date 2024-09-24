package progress

import "fmt"

func NormalizeRows(rows [][]string) [][]string {
	maxColumns := 0

	for _, row := range rows {
		if len(row) > maxColumns {
			maxColumns = len(row)
		}
	}

	for i := range rows {
		for len(rows[i]) < maxColumns {
			rows[i] = append(rows[i], "")
		}
	}

	return rows
}

func CollectData(p *Progress) [][]string {
	var result [][]string

	for _, subTask := range p.subTasks {
		progress, _ := p.CalculateTotalProgress()
		subTaskProgress, _ := subTask.CalculateTotalProgress()

		currentNames := []string{fmt.Sprintf("%s[%.0f]%.2f%%", p.name, p.percentage, progress), fmt.Sprintf("%s[%.0f]%.2f%%", subTask.name, subTask.percentage, subTaskProgress)}

		if len(subTask.subTasks) > 0 {
			subData := collectSubData(subTask, currentNames)
			result = append(result, subData...)
		} else {
			// 如果是最终子任务
			result = append(result, currentNames)
		}
	}

	if len(p.subTasks) == 0 {
		progress, _ := p.CalculateTotalProgress()
		row := []string{fmt.Sprintf("%s[%.0f]%.2f%%", p.name, p.percentage, progress)}
		result = append(result, row)
	}

	return result
}

// 辅助函数，用于递归处理子任务
func collectSubData(p *Progress, parentNames []string) [][]string {
	var result [][]string

	for _, subTask := range p.subTasks {
		currentNames := append([]string{}, parentNames...)
		progress, _ := subTask.CalculateTotalProgress()
		currentNames = append(currentNames, fmt.Sprintf("%s[%.0f]%.2f%%", subTask.name, subTask.percentage, progress))

		if len(subTask.subTasks) > 0 {
			subData := collectSubData(subTask, currentNames)
			result = append(result, subData...)
		} else {
			result = append(result, currentNames)
		}
	}

	if len(p.subTasks) == 0 {
		row := append([]string{}, parentNames...)
		progress, _ := p.CalculateTotalProgress()
		row = append(row, fmt.Sprintf("%s[%.0f]%.2f%%", p.name, p.percentage, progress))
		result = append(result, row)
	}

	return result
}
