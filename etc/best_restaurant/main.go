package main

import (
	"bufio"
	_ "encoding/json"
	"fmt"
	"io"
	_ "io/ioutil"
	_ "log"
	_ "net/http"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'bestRestaurant' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING city
 *  2. INTEGER cost
 * API URL: https://jsonmock.hackerrank.com/api/food_outlets?city={city}&page={page}
 */

func bestRestaurant(city string, cost int32) string {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	city := readLine(reader)

	costTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	cost := int32(costTemp)

	result := bestRestaurant(city, cost)

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
