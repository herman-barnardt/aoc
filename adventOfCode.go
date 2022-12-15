package aoc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
	"time"
)

type solution struct {
	Part1 func([]string) interface{}
	Part2 func([]string) interface{}
}

var allSolutions = make(map[int]map[int]solution)

var filenameIncludesYear = false
var fileExtension = ".go"

func Register(year int, dayNumber int, a, b func([]string) interface{}) {
	if _, ok := allSolutions[year]; !ok {
		allSolutions[year] = make(map[int]solution)
	}
	allSolutions[year][dayNumber] = solution{
		Part1: a,
		Part2: b,
	}
}

func SetFilenameIncludesYear(b bool) {
	filenameIncludesYear = b
}

func SetFileExtension(s string) {
	fileExtension = s
}

func SetSolutionTemplate(s string) {
	solutionTemplate = s
}

func Run(cmd string, year int, day int, part int) error {
	switch cmd {
	case "create":
		return create(year, day)
	case "download":
		return download(year, day)
	case "solve":
		return solve(fmt.Sprintf("./data/%d/day%d", year, day), year, day, part)
	case "test":
		return solve(fmt.Sprintf("./data/%d/day%dTest", year, day), year, day, part)
	default:
		return errors.New(fmt.Sprintf("unsupported command: %q", cmd))

	}
}

func makeDir(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func getFile(year, day int) {
	cookie, _ := readLines("session")
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie[0]})

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create(fmt.Sprintf("./data/%d/day%d", year, day))
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.Write(body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func readLines(file string) ([]string, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	values := make([]string, 0)

	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return values, nil
}

func solve(path string, year, day, part int) error {
	if _, ok := allSolutions[year][day]; !ok {
		err := create(year, day)
		if err != nil {
			return fmt.Errorf("Solution not implemented\n%v", err)
		}
		return errors.New("Solution not implemented\nTemplate has been generated")
	}

	lines, err := readLines(path)

	if err != nil || len(lines) == 0 {
		return fmt.Errorf("Missing input file: %s", path)
	}

	if part == 1 || part == 0 {
		startTime := time.Now()
		solution1 := allSolutions[year][day].Part1(lines)
		duration1 := time.Since(startTime)
		fmt.Printf("Part 1\n%v (%v)\n", solution1, duration1)
	}
	if part == 2 || part == 0 {
		startTime := time.Now()
		solution2 := allSolutions[year][day].Part2(lines)
		duration2 := time.Since(startTime)
		fmt.Printf("Part 2\n%v (%v)\n", solution2, duration2)
	}
	return nil
}

func download(year, day int) error {
	makeDir("./data")

	path := fmt.Sprintf("./data/%d/day%d", year, day)
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Downloading file")
		makeDir(fmt.Sprintf("./data/%d", year))
		getFile(year, day)
	} else {
		return err
	}
	return nil
}

func create(year, day int) error {
	filename := fmt.Sprintf("day%d", day)

	if filenameIncludesYear {
		filename = fmt.Sprintf("%dday%d", year, day)
	}

	basepath, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := path.Join(basepath, filename)

	_, err = os.Stat(filePath)
	if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("file=%s already exists (%v)", filename, err)
	}

	filePath += fileExtension

	templateString := solutionTemplate

	data := map[string]interface{}{
		"Year": year,
		"Day":  day,
	}

	t := template.Must(template.New("template").Parse(templateString))

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return fmt.Errorf("Failed to populate solution template %v", err)
	}

	templateString = buf.String()

	return os.WriteFile(filePath, []byte(templateString), 0644)
}

var solutionTemplate = `package main
import (
	"errors"

	aoc "github.com/herman-barnardt/aoc"
)
func init() {
	aoc.Register(
		{{.Year}},
		{{.Day}},
		solve{{.Year}}Day{{.Day}}Part1,
		solve{{.Year}}Day{{.Day}}Part2,
	)
}
func solve{{.Year}}Day{{.Day}}Part1(lines []string) interface{} {
	return errors.New("Not yet implemented")
}
func solve{{.Year}}Day{{.Day}}Part2(lines []string) interface{} {
	return errors.New("Not yet implemented")
}
`
