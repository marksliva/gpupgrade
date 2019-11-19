package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	regexes := []*regexp.Regexp{}

	regexStrings := []string{
		"Name: plpgsql; Type: EXTENSION; Schema",
		"CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;",
		"COMMENT ON EXTENSION plpgsql IS",
		"ALTER DATABASE postgres SET gp_use_legacy_hashops TO",
		"COMMENT ON DATABASE postgres IS",
		"Name: EXTENSION plpgsql; Type: COMMENT;",
		"foobar",
		"does not matter...",
	}
	for _, regexString := range regexStrings {
		regexes = append(regexes, regexp.MustCompile(regexString))
	}

	inFile, err := os.Open("new.sql")
	if err != nil {
		fmt.Printf("err: %+v", err)
		return
	}
	if inFile != nil {
		defer inFile.Close()
	}

	outFile, err := os.Create("new-sanitized.sql")
	if err != nil {
		fmt.Printf("err: %+v", err)
		return
	}
	if outFile != nil {
		defer outFile.Close()
	}
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	counter := 0
	for {
		var buffer bytes.Buffer

		var l []byte
		l, err = reader.ReadBytes('\n')
		buffer.Write(l)

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("\nerr: %+v\n\n", err)
			break
		}

		line := buffer.String()

		matched := false
		for _, regex := range regexes {
			matches := regex.FindAllString(line, -1)
			if matches != nil {
				matched = true
				break
			}
		}
		if !matched {
			counter += 1
			writtenBytes, err := writer.WriteString(line)
			counter += writtenBytes
			if err != nil {
				fmt.Printf("err writing line: %+v", err)
				break
			}
			err = outFile.Sync()
			if err != nil {
				fmt.Printf("sync failure %+v", err)
			}
		}
	}

	fmt.Printf("Wrote %d lines\n", counter)

	fmt.Printf("finished processing file\n")
}
