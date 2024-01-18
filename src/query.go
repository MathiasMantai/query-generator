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
func Process(data string, query string, toInsertBefore string, toInsertAfter string, excludeLastElement bool, useAsIn bool, replaceDoubleQuotes bool) string {
	processed_data := ""

	dataLines := strings.Split(data, "\n")
	
	if !useAsIn {
		for i := range dataLines {
			if replaceDoubleQuotes {
				dataLines[i] = strings.Replace(dataLines[i], "\"", "'", -1)
			}

			line := dataLines[i]
			line = sanitizeLineBreaks(line)

			if i < len(dataLines) - 1 {
				line = InsertAfter(line, toInsertAfter)
			}

			line = InsertBefore(line, toInsertBefore)

			processed_data = processed_data + strings.Replace(query, "?", line, 1) + "\n"
		}
	} else {
		//prepare data for in clause
		lines := ""
		for i := range dataLines {
            line := dataLines[i]
            line = sanitizeLineBreaks(line)

			if replaceDoubleQuotes {
				line = strings.Replace(line, "\"", "'", -1)
			}

            line = InsertAfter(line, toInsertAfter)

			if excludeLastElement && i < len(dataLines) - 1 {
				line = line[:len(line) - 2]
			}


			line = InsertBefore(line, toInsertBefore)
			lines = lines + line
			if i < len(dataLines) - 1 {
				lines = lines + ", "
			}
		}
		
		processed_data = strings.Replace(query, "?", lines, 1)
	}

	return processed_data
}