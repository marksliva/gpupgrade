package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var lineRegexes []*regexp.Regexp
var blockRegexes []*regexp.Regexp

func init() {
	linePatterns := []string{
		"ALTER DATABASE .+ SET gp_use_legacy_hashops TO 'on';",
	}

	blockPatterns := []string{
		"Name: plpgsql; Type: EXTENSION; Schema",
		"CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;",
		"COMMENT ON EXTENSION plpgsql IS",
		"COMMENT ON DATABASE postgres IS",
		"Name: EXTENSION plpgsql; Type: COMMENT;",
	}

	for _, pattern := range linePatterns {
		lineRegexes = append(lineRegexes, regexp.MustCompile(pattern))
	}
	for _, pattern := range blockPatterns {
		blockRegexes = append(blockRegexes, regexp.MustCompile(pattern))
	}
}

func write(out io.Writer, lines ...string) {
	for _, line := range lines {
		_, err := fmt.Fprintln(out, line)
		if err != nil {
			log.Fatalf("writing output: %+v", err)
		}
	}
}

func Filter(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	var buf []string      // lines buffered for look-ahead
	var discardEmpty bool // should we discard the next empty line?

nextline:
	for scanner.Scan() {
		line := scanner.Text()

		if discardEmpty && len(line) == 0 {
			discardEmpty = false
			continue nextline
		}

		for _, r := range lineRegexes {
			if r.MatchString(line) {
				continue nextline
			}
		}

		if strings.HasPrefix(line, "--") || len(line) == 0 {
			// A comment or an empty line. We only want to output this section
			// if the SQL it's attached to isn't filtered.
			buf = append(buf, line)
			continue nextline
		}

		for _, r := range blockRegexes {
			if r.MatchString(line) {
				// Discard this line, any buffered comment block, and any blank
				// line directly after this block.
				buf = buf[:0]
				discardEmpty = true
				continue nextline
			}
		}

		// Flush and empty our buffer.
		if len(buf) > 0 {
			write(out, buf...)
			buf = buf[:0]
		}

		write(out, line)
	}

	if scanner.Err() != nil {
		log.Fatalf("scanning stdin: %+v", scanner.Err())
	}

	/*
		// Flush and empty our buffer.
		if len(buf) > 0 {
			write(out, buf...)
			buf = buf[:0]
		}
	*/
}

func main() {
	Filter(os.Stdin, os.Stdout)
}
