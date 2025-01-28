package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	err := errorMain()
	if err != nil {
		log.Fatal(err)
	}
}

func errorMain() error {
	location := flag.String("location", ".", "The location to run grep")
	format := flag.String("format", "", "The format to search for")
	fix := flag.String("fix", "", "Auto fix duplicates or invalid if possible")
	flag.Parse()

	//fmt.Println("Hello, playground")
	lines, err := runGrep(*location)
	if err != nil {
		return fmt.Errorf("runGrep: %w", err)
	}
	//fmt.Println(lines)

	codes := []grepParts{}
	missings := []grepParts{}
	for _, line := range lines {
		gp, err := extractCode(line)
		if err != nil {
			return fmt.Errorf("extractCode: %w", err)
		}
		if gp.code == "" {
			missings = append(missings, gp)
			continue
		}
		codes = append(codes, gp)
	}
	if len(missings) > 0 {
		lines := []string{}
		for _, m := range missings {
			lines = append(lines, fmt.Sprintf("%s:%s", m.filePath, m.line))
		}
		fmt.Printf("Missing codes: \n%s\n\n", lines)
	}

	duplicates := getDuplicates(codes)
	if len(duplicates) > 0 {
		codes := []string{}
		for _, d := range duplicates {
			codes = append(codes, d.code)
		}
		fmt.Printf("Duplicate codes:\n%s\n\n", codes)
	}

	invalid := invalidCodes(codes, *format)
	if len(invalid) > 0 {
		codes := []string{}
		for _, d := range invalid {
			codes = append(codes, d.code)
		}
		fmt.Printf("Invalid codes: \n%s\n\n", codes)
	}

	if v, _ := strconv.ParseBool(*fix); v {
		badLines := append(duplicates, invalid...)
		badLines = append(badLines, missings...)
		err = fixCodes(*format, badLines)
		if err != nil {
			return fmt.Errorf("fixMissingCodes: %w", err)
		}
	}

	return nil
}

func runGrep(location string) ([]string, error) {
	// Run grep command
	cmd := exec.Command("grep", "-rn", "--include=*.go", "-E", "ctxerr\\.(New|Wrap)", location)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(output), "\n")
	var filteredLines []string
	for _, line := range lines {
		if strings.Contains(line, "Binary file") {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Filter out comments
		split := strings.SplitN(line, ":", 3)
		if len(split) == 3 {
			if strings.HasPrefix(strings.TrimSpace(split[2]), "//") {
				continue
			}
		}

		filteredLines = append(filteredLines, line)
	}
	return filteredLines, nil
}

var reWrap = regexp.MustCompile(`(ctxerr\.Wrap(?:HTTP)?(?:f)?\([^,]+, [^,]+, ")(?<code>[^"]*)(".*)`)
var reNew = regexp.MustCompile(`(ctxerr\.New(?:HTTP)?(?:f)?\([^,]+, ")(?<code>[^"]*)(".*)`)

type grepParts struct {
	filePath string
	line     string
	content  string
	code     string
}

func extractCode(line string) (grepParts, error) {
	parts := strings.SplitN(line, ":", 3)
	if len(parts) != 3 {
		return grepParts{}, fmt.Errorf("invalid line format: %s", line)
	}
	r := grepParts{
		filePath: parts[0],
		line:     parts[1],
		content:  parts[2],
	}

	switch {
	case reWrap.MatchString(line):
		matches := reWrap.FindStringSubmatch(line)
		//fmt.Println("wrap", matches)
		if len(matches) > 2 {
			r.code = matches[2]
		}
	case reNew.MatchString(line):
		matches := reNew.FindStringSubmatch(line)
		//fmt.Println("new", matches)
		if len(matches) > 2 {
			r.code = matches[2]
		}
	default:
		return r, fmt.Errorf("No match found: %s", line)
	}

	//log.Println(line)
	//log.Printf("%#v\n", r)
	//panic("hello")

	return r, nil
}

func getDuplicates(gps []grepParts) []grepParts {
	codeMap := make(map[string]struct{})
	duplicates := []grepParts{}

	for _, gp := range gps {
		if _, exists := codeMap[gp.code]; exists {
			duplicates = append(duplicates, gp)
			continue
		}
		codeMap[gp.code] = struct{}{}
	}

	return duplicates
}

func invalidCodes(gps []grepParts, format string) []grepParts {
	var f func(code string) bool
	switch format {
	case "uuid":
		f = func(code string) bool {
			_, err := uuid.Parse(code)
			return err != nil
		}
	case "uppercase":
		f = func(code string) bool {
			return code != strings.ToUpper(code)
		}
	}
	if f == nil {
		return nil
	}

	invalid := []grepParts{}
	for _, gp := range gps {
		if f(gp.code) {
			invalid = append(invalid, gp)
		}
	}
	return invalid
}

func fixCodes(format string, gps []grepParts) error {
	if format != "uuid" {
		return nil
	}

	for _, gp := range gps {
		newCode := uuid.New().String()
		var newContent string
		switch {
		case reWrap.MatchString(gp.content):
			newContent = reWrap.ReplaceAllString(gp.content, fmt.Sprintf(`${1}%s${3}`, newCode))
		case reNew.MatchString(gp.content):
			newContent = reNew.ReplaceAllString(gp.content, fmt.Sprintf(`${1}%s${3}`, newCode))
		default:
			return fmt.Errorf("could not replace in line %+v", gp)
		}
		log.Println("newContent", newContent)
		fmt.Println()

		// Replace the entire line with the updated value in the variable `newContent`
		cmd := exec.Command("sed", "-i", "", fmt.Sprintf("%ss/.*/%s/", gp.line, newContent), gp.filePath)
		var stderr strings.Builder
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to replace line in file %s at line %s: %w, stderr: %s", gp.filePath, gp.line, err, stderr.String())
		}
		fmt.Printf("Updated code '%s' in file %s at line %s with code %s\n", gp.code, gp.filePath, gp.line, newCode)
	}
	return nil
}
