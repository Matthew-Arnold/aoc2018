package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	data := make(map[int][][60]bool)
	records := make([]string, 0)

	var currentGuard int
	awake := true
	startMinute := 0
	var night [60]bool

	for stdin.Scan() {
		line := stdin.Text()
		records = append(records, line)
	}

	sort.Strings(records)

	for _, record := range records {
		minute, _ := strconv.Atoi(record[15:17])
		event := record[19:]
		//fmt.Println(minute, ":", event)

		//resolve previous state
		if !awake {
			for i := startMinute; i < minute; i++ {
				night[i] = true
			}
		}

		startMinute = minute
		if strings.Contains(event, "begins shift") {
			data[currentGuard] = append(data[currentGuard], night)
			night = [60]bool{}

			guardID, _ := strconv.Atoi(strings.Split(event, " ")[1][1:])
			currentGuard = guardID
		} else if event == "falls asleep" {
			awake = false
		} else if event == "wakes up" {
			awake = true
		} else {
			fmt.Fprintln(os.Stderr, "Error: invalid event: ", event)
		}
	}

	data[currentGuard] = append(data[currentGuard], night)

	maxGuardID := 0
	maxMinute := 0
	maxMinuteCount := 0

	for guardID, nights := range data {
		guardMaxMinute := 0
		guardMaxMinuteCount := 0

		for i := 0; i < 60; i++ {
			minuteCount := 0

			for _, night := range nights {
				if night[i] {
					minuteCount++
				}
			}

			if minuteCount > guardMaxMinuteCount {
				guardMaxMinute = i
				guardMaxMinuteCount = minuteCount
			}
		}

		//fmt.Println(guardID, guardMaxMinute, guardMaxMinuteCount)

		if guardMaxMinuteCount > maxMinuteCount {
			maxGuardID = guardID
			maxMinute = guardMaxMinute
			maxMinuteCount = guardMaxMinuteCount
		}
	}

	//fmt.Println(maxGuardID, maxMinute)
	fmt.Println(maxMinute * maxGuardID)
}
