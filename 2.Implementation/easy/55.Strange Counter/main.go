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
 * Complete the 'strangeCounter' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER t as parameter.
 */

// 26min

/*
Solution Description

model as:
n=1
i		1 2 3
value	3 2 1
(size=3)

n=2
i		4 5 6 7 8 9
value	6 5 4 3 2 1
(size=6)

n=3
i		10 11 12 13 14 15 16 17 18 19 20 21
value	12 11 10 9  8  7  6  5  4  3  2  1
(size=12)

note that
i of some element is the same as:
i = t = Soma de sizes (ate o n anterior) + (t - Soma de sizes (ate o n anterior))
Soma de sizes = 3*(2^n-1) (PG)
n = log2[(Soma de sizes + 1)/3]
se t for o ultimo do grupo se t estiver no meio, n vai ficar como o grupo anterior que ele esta


*/

func strangeCounter(t int64) int64 {
	n := int64(math.Ceil(math.Log2(float64(t)/3 + 1)))
	ini := int64(3 * int64(math.Pow(2, float64(n-1))))
	missing := t - (int64(3*(math.Pow(2, float64(n-1))-1)) + 1)
	return ini - missing
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	t, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := strangeCounter(t)

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
