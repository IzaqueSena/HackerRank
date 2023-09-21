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
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

// 17min

func numberOfTopicsKnown(i string, j string) int32 {
	n := int32(0)
	for ind := 0; ind < len(i); ind++ {
		if i[ind] == '1' || j[ind] == '1' {
			n++
		}
	}
	return n
}

func acmTeam(topic []string) []int32 {
	out := []int32{0, 0}
	for i := 0; i < len(topic)-1; i++ {
		for j := i + 1; j < len(topic); j++ {
			n := numberOfTopicsKnown(topic[i], topic[j])
			fmt.Println(fmt.Sprintf("i=%d, j=%d, n=%d", i, j, n))
			if n > out[0] {
				out[0] = n
				out[1] = 1
			} else if n == out[0] {
				out[1]++
			}
		}
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	_ = int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
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
