package src

import (
	"strings"
)

func InsertBefore(value string, toInsert string) string {
	value = toInsert + value
	return value
}

func InsertAfter(value string, toInsert string) string {
	value = value + toInsert
	return value
}

func sanitizeLineBreaks(value string) string {
	value = strings.Replace(value, "\r\n", "", -1)
    value = strings.Replace(value, "\r", "", -1)
	value = strings.Replace(value, "\n", "", -1)
    return value
}

//process data
func Process(data string, toInsertBefore string, toInsertAfter string, excludeLastElement bool) string {
	processed_data := ""

	dataLines := strings.Split(data, "\n")
	
	for i := range dataLines {
		line := dataLines[i]
		line = sanitizeLineBreaks(line)

		if i < len(dataLines) - 1 {
			line = InsertAfter(line, toInsertAfter)
		}

		line = InsertBefore(line, toInsertBefore)

		processed_data = processed_data + line + "\n"
	}

	return processed_data
}