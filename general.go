package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func stringIsExistsInSlice(s string, slice []string) bool {
	for i := range slice {
		if slice[i] == s {
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
	// If there is empty element in slices remove it
	var should_delete_indicies []int
	for i := range Dares {
		if strings.TrimSpace(Dares[i]) == "" {
			should_delete_indicies = append(should_delete_indicies, i)
		}
	}
	for i := range should_delete_indicies {
		Dares = removeFromSlice(Dares, should_delete_indicies[i])
	}
	should_delete_indicies = make([]int, 0)

	// Repeat operations for truth slice
	for i := range Truths {
		if strings.TrimSpace(Truths[i]) == "" {
			should_delete_indicies = append(should_delete_indicies, i)
		}
	}
	for i := range should_delete_indicies {
		Truths = removeFromSlice(Truths, should_delete_indicies[i])
	}
}
func removeFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
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
func readCheatingUsernames() {
	f_bytes, err := ioutil.ReadFile("./cheating.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data map[string]interface{}
	err = json.Unmarshal(f_bytes, &data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	untyped_cheat_time, ok := data["MAX_CHEAT_TIME"]
	if !ok {
		fmt.Println("There's no CHEAT_TIME field in cheating.json file")
	} else {
		cheat_time_float, ok := untyped_cheat_time.(float64)
		if !ok {
			fmt.Println("Unknown error occurred during convert CHEAT_TIME to intger")
		} else {
			MAX_CHEAT_TIME = int(cheat_time_float)
		}
	}

	untyped_good_usernames, ok := data["GOOD_USERNAMES"]
	if !ok {
		fmt.Println("There's no GOOD_USERNAMES field in cheating.json file")
	} else {
		good_usernames_interface_slice, ok := untyped_good_usernames.([]interface{})
		if !ok {
			fmt.Println("Error occurred during convert GOOD_USERNAMES to interface slice")
		} else {
			var good_usernames_slice = make([]string, len(good_usernames_interface_slice))
			for i, v := range good_usernames_interface_slice {
				good_usernames_slice[i] = fmt.Sprint(v)
			}
			GOOD_USERNAMES = good_usernames_slice
		}
	}
}
