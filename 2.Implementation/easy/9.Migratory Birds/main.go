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
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

// 19min

func migratoryBirds(arr []int32) int32 {
	mapNumberFrequency := map[int32]int32{}
	for i := 0; i < len(arr); i++ {
		f, isInMap := mapNumberFrequency[arr[i]]
		if !isInMap {
			mapNumberFrequency[arr[i]] = 1
		} else {
			mapNumberFrequency[arr[i]] = f + 1
		}
	}
	fmt.Println(fmt.Sprint("map: ", mapNumberFrequency))
	var maximuns []int32
	max := int32(0)
	for n, f := range mapNumberFrequency {
		if f > max {
			maximuns = []int32{n}
			max = f
		} else if f == max {
			maximuns = append(maximuns, n)
		}
	}
	fmt.Println(fmt.Sprint("maximuns: ", maximuns))
	min := maximuns[0]
	for i := 0; i < len(maximuns); i++ {
		if maximuns[i] < min {
			min = maximuns[i]
		}
	}
	return min

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

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
