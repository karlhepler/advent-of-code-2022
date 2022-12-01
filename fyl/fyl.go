package fyl

import (
	"bufio"
	"os"
)

// ReadEachLine opens a file and calls the callback function for
// each line of text. It may return an error if there is an
// issue reading the file. If the callback function returns an
// error, then it will simply pass that error up to the caller.
func ReadEachLine(filepath string, callback func(line string) error) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		err := callback(line)
		if err != nil {
			return err
		}
	}

	return nil
}
