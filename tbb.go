package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func floatfind(slice []float64, value float64) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

type recollection struct {
	Result []float64
}

type theta struct {
	ID     string
	Text   string
	Vector []float64
}

type phi struct {
	Token  string
	Vector []float64
}

func readTheta(file string) ([]theta, []string) {
	var topics []string
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("could not open file")
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}
	var result []theta
	for i, line := range lines {
		if i == 0 {
			for i := range line {
				if i < 3 {
					continue
				}
				topics = append(topics, line[i])
			}
			continue
		}
		identifier := line[1]
		text := line[2]
		vector := []float64{}
		for i := range line[3:len(line)] {
			index := i + 3
			floatvalue, _ := strconv.ParseFloat(line[index], 64)
			vector = append(vector, floatvalue)
		}
		result = append(result, theta{ID: identifier, Text: text, Vector: vector})
	}
	return result, topics
}

func mpair(x, y float64) float64 {
	switch {
	case y < x:
		return x - y
	case x < y:
		return y - x
	default:
		return 0
	}
}

func manhattan(x, y []float64) float64 {
	var result float64
	for i := range x {
		result = result + mpair(x[i], y[i])
	}
	return result
}

func manhattan_wghted(x, y, weight []float64) float64 {
	var result float64
	for i := range x {
		result = result + mpair(x[i], y[i])*weight[i]
	}
	return result
}

func testmanhattan(query theta, dataset []theta) []float64 {
	var result []float64
	for i := range dataset {
		result = append(result, manhattan(query.Vector, dataset[i].Vector))
	}
	return result
}

func sortresults(result []float64, number int) []float64 {
	var sorted_result []float64
	for i := range result {
		sorted_result = append(sorted_result, result[i])
	}
	sort.Float64s(sorted_result)
	sorted_result = sorted_result[0:number]
	return sorted_result
}

func reversesortresults(result []float64, number int) []float64 {
	var sorted_result []float64
	for i := range result {
		sorted_result = append(sorted_result, result[i])
	}
	sort.Float64s(sorted_result)
	for i := len(sorted_result)/2 - 1; i >= 0; i-- {
		opp := len(sorted_result) - 1 - i
		sorted_result[i], sorted_result[opp] = sorted_result[opp], sorted_result[i]
	}
	sorted_result = sorted_result[0:number]
	return sorted_result
}

func findURN(testset []theta, urn string) int {
	var result int
	for i := range testset {
		if testset[i].ID == urn {
			result = i
		}
	}
	return result
}

func main() {
	significant := float64(0.1)
	if len(os.Args) > 1 {
		significant, _ = strconv.ParseFloat(os.Args[1], 64)
	}

	fmt.Println("Reading file.")
	var number, position int
	testset, topics := readTheta("theta.csv")
	fmt.Println("All is read.")
	fmt.Println("Browsing through the space of the following topic dimensions:")
	for i := range topics {
		fmt.Println("Topic", i+1, topics[i])
	}
	fmt.Println("-------------------------------")
	fmt.Println("Significant distance has been set to: ", significant)
	fmt.Println("Happy navigating!")

	for {
		var input string
		randomid := testset[rand.Intn(len(testset))].ID

		fmt.Print("Enter URN (e.g.", randomid, "):")
		fmt.Scanln(&input)
		if input == "" {
			input = randomid
		}
		position = findURN(testset, input)
		fmt.Print("Enter number of similar passages(e.g. 3):")
		fmt.Scanln(&input)
		number, _ = strconv.Atoi(input)
		number = number + 1

		result := testmanhattan(testset[position], testset)
		sorted_result := sortresults(result, number)
		for i := range sorted_result {
			index := floatfind(result, sorted_result[i])
			fmt.Println("------------------------------------------")
			switch {
			case i == 0:
				fmt.Println("You queried:")
			case i > 0:
				fmt.Println("Result:", i)
				fmt.Println("Distance:", sorted_result[i])
			}
			fmt.Println("URN:", testset[index].ID)
			fmt.Println("Text:", testset[index].Text)
			fmt.Println()
			fmt.Println("Three most important topics:")
			sorted_indiresult := reversesortresults(testset[index].Vector, 3)
			for j := range sorted_indiresult {
				indi_index := floatfind(testset[index].Vector, sorted_indiresult[j])
				fmt.Println("Topic", indi_index+1, topics[indi_index], ":", testset[index].Vector[indi_index])
			}
			if i > 0 {
				fmt.Println()
				fmt.Println("Topics with significant distance:")
				for j := range testset[index].Vector {
					topicdistance := mpair(testset[index].Vector[j], testset[position].Vector[j])
					if topicdistance > significant {
						fmt.Println("Distance Topic", j+1, topics[j], ":", topicdistance)
					}
				}
			}
			fmt.Println("------------------------------------------")
		}
		fmt.Print("Don't want another search? Input n, otherwise just hit return:")
		fmt.Scanln(&input)
		if input == "n" {
			break
		}
	}
}
