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
var solutionDirectory = ""

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

func SetSolutionDirectory(s string) {
	solutionDirectory = s
}

func Run(cmd string, year int, day int, part int) error {
	switch cmd {
	case "create":
		if day > 0 {
			return create(year, day)
		} else {
			for i := 1; i <= 25; i++ {
				if err := create(year, i); err != nil {
					return err
				}
			}
			return nil
		}
	case "download":
		if day > 0 {
			return download(year, day)
		} else {
			for i := 1; i <= 25; i++ {
				if err := download(year, i); err != nil {
					return err
				}
			}
			return nil
		}

	case "solve":
		return solve(fmt.Sprintf("./data/%d/day%d", year, day), year, day, part)
	case "test":
		return solve(fmt.Sprintf("./data/%d/day%dTest", year, day), year, day, part)
	default:
		return errors.New(fmt.Sprintf("unsupported command: %q", cmd))

	}
}

func makeDir(path string) (bool, error) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
		return true, nil
	}
	return false, nil
}

func getFile(year, day int) error {
	cookie, err := readLines("session")

	if err != nil {
		return err
	}

	if len(cookie) == 0 {
		return errors.New("Missing session file")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie[0]})

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	f, err := os.Create(fmt.Sprintf("./data/%d/day%d", year, day))
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	if err != nil {
		return err
	}
	return nil
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
			return fmt.Errorf("Solution not implemented\nTemplate generation has failed %v", err)
		}
		return errors.New("Solution not implemented\nTemplate has been generated")
	}

	lines, err := readLines(path)

	if err != nil {
		return fmt.Errorf("Error reading file: %s %v", path, err)
	}
	if len(lines) == 0 {
		return fmt.Errorf("Input file %s is empty", path)
	}

	fmt.Printf("Solving %d Day %d\n", year, day)

	if part == 1 || part == 0 {
		startTime := time.Now()
		solution1 := allSolutions[year][day].Part1(lines)
		duration1 := time.Since(startTime)
		fmt.Printf("Part 1: %v (%v)\n", solution1, duration1)
	}
	if part == 2 || part == 0 {
		startTime := time.Now()
		solution2 := allSolutions[year][day].Part2(lines)
		duration2 := time.Since(startTime)
		fmt.Printf("Part 2: %v (%v)\n", solution2, duration2)
	}
	return nil
}

func download(year, day int) error {
	path := fmt.Sprintf("./data/%d/day%d", year, day)
	fmt.Println("Downloading file", path)
	created, err := makeDir("./data")
	if err != nil {
		return err
	}
	if created {
		fmt.Println("Created directory ./data")
	}
	created, err = makeDir(fmt.Sprintf("./data/%d", year))
	if err != nil {
		return err
	}
	if created {
		fmt.Printf("Created directory ./data/%d\n", year)
	}
	if err = getFile(year, day); err != nil {
		return err
	}

	return nil
}

func create(year, day int) error {
	filename := fmt.Sprintf("day%d", day)

	if filenameIncludesYear {
		filename = fmt.Sprintf("%dday%d", year, day)
	}

	fmt.Println("Creating file from template: ", filename)

	basepath, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := path.Join(basepath, solutionDirectory, filename)

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
	aoc.Register({{.Year}}, {{.Day}}, solve{{.Year}}Day{{.Day}}Part1, solve{{.Year}}Day{{.Day}}Part2)
}

func solve{{.Year}}Day{{.Day}}Part1(lines []string) interface{} {
	return errors.New("Not yet implemented")
}

func solve{{.Year}}Day{{.Day}}Part2(lines []string) interface{} {
	return errors.New("Not yet implemented")
}
`
