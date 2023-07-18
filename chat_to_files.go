package main

import (
	"regexp"
)

type fileItem struct {
	path string
	code string
}

func parseChat(chat string) []fileItem {
	regex := regexp.MustCompile("(\\S+)\n\\s*```[^\n]*\n([\\s\\S.]+?)```")
	matches := regex.FindAllStringSubmatch(chat, -1)

	files := make([]fileItem, 0, len(matches))

	for _, match := range matches {
		path := match[1]

		// Strip the filename of any non-allowed characters and convert / to \
		regex = regexp.MustCompile(`[<>"|?*]`)
		path = regex.ReplaceAllString(path, "")

		// Remove leading and trailing brackets
		regex = regexp.MustCompile(`^\[(.*)\]$`)
		path = regex.ReplaceAllString(path, "\\1")

		// Remove leading and trailing backticks
		regex = regexp.MustCompile("^`(.*)`$")
		path = regex.ReplaceAllString(path, "\\1")

		// Remove trailing ]
		regex = regexp.MustCompile("\\]$")
		path = regex.ReplaceAllString(path, "")

		// Get the code
		code := match[2]

		// Add the file to the list
		files = append(files, fileItem{path, code})
	}

	return files
}

func toFiles(chat string, workspace *DB) {
	workspace.Set("all_output.txt", chat)

	files := parseChat(chat)
	for _, file := range files {
		workspace.Set(file.path, file.code)
	}
}
