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

	codes := []string{}
	for _, line := range lines {
		code, err := extractCode(line)
		if err != nil {
			return fmt.Errorf("extractCode: %w", err)
		}
		codes = append(codes, code)
	}

	duplicates := getDuplicates(codes)
	if len(duplicates) > 0 {
		fmt.Printf("Duplicate codes found:\n%s\n", duplicates)
	}

	invalid := invalidCodes(codes, *format)
	if len(invalid) > 0 {
		fmt.Printf("Invalid codes found: \n%s\n", invalid)
	}

	if v, _ := strconv.ParseBool(*fix); v {
		err := regenerateCodes(*format, duplicates, invalid, *location)
		if err != nil {
			return fmt.Errorf("regenerateCodes: %w", err)
		}
	}

	return nil
}

func runGrep(location string) ([]string, error) {
	// Run grep command
	cmd := exec.Command("grep", "-r", "-E", "ctxerr\\.(New|Wrap)", location)
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
		if strings.TrimSpace(line) == "" {
			continue
		}
		filteredLines = append(filteredLines, line)
	}
	return filteredLines, nil
}

func extractCode(line string) (string, error) {
	reWrap := regexp.MustCompile(`ctxerr\.Wrap(?:HTTP)?\([^,]+, [^,]+, "([^"]+)"`)
	reNew := regexp.MustCompile(`ctxerr\.New(?:HTTP)?\([^,]+, "([^"]+)"`)
	var code string
	switch {
	case reWrap.MatchString(line):
		matches := reWrap.FindStringSubmatch(line)
		if len(matches) > 1 {
			code = matches[1]
		}
	case reNew.MatchString(line):
		matches := reNew.FindStringSubmatch(line)
		if len(matches) > 1 {
			code = matches[1]
		}
	default:
		return "", fmt.Errorf("No match found: %s", line)
	}
	if code == "" {
		return "", fmt.Errorf("missing code: %s", strings.TrimRight(line, ":"))
	}
	return code, nil
}

func getDuplicates(codes []string) []string {
	codeMap := make(map[string]struct{})
	var duplicates []string

	for _, code := range codes {
		if _, exists := codeMap[code]; exists {
			duplicates = append(duplicates, code)
			continue
		}
		codeMap[code] = struct{}{}
	}

	return duplicates
}

func invalidCodes(codes []string, format string) []string {
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

	invalid := []string{}
	for _, code := range codes {
		if f(code) {
			invalid = append(invalid, code)
		}
	}
	return invalid
}

func regenerateCodes(format string, duplicates, invalid []string, location string) error {
	var codes []string
	switch format {
	case "uuid":
		codes = append(duplicates, invalid...)
	case "uppercase":
		codes = invalid
	}

	for _, code := range codes {
		var newCode string
		switch format {
		case "uuid":
			newCode = uuid.New().String()
		case "uppercase":
			newCode = strings.ToUpper(code)
		default:
			continue
		}
		cmd := exec.Command("sed", "-i", "", fmt.Sprintf("0,/\"%s\"/s//\"%s\"/", code, newCode), location)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to update code %s: %w", code, err)
		}
		fmt.Printf("Updated code %s to %s\n", code, newCode)
	}

	return nil
}
