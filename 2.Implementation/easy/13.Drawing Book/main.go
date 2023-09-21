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
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

// 23min30

func countFront(n int32, p int32) int32 {
	pages := []int32{0, 1}
	out := int32(0)
	for pages[1] <= n {
		if pages[0] == p || pages[1] == p {
			return out
		}
		out++
		pages[0] += 2
		pages[1] += 2
	}
	return out
}

func countBack(n int32, p int32) int32 {
	var pages []int32
	if n%2 == 0 {
		pages = []int32{n, n + 1}
	} else {
		pages = []int32{n - 1, n}
	}
	out := int32(0)
	for pages[0] >= 0 {
		if pages[0] == p || pages[1] == p {
			return out
		}
		pages[0] -= 2
		pages[1] -= 2
		out++
	}
	return out
}

func pageCount(n int32, p int32) int32 {
	turnsFront := countFront(n, p)
	turnsBack := countBack(n, p)
	if turnsFront < turnsBack {
		return turnsFront
	} else {
		return turnsBack
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

	fmt.Fprintf(writer, "%d\n", result)

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
