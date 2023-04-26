package utils

import (
	"bufio"
	"os"

	"github.com/rs/zerolog/log"
)

func ReadFile(filename *string) ([]string, error) {
	log.Debug().Msgf("file name %v", *filename)
	file, err := os.Open(*filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
