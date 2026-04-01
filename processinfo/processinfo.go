package processinfo

import (
	"bufio"
	"strings"

	"github.com/probeldev/aerospace-autokill/bash"
)

type ProcessInfo struct {
	PID     string
	CPU     string
	Memory  string
	Command string
	WorkDir string
	Start   string
	Time    string
}

func GetProcessByName(
	name string,
) (
	[]ProcessInfo,
	error,
) {

	command := `ps aux | grep ` + name + ` | grep -v grep`

	output, err := bash.RunCommand(command)

	if err != nil {
		return nil, err
	}

	process := parseProcesses(output)

	return process, nil
}

func parseProcesses(output string) []ProcessInfo {
	var processes []ProcessInfo

	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Разделяем по пробелам, сохраняя команду с пробелами
		fields := strings.Fields(line)
		if len(fields) < 11 {
			continue
		}

		// Структура ps aux на macOS:
		// USER PID %CPU %MEM VSZ RSS TT STAT STARTED TIME COMMAND
		proc := ProcessInfo{
			PID:    fields[1],
			CPU:    fields[2],
			Memory: fields[3],
			Start:  fields[8],
			Time:   fields[9],
		}

		// Команда - это всё, что после TIME (индекс 10)
		if len(fields) > 10 {
			proc.Command = strings.Join(fields[10:], " ")

			// Извлекаем флаг -w если есть
			if idx := strings.Index(proc.Command, "-w "); idx != -1 {
				rest := proc.Command[idx+3:]
				// Ищем путь (он может быть в кавычках или до следующего пробела)
				if len(rest) > 0 {
					if rest[0] == '"' {
						// Путь в кавычках
						endQuote := strings.Index(rest[1:], `"`)
						if endQuote != -1 {
							proc.WorkDir = rest[1 : endQuote+1]
						}
					} else {
						// Путь без кавычек
						pathParts := strings.Fields(rest)
						if len(pathParts) > 0 {
							proc.WorkDir = pathParts[0]
						}
					}
				}
			}
		}

		processes = append(processes, proc)
	}

	return processes
}
