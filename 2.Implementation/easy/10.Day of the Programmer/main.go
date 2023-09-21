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
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func leap(year int32, isJulian bool) bool {
	if isJulian {
		if year%4 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			return true
		} else {
			return false
		}
	}
}

func dayOfProgrammer(year int32) string {
	if year == 1918 {
		fmt.Print("26.09.1918")
		return "26.09.1918"
	}
	var isJulian bool
	if year >= 1700 && year <= 1917 {
		isJulian = true
	} else {
		isJulian = false
	}
	isLeap := leap(year, isJulian)
	if isLeap {
		fmt.Printf("12.09.%s", fmt.Sprint(year))
		return fmt.Sprintf("12.09.%s", fmt.Sprint(year))
	} else {
		fmt.Printf("13.09.%s", fmt.Sprint(year))
		return fmt.Sprintf("13.09.%s", fmt.Sprint(year))
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

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
