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
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY stringList
 *  2. STRING_ARRAY queries
 */

// 3min

func matchingStrings(stringList []string, queries []string) []int32 {
	var count []int32
	for i := 0; i < len(queries); i++ {
		count = append(count, 0)
	}
	for i := 0; i < len(queries); i++ {
		for j := 0; j < len(stringList); j++ {
			if queries[i] == stringList[j] {
				count[i]++
			}
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stringListCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var stringList []string

	for i := 0; i < int(stringListCount); i++ {
		stringListItem := readLine(reader)
		stringList = append(stringList, stringListItem)
	}

	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []string

	for i := 0; i < int(queriesCount); i++ {
		queriesItem := readLine(reader)
		queries = append(queries, queriesItem)
	}

	res := matchingStrings(stringList, queries)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
