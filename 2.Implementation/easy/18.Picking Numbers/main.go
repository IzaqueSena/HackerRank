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
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

// 24min30s

func frequenciesMap(ar []int32) map[int32]int32 {
	frequencies := map[int32]int32{}
	for i := 0; i < len(ar); i++ {
		f, isInMap := frequencies[ar[i]]
		if !isInMap {
			frequencies[ar[i]] = 1
		} else {
			frequencies[ar[i]] = f + 1
		}
	}
	return frequencies
}

func pickingNumbers(a []int32) int32 {
	frequencies := frequenciesMap(a)
	maxSum := int32(0)
	for n, f := range frequencies {
		maxSecondFreq := int32(0)
		for n2, f2 := range frequencies {
			dif := n - n2
			if dif == 1 || dif == -1 {
				if f2 > maxSecondFreq {
					maxSecondFreq = f2
				}
			}
		}
		if f+maxSecondFreq > maxSum {
			maxSum = f + maxSecondFreq
		}
	}
	return maxSum
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

	result := pickingNumbers(a)

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
