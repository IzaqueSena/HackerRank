package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func strTwoToInt(s string) int64 {
	x := int64(s[0]-'0')*10 + int64(s[1]-'0')
	return x
}

func am(s string) string {
	if s[:2] != "12" {
		return s
	} else {
		return "00" + s[2:]
	}
}

func pm(s string) string {
	if s[:2] != "12" {
		x := strTwoToInt(s[:2])
		x += 12
		return fmt.Sprint(x) + s[2:]
	} else {
		return s
	}
}

func timeConversion(s string) string {
	if s[len(s)-2:] == "AM" {
		s = am(s)
		return s[:len(s)-2]
	} else {
		s = pm(s)
		return s[:len(s)-2]
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
