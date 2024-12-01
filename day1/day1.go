package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list1 := []int{}
	list2 := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	var total = 0
	slices.Sort(list1)
	slices.Sort(list2)
	for i := 0; i < len(list1); i++ {
		fmt.Println(list1[i], list2[i])
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println("Total:", total)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
	}
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	list1 := []int{}
	list2 := []int{}
	list2Map := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, "   ")
		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])
		if value, exists := list2Map[value2]; exists {
			list2Map[value2] = value + 1
		} else {
			list2Map[value2] = 1
		}
		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	var total = 0
	for i := 0; i < len(list1); i++ {
		fmt.Println(list1[i], list2[i])
		value1 := list1[i]
		if value, exists := list2Map[value1]; exists {
			total += value1 * value
		}

	}
	fmt.Println("Total:", total)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
	}
}

func main() {

	part1()
	part2()
}
