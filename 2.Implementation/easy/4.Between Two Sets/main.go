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
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func gcd(a, b int32) int32 {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func commonDivisors(l []int32) ([]int32, int32) {
	gcdNumber := l[0]
	for i := 1; i < len(l); i++ {
		gcdNumber = gcd(gcdNumber, l[i])
	}
	var commonDivisorsList []int32
	for i := int32(1); i <= gcdNumber; i++ {
		if gcdNumber%i == 0 {
			commonDivisorsList = append(commonDivisorsList, i)
		}
	}
	return commonDivisorsList, gcdNumber
}

func lcm(a, b int32) int32 {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / gcd(a, b)
}

func getTotalX(a []int32, b []int32) int32 {
	commonDivisorsB, gcdB := commonDivisors(b)
	fmt.Println(fmt.Sprintf("gcdB: ", gcdB))
	fmt.Println(fmt.Sprintf("commonDivisorsB: ", commonDivisorsB))
	if len(a) == 0 || len(b) == 0 {
		return int32(0)
	}
	lcmA := a[0]
	for i := 1; i < len(a); i++ {
		lcmA = lcm(lcmA, a[i])
	}
	fmt.Println(fmt.Sprintf("lcmA: ", lcmA))
	count := int32(0)
	var outputList []int32
	for i := int32(0); i*lcmA <= gcdB; i++ {
		for j := 0; j < len(commonDivisorsB); j++ {
			if i*lcmA == commonDivisorsB[j] {
				count++
				outputList = append(outputList, i*lcmA)
			}
		}
	}
	fmt.Println(fmt.Sprint("a: ", a))
	fmt.Println(fmt.Sprint("b: ", b))
	fmt.Println(outputList)
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

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
