package fileutils

import (
	"os"
	"path/filepath"
)

func FindPath(path string, baseSearchPaths []string, filter func(info os.FileInfo) bool) string {
	if filepath.IsAbs(path) {
		if _, err := os.Stat(path); err == nil {
			return path
		}

		return ""
	}

	searchPaths := []string{}
	searchPaths = append(searchPaths, baseSearchPaths...)

	// Additionally attempt to search relative to the location of the running binary.
	var binaryDir string
	if exe, err := os.Executable(); err == nil {
		if exe, err = filepath.Abs(exe); err == nil {
			binaryDir = filepath.Dir(exe)
		}
	}

	if binaryDir != "" {
		for _, baseSearchPaths := range baseSearchPaths {
			searchPaths = append(
				searchPaths,
				filepath.Join(binaryDir, baseSearchPaths),
			)
		}
	}

	for _, parent := range searchPaths {
		found, err := filepath.Abs(filepath.Join(parent, path))
		if err != nil {
			continue
		} else if fileInfo, err := os.Stat(found); err == nil {
			if filter != nil {
				if filter(fileInfo) {
					return found
				}
			} else {
				return found
			}
		}
	}

	return ""
}

// fileutils.FindDir looks for the given directory in nearby ancestors relative to the current working
// directory as well as the directory of the executable, falling back to './' if not found.
func FindDir(dir string) (string, bool) {
	found := Find
}
