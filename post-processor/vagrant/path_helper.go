package vagrant

import (
	"path/filepath"
	"strings"
)

func getDirectoryPaths(filePaths []string) []string {
	directoryPaths := make([]string, len(filePaths))
	for i, path := range filePaths {
		directoryPaths[i] = filepath.Dir(path)
	}
	return directoryPaths
}

func alignLength(pathsSegments [][]string, length int) {
	for i := 0; i < len(pathsSegments); i++ {
		alignedPathSegments := make([]string, length)
		copy(alignedPathSegments, pathsSegments[i][0:length])
		pathsSegments[i] = alignedPathSegments
	}
}

func splitToSameLengthPathSegments(directoryPaths []string, seperator string) [][]string {
	pathsSegments := make([][]string, len(directoryPaths))
	minLen := -1

	for i, dirPath := range directoryPaths {
		pathSegments := strings.Split(dirPath, seperator)
		if minLen == -1 || len(pathSegments) < minLen {
			minLen = len(pathSegments)
		}
		pathsSegments[i] = pathSegments
	}

	alignLength(pathsSegments, minLen)
	return pathsSegments
}

// Return path to a directory that is common to all paths if one exists or an empty string if not.
func getCommonDirectory(filePaths []string, pathSeperator string) string {
	directoriesPaths := getDirectoryPaths(filePaths)
	switch len(filePaths) {
	case 0:
		return ""
	case 1:
		return directoriesPaths[0]
	}

	directoryPathsSegments := splitToSameLengthPathSegments(directoriesPaths, pathSeperator)
	commonPathSegments := directoryPathsSegments[0]

	for i := 1; i < len(directoryPathsSegments); i++ {
		currentDirPathSegments := directoryPathsSegments[i][0:]
		for j := 0; j < len(currentDirPathSegments); j++ {
			if len(commonPathSegments) <= j || currentDirPathSegments[j] == commonPathSegments[j] {
				continue
			} else {
				commonPathSegments = commonPathSegments[0:j]
			}
		}
	}

	return strings.Join(commonPathSegments, pathSeperator)
}
