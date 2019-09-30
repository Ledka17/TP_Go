package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type ByColumn struct{
	isNum bool
	colNum int
	data []string
}
func (a *ByColumn) Len() int           { return len(a.data) }
func (a *ByColumn) Swap(i, j int)      { a.data[i], a.data[j] = a.data[j], a.data[i] }
func (a *ByColumn) Less(i, j int) bool {
	if a.isNum {
		aInt := ByInt(a.data)
		return aInt.Less(i, j)
	}
	return a.data[i] < a.data[j]
}

type ByInt []string
func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInt) Less(i, j int) bool {
	iInt, err1 := atoi(a[i])
	jInt, err2 := atoi(a[i])
	if (err1 != nil) && (err2 != nil) {
		panic(err1)
		return a[i] < a[j]
	}
	return iInt < jInt
}

type ByStr []string
func (a ByStr) Len() int           { return len(a) }
func (a ByStr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStr) Less(i, j int) bool { return a[i] < a[j] }

type IgnoreCase []string
func (a IgnoreCase) Len() int           { return len(a) }
func (a IgnoreCase) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IgnoreCase) Less(i, j int) bool { return strings.ToLower(a[i]) < strings.ToLower(a[j]) }

type ByAsc []string
func (a ByAsc) Len() int           { return len(a) }
func (a ByAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAsc) Less(i, j int) bool { return a[i] < a[j] }

func atoi(element string) (int, error) {
	elementInt, err := strconv.Atoi(element)
	if err != nil {
		return elementInt, err
	}
	return elementInt, nil
}

func unique(arrayStr []string) []string {
	keys := make(map[string]bool)
	var arrayUnique []string

	for _, element := range arrayStr {
		if _, value := keys[element]; !value {
			keys[element] = true
			arrayUnique = append(arrayUnique, element)
		}
	}
	return arrayUnique
}

func readFile(fileName string) ([]string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, err
}

func printFile(data []string, fileOutput string) error {
	file, err := os.Create(fileOutput)
	if err != nil {
		return err
	}

	for line := range data {
		_, _ = file.WriteString(data[line] + "\n")
	}
	defer file.Close()
	return err
}

func main() {

	// flags
	ignoreCase := flag.Bool("f", false, "ignore case")
	uniqueSort := flag.Bool("u", false, "unique words")
	reverseSort := flag.Bool("r", false, "reverse sort")
	numberSort := flag.Bool("n", false, "numbers sort")
	fileOutput := flag.String("o", "stdout", "output file")

	flag.Parse()

	// read file
	fileName := os.Args[len(os.Args) - 1]
	data, err := readFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	//data = []string{"apple", "abc", "napkin", "book"}
	//data = []string{"78", "66", "105", "40", "9", "9"}

	// sort
	dataInt := []int{}
	if *numberSort {
		dataInt, err = atoi(data)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sort.Ints(dataInt)
	}

	if *ignoreCase {
		sort.Sort(IgnoreCase(data))
	}

	if *uniqueSort {
		data = unique(data)
		sort.Strings(data)
	}

	if *reverseSort {
		sort.Sort(sort.Reverse(ByAsc(data)))
	} else {
		sort.Strings(data)
	}

	// print result
	if *fileOutput != "stdout" {
		err = printFile(data, *fileOutput)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		if len(dataInt) > 0 {
			fmt.Println(dataInt)
		} else {
			fmt.Println(data)
		}
	}
}