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
 * Complete the 'jumpingOnClouds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY c as parameter.
 */

//24min

func jumpClouds(c []int32, currentPosition int32, currentJumps int32, jumps []int32) []int32 {
	if currentPosition == int32(len(c))-1 {
		jumps = append(jumps, currentJumps)
		return jumps
	}
	if currentPosition+2 <= int32(len(c))-1 && c[currentPosition+2] == 0 {
		jumps = jumpClouds(c, currentPosition+2, currentJumps+1, jumps)
	}
	if c[currentPosition+1] == 0 {
		jumps = jumpClouds(c, currentPosition+1, currentJumps+1, jumps)
	}
	return jumps
}

func jumpingOnClouds(c []int32) int32 {
	var jumps []int32
	jumps = jumpClouds(c, 0, 0, jumps)
	fmt.Println(jumps)
	min := int32(math.MaxInt32)
	for i := 0; i < len(jumps); i++ {
		if jumps[i] < min {
			min = jumps[i]
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

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var c []int32

	for i := 0; i < int(n); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int32(cItemTemp)
		c = append(c, cItem)
	}

	result := jumpingOnClouds(c)

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
