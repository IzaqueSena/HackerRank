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
 * Complete the 'stones' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER a
 *  3. INTEGER b
 */

func updateList(list []int32, x int32) []int32 {
	var out []int32
	alreadyPut := false
	for i := 0; i < len(list); i++ {
		if list[i] >= x && !alreadyPut {
			if list[i] > x {
				out = append(out, x)
			}
			alreadyPut = true
		}
		out = append(out, list[i])
	}
	if !alreadyPut {
		out = append(out, x)
	}
	return out
}

/*
Solution description

The point is to see the problem in a manner that you can notice that the last stone is
the sum of all a's or 'bs that you chose to create the trail.

So, it is all the possible combinations with repetitions of size n-1 that you can do
using a or b

*/

//30 min

func stones(n int32, a int32, b int32) []int32 {
	var out []int32
	var sumList []int32
	for i := int32(0); i < n-1; i++ {
		sumList = append(sumList, a)
	}
	for i := int32(0); i <= n-1; i++ {
		sum := int32(0)
		if i-1 >= 0 {
			sumList[i-1] = b
		}
		for j := int32(0); j < n-1; j++ {
			sum += sumList[j]
		}
		fmt.Println(fmt.Sprintf("sumList:%v, sum:%d, out:%v", sumList, sum, out))
		out = updateList(out, sum)
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		aTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := stones(n, a, b)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
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
