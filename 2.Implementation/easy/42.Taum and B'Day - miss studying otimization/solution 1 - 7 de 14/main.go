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
 * Complete the 'taumBday' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER b
 *  2. INTEGER w
 *  3. INTEGER bc
 *  4. INTEGER wc
 *  5. INTEGER z
 */

// 13min

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	min := int64(math.MaxInt64)
	for b1 := int64(0); b1 <= int64(b); b1++ {
		for w1 := int64(0); w1 <= int64(w); w1++ {
			C := b1*int64(bc) + (int64(b)-b1)*(int64(wc)+int64(z)) + w1*int64(wc) + (int64(w)-w1)*(int64(bc)+int64(z))
			if C < min {
				min = C
			}
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

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		b := int32(bTemp)

		wTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		w := int32(wTemp)

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bcTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		bc := int32(bcTemp)

		wcTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		wc := int32(wcTemp)

		zTemp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := taumBday(b, w, bc, wc, z)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
