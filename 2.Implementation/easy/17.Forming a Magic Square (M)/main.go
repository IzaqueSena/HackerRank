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
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

// 30min+: trying solving, and I could't solve
// after studies: 7min20s

var magicSquareInstances = [][][]int32{
	{{8, 1, 6}, {3, 5, 7}, {4, 9, 2}},
	{{6, 1, 8}, {7, 5, 3}, {2, 9, 4}},
	{{4, 9, 2}, {3, 5, 7}, {8, 1, 6}},
	{{2, 9, 4}, {7, 5, 3}, {6, 1, 8}},
	{{8, 3, 4}, {1, 5, 9}, {6, 7, 2}},
	{{4, 3, 8}, {9, 5, 1}, {2, 7, 6}},
	{{6, 7, 2}, {1, 5, 9}, {8, 3, 4}},
	{{2, 7, 6}, {9, 5, 1}, {4, 3, 8}},
}

func getCost(s [][]int32, magicSquare [][]int32) int32 {
	cost := int32(0)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if magicSquare[i][j] != s[i][j] {
				if magicSquare[i][j] > s[i][j] {
					cost += magicSquare[i][j] - s[i][j]
				} else {
					cost += s[i][j] - magicSquare[i][j]
				}
			}
		}
	}
	return cost
}

func formingMagicSquare(s [][]int32) int32 {
	minCost := getCost(s, magicSquareInstances[0])
	for i := 1; i < len(magicSquareInstances); i++ {
		if cost := getCost(s, magicSquareInstances[i]); cost < minCost {
			minCost = cost
		}
	}
	return minCost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
