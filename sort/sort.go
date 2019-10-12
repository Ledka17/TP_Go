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

func takeColumn(a ByColumn) ByColumn {
	aColumn := a
	if *ProgFlags.columnSort > 1 {
		for index := 1; index < *ProgFlags.columnSort; index++ {
			for elem, _ := range aColumn {
				aColumn[elem] = aColumn[elem][strings.Index(aColumn[elem], " ") + 1:]
			}
		}
	}
	return aColumn
}

type ByColumn []string
func (a ByColumn) Len() int           { return len(a) }
func (a ByColumn) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByColumn) Less(i, j int) bool {
	aColumn := takeColumn(a)
	if *ProgFlags.numberSort {
		return ByInt(aColumn).Less(i, j)
	}
	return aColumn[i] < aColumn[j]
}

type ByInt []string
func (a ByInt) Len() int           { return len(a) }
func (a ByInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByInt) Less(i, j int) bool {
	iInt, err1 := atoI(a[i])
	jInt, err2 := atoI(a[j])
	if (err1 != nil) && (err2 != nil) {
		panic(err1)
	}
	return iInt < jInt
}

type ByStr []string
func (a ByStr) Len() int           { return len(a) }
func (a ByStr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStr) Less(i, j int) bool { return a[i] < a[j] }

type flags struct {
	ignoreCase *bool
	uniqueSort *bool
	reverseSort *bool
	numberSort *bool
	columnSort *int
	fileOutput *string
}

var ProgFlags flags

func main() {

	// flags
	ProgFlags.ignoreCase = flag.Bool("f", false, "ignore case")
	ProgFlags.uniqueSort = flag.Bool("u", false, "unique words")
	ProgFlags.reverseSort = flag.Bool("r", false, "reverse sort")
	ProgFlags.numberSort = flag.Bool("n", false, "numbers sort")
	ProgFlags.columnSort = flag.Int( "k", 1, "sort by column")
	ProgFlags.fileOutput = flag.String("o", "stdout", "output file")

	flag.Parse()

	// read file
	fileName := os.Args[len(os.Args) - 1]
	data, err := readFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sort
	data = MySort(data)

	// print result
	if *ProgFlags.fileOutput != "stdout" {
		err = printFile(data, *ProgFlags.fileOutput)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println(data)
	}
}

func MySort(data []string) []string {
	if *ProgFlags.uniqueSort {
		data = unique(data, *ProgFlags.ignoreCase)
	}

	sort.Sort(ByStr(data))
	if *ProgFlags.numberSort {
		sort.Sort(ByInt(data))
	} else if *ProgFlags.columnSort > 1 {
		sort.Sort(ByColumn(data))
	}

	if *ProgFlags.reverseSort {
		j := len(data) - 1
		for i := 0; i < j; i++ {
			data[i], data[j] = data[j], data[i]
			j--
		}
	}
	return data
}

func atoI(element string) (int, error) {
	element = element + " "
	element = element[:strings.Index(element, " ")]
	elementInt, err := strconv.Atoi(element)
	if err != nil {
		return elementInt, err
	}
	return elementInt, nil
}

func unique(arrayStr []string, ignoreCase bool) []string {
	keys := make(map[string]bool)
	var arrayUnique []string

	if !ignoreCase {
		for _, element := range arrayStr {
			if _, value := keys[element]; !value {
				keys[element] = true
				arrayUnique = append(arrayUnique, element)
			}
		}
	} else {
		for _, element := range arrayStr {
			if _, value := keys[strings.ToLower(element)]; !value {
				keys[strings.ToLower(element)] = true
				arrayUnique = append(arrayUnique, element)
			}
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
