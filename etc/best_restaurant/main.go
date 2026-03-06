package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type UserRating struct {
	AverageRating float64 `json:"average_rating"`
	Votes         int     `json:"votes"`
}

type Outlet struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	City          string     `json:"city"`
	EstimatedCost int        `json:"estimated_cost"`
	UserRating    UserRating `json:"user_rating"`
}

type APIResponse struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	Total      int      `json:"total"`
	TotalPages int      `json:"total_pages"`
	Data       []Outlet `json:"data"`
}

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
	baseURL := "https://jsonmock.hackerrank.com/api/food_outlets"
	bestOutlet := Outlet{}

	page := 1
	for {
		resp, err := http.Get(fmt.Sprintf("%s?city=%s&page=%d", baseURL, city, page))
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)

		var apiResp APIResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			log.Fatal(err)
		}

		for _, outlet := range apiResp.Data {
			outletCost := outlet.EstimatedCost
			outletRating := outlet.UserRating.AverageRating

			if outletCost > int(cost) {
				continue
			}

			bestOutletRating := bestOutlet.UserRating.AverageRating
			bestOutletCost := bestOutlet.EstimatedCost

			if outletRating > bestOutletRating || (outletRating == bestOutletRating && outletCost < bestOutletCost) {
				bestOutlet = outlet
			}
		}

		if page >= apiResp.TotalPages {
			break
		}
		page++
	}

	return bestOutlet.Name
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
