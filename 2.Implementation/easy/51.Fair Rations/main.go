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
 * Complete the 'fairRations' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY B as parameter.
 */

// 32min

func findOdds(B []int32) []int32 {
	var odds []int32
	for i := 0; i < len(B); i++ {
		if B[i]%2 != 0 {
			odds = append(odds, int32(i))
		}
	}
	return odds
}

func addBreads(B []int32, oddIndexes []int32) []int32 {
	if oddIndexes[1] == oddIndexes[0]+1 {
		B[oddIndexes[0]]++
		B[oddIndexes[1]]++
		var newOddIndexes []int32
		for i := 2; i < len(oddIndexes); i++ {
			newOddIndexes = append(newOddIndexes, oddIndexes[i])
		}
		return newOddIndexes
	}
	B[oddIndexes[0]]++
	B[oddIndexes[0]+1]++
	oddIndexes[0]++
	return oddIndexes
}

func fairRations(B []int32) string {
	breads := int32(0)
	oddsIndexes := findOdds(B)
	if len(oddsIndexes)%2 != 0 {
		return "NO"
	}
	for len(oddsIndexes) > 0 {
		oddsIndexes = addBreads(B, oddsIndexes)
		breads += 2
	}
	return fmt.Sprintf("%d", breads)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)

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
