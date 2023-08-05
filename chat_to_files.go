package main

import (
	"fmt"
	"github.com/go-pax/gpt-engineer/database"
	"github.com/go-pax/gpt-engineer/io"
	"regexp"
)

type fileItem struct {
	path string
	code string
}

func parseChat(chat string) []fileItem {
	// Get the line above markdown
	regex := regexp.MustCompile("([^\n]+)\n\\s*```[^\n]*\n([\\s\\S.]+?)```")
	matches := regex.FindAllStringSubmatch(chat, -1)

	files := make([]fileItem, 0, len(matches))

	for _, match := range matches {
		path := match[1]

		regex = regexp.MustCompile(`([\w-]+\\)*?[\w-]+\.\w+`)
		path = regex.FindString(path)

		// Get the code
		code := match[2]

		// Add the file to the list
		files = append(files, fileItem{path, code})
	}

	return files
}

func toFiles(chat string, workspace database.Database) {
	if err := workspace.Set("all_output.txt", chat); err != nil {
		panic(err)
	}

	files := parseChat(chat)
	for _, file := range files {
		if file.path == "" {
			file.path = io.GenerateRandomFilename()
		}
		if err := workspace.Set(file.path, file.code); err != nil {
			fmt.Println("error, possibly openai returned invalid filename.")
			panic(err)
		}
	}
}
