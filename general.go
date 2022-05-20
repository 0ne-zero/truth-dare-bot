package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func stringIsExistsInSlice(s string, slice []string) bool {
	for i := range slice {
		if s == slice[i] {
			return true
		}
	}
	return false
}

func readFileAsString(path string) (string, error) {
	file_bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file_bytes), nil
}

func loadTruthsAndDares() {
	dare_texts, err := readFileAsString("questions/dare.txt")
	if err != nil {
		os.Exit(1)
	}
	splitted_dare_texts := strings.Split(dare_texts, "|")
	Dares = append(Dares, splitted_dare_texts...)

	truth_texts, err := readFileAsString("questions/truth.txt")
	if err != nil {
		os.Exit(1)
	}
	splitted_truth_texts := strings.Split(truth_texts, "|")
	Truths = append(Truths, splitted_truth_texts...)
}
func breakStringSliceInLines(slice []string) string {
	var result string
	s_len := len(slice)
	for i := range slice {
		if i == s_len {
			result += slice[i]
		} else {
			result += slice[i] + "\n"
		}
	}
	return result
}
