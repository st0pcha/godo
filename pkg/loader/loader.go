package loader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadGodoFile() (map[string]string, error) {
	file, err := os.Open(".godo")
	if err != nil {
		return nil, fmt.Errorf("failed to open .godo file\n%w", err)
	}
	defer file.Close()

	commands := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line in .godo file: %s", line)
		}

		command := strings.TrimSpace(parts[0])
		script := strings.TrimSpace(parts[1])
		commands[command] = script
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read .godo file\n%w", err)
	}

	return commands, nil
}