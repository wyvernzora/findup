package findup

import (
	"errors"
	"os"
	"path/filepath"
)

/*
Find() recursively searches ancestor directories of current working directory looking for {filename}.
Returns the path of the first {filename} found; returns error if reaches file system root without finding anything.
*/
func Find(filename string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return FindInDir(cwd, filename)
}

/*
FindInDir() does the same thing that Find() does, but allows user to specify what directory to start from.
*/
func FindInDir(path string, filename string) (string, error) {
	current, _ := filepath.Abs(path)

	for {
		matches, err := filepath.Glob(filepath.Join(current, filename))

		if err != nil {
			return "", err
		}

		if len(matches) > 0 {
			return matches[0], nil
		}

		// Give up if reached file system root
		if current == "/" {
			break
		}

		current = filepath.Join(current, "..")
	}

	return "", errors.New("no matching files found")
}
