package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumDistances' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */
func indexes(a []int32) map[int32][]int32 {
	indexesMap := map[int32][]int32{}
	for i := int32(0); i < int32(len(a)); i++ {
		indexes, isInMap := indexesMap[a[i]]
		if !isInMap {
			indexesMap[a[i]] = []int32{i}
		} else {
			indexes = append(indexes, i)
			indexesMap[a[i]] = indexes
		}
	}
	return indexesMap
}

func minimunDistance(a []int32) int32 {
	min := int32(math.MaxInt32)
	for i := 1; i < len(a); i++ {
		d := a[i] - a[i-1]
		if d < min {
			min = d
		}
	}
	return min
}

// 20min

func minimumDistances(a []int32) int32 {
	min := int32(math.MaxInt32)
	indexesMap := indexes(a)
	fmt.Println(indexesMap)
	for _, indexes := range indexesMap {
		if len(indexes) > 1 {
			d := minimunDistance(indexes)
			if d < min {
				min = d
			}
		}
	}
	if min == int32(math.MaxInt32) {
		min = int32(-1)
	}
	return min
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

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := minimumDistances(a)

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
