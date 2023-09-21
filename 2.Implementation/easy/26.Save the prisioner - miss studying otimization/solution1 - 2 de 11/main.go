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
 * Complete the 'saveThePrisoner' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER s
 */

// 9min

func saveThePrisoner(n int32, m int32, s int32) int32 {
    count := int32(0)
    var i int32
    for i = s; i <= n; i++ {
        count++
        if count == m {
            break
        }
        if i == n {
            i = 0
        }
    }
    return i

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        m := int32(mTemp)

        sTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
        checkError(err)
        s := int32(sTemp)

        result := saveThePrisoner(n, m, s)

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

