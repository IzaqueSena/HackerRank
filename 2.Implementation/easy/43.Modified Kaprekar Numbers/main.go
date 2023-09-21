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
 * Complete the 'kaprekarNumbers' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER p
 *  2. INTEGER q
 */

// 57min
// Problem: Poorly planned, resulting in a long time wasted on debugging.

func numberOfDigits(n int32) int32 {
	d := int32(0)
	for x := n; x > 0; d++ {
		x /= 10
	}
	return d
}

func kaprekar(n int32) bool {
	// fmt.Print(fmt.Sprintf("n:%d ", n))
	square := int64(n) * int64(n)
	nDigits := int64(numberOfDigits(n))
	r := int64(0)
	dec := int64(1)
	for rDigits := int64(0); rDigits < nDigits; square /= 10 {
		r += (square % 10) * dec
		rDigits++
		dec *= 10
		// fmt.Print(fmt.Sprintf("r:%d sq:%d ", r, square))
	}
	l := square
	// fmt.Println(fmt.Sprintf("r:%d l:%d", r, l))
	if r+l == int64(n) {
		return true
	} else {
		return false
	}
}

func kaprekarNumbers(p int32, q int32) {
	isInvalid := true
	for i := p; i <= q; i++ {
		if kaprekar(i) == true {
			isInvalid = false
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	if isInvalid {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	pTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	p := int32(pTemp)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	kaprekarNumbers(p, q)
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
