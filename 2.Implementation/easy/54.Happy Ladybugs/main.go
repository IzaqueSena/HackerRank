package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'happyLadybugs' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING b as parameter.
 */

func frequencies(s string) map[string]int {
	freqMap := map[string]int{}
	for i := 0; i < len(s); i++ {
		f, isInMap := freqMap[string(s[i])]
		if !isInMap {
			freqMap[string(s[i])] = 1
		} else {
			freqMap[string(s[i])] = f + 1
		}
	}
	return freqMap
}

func isAllLadybugsHappy(b string) bool {
	var queue []string
	for i := 0; i < len(b); i++ {
		if len(queue) > 0 {
			if string(b[i]) != queue[len(queue)-1] {
				return false
			} else {
				if i == len(b)-1 || string(b[i+1]) != queue[len(queue)-1] {
					queue = []string{}
				}
			}
		} else {
			queue = append(queue, string(b[i]))
		}
	}
	if len(queue) > 0 {
		return false
	}
	return true
}

func happyLadybugs(b string) string {
	freqMap := frequencies(b)
	_, isThereUnderline := freqMap["_"]
	if !isThereUnderline {
		if isAllLadybugsHappy(b) {
			return "YES"
		} else {
			return "NO"
		}
	} else {
		for key, f := range freqMap {
			if key != string('_') && f == 1 {
				return "NO"
			}
		}
		return "YES"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		_ = int32(nTemp)

		b := readLine(reader)

		result := happyLadybugs(b)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
