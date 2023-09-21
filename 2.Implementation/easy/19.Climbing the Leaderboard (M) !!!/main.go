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
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

// func withoutRepetitions(ar []int32) []int32 {
//    frequencies := map[int32]int32{}
//    for i := 0; i < len(ar); i++ {
//        f, isInMap := frequencies[ar[i]]
//        if !isInMap {
//            frequencies[ar[i]] = 1
//        } else {
//            frequencies[ar[i]] = f + 1
//        }
//    }
//    var out []int32
//    for n, _ := range frequencies {
//        out = append(out, n)
//    }
//    return out
// }

func withoutRepetitions(ar []int32) []int32 {
	var out []int32
	var step int
	for i := 0; i < len(ar); i += step {
		step = 1
		isEqual := true
		for j := i + 1; isEqual && j < len(ar); j++ {
			if ar[i] != ar[j] {
				break
			}
			step++
		}
		out = append(out, ar[i])
	}
	return out
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	out := make([]int32, len(player))
	rank := withoutRepetitions(ranked)
	fmt.Println(rank)
	for i := 0; i < len(player); i++ {
		var j int
		for j = len(rank) - 1; j >= 0; j-- {
			if player[i] < rank[j] {
				break
			}
		}
		out[i] = int32(j) + 2
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ranked []int32

	for i := 0; i < int(rankedCount); i++ {
		rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
		checkError(err)
		rankedItem := int32(rankedItemTemp)
		ranked = append(ranked, rankedItem)
	}

	playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var player []int32

	for i := 0; i < int(playerCount); i++ {
		playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
		checkError(err)
		playerItem := int32(playerItemTemp)
		player = append(player, playerItem)
	}

	result := climbingLeaderboard(ranked, player)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
