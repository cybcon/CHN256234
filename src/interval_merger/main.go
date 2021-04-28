/**
 * This solution is part of a task from #CHN256234.
 * Author: Michael Oberdorf <michael.oberdorf@bridging-it>
 * Date: 2021-04-26
 **/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

const VERSION = "1.0.1"

/**
 * Interval
 * Structure of 2 integers with the names "Lo" and "Hi" that defines an interval
 **/
type Interval struct {
	Lo, Hi int
}

/**
 * parseJson
 * @desc: parses an interval json string of format "[[a, b],[c, d],...]"
 * @param s: string, the string to parse
 * @return: array of struct Interval, the merged intervals
 **/
func parseJson(b *[]byte) []Interval {
	var jsonIntervals [][]int

	// we unmarshal our byteArray which contains our
	byteValue := *b
	json.Unmarshal(byteValue, &jsonIntervals)

	var inputIntervals []Interval

	for i := 0; i < len(jsonIntervals); i++ {
		if jsonIntervals[i][0] <= jsonIntervals[i][1] {
			inputIntervals = append(inputIntervals, Interval{Lo: jsonIntervals[i][0], Hi: jsonIntervals[i][1]})
		} else {
			inputIntervals = append(inputIntervals, Interval{Lo: jsonIntervals[i][1], Hi: jsonIntervals[i][0]})
		}
	}

	return inputIntervals
}

/**
 * merge
 * @desc: the merge function will get the different intervals and merge overlapping ones
 * @param ivs: array of struct Interval, the given intervals
 * @return: array of struct Interval, the merged intervals
 * @see: https://stackoverflow.com/questions/55201821/merging-overlapping-intervals-using-double-for-loop
 **/
func merge(ivs []Interval) []Interval {
	// define m as the merged array, build from an empty array of struct Interval and add the given intervals
	m := append([]Interval(nil), ivs...)
	// if the array has no or just one entry, return the intervals
	if len(m) <= 1 {
		return m
	}

	// sort the given intervals by minimum value, using an unnamed function
	sort.Slice(m,
		func(i, j int) bool {
			if m[i].Lo < m[j].Lo {
				return true
			}
			if m[i].Lo == m[j].Lo && m[i].Hi < m[j].Hi {
				return true
			}
			return false
		},
	)

	// loop over intervals and merge the intervals by inspect the current and the next interval in the array
	j := 0
	for i := 1; i < len(m); i++ {
		if m[j].Hi >= m[i].Lo {
			if m[j].Hi < m[i].Hi {
				m[j].Hi = m[i].Hi
			}
		} else {
			j++
			m[j] = m[i]
		}
	}

	// return the merged intervals
	return append([]Interval(nil), m[:j+1]...)
}

/**
 * outputMergedJson
 * @desc: the outputMergedJson function will transform the array of intervals into a json string and outputs that
 * @param m: array of struct Interval, the given array of intervals
 **/
func outputMergedJson(m []Interval) {
	fmt.Print("[")
	for i := 0; i < len(m); i++ {
		if i == 0 {
			fmt.Print("[", m[i].Lo, ",", m[i].Hi, "]")
		} else {
			fmt.Print(",[", m[i].Lo, ",", m[i].Hi, "]")
		}
	}
	fmt.Println("]")
}




// main function
func main() {
	// prepare an input filename
	filename := flag.String("file", "", "The json file that contains the intervals, e.g. \"--file=intervals.json\"")
	// prepare the input parameter -i as a string
	var intervals string
	flag.StringVar(&intervals, "i", "", "A json string that contains the intervals, e.g. -i \"[[25,30],[2,19],[14, 23],[4,8]]\"")
	flag.Parse()

	if *filename != "" && intervals != "" {
		fmt.Println("ERROR: concurent input, please specify interval (-i=\"<intervals>\") or an input file (-file=<file.json>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *filename == "" && intervals == "" {
		fmt.Println("ERROR: missing input, please specify interval (-i=\"<intervals>\") or an input file (-file=<file.json>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// create the byteValue out of the input parameters
	var byteValue []byte
	if intervals != "" {
		byteValue = []byte(intervals)
	}
	if *filename != "" {
		// read the json file
		jsonFile, err := os.Open(*filename)
		// handle errors during file open
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// read our opened jsonFile as a byte array.
		byteValue, _ = ioutil.ReadAll(jsonFile)
	}

	// array of struct Intervals that will have the validated and parsed input params
	var inputIntervals []Interval
	inputIntervals = parseJson(&byteValue)

	var mergedIntervals []Interval
	mergedIntervals = merge(inputIntervals)

	outputMergedJson(mergedIntervals)
}
