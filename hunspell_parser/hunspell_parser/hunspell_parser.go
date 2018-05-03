package main

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "strings"
    "encoding/json"
    "errors"
    "github.com/ewgRa/ci-utils/hunspell_parser"
)


type TypeResult struct {
    results []Result
}

type Result struct {
    line int
    word string
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    types := parseHunspellOutput(scanner);

    _, ok := types["&"]

    if ok {
        var dropCol = regexp.MustCompile(`^([^ ]+) \d+ \d+: (.*)$`)
        var minimumWord = regexp.MustCompile(`^[^ ]{3}`)

        for _, result := range types["&"].results {
            if minimumWord.MatchString(result.word) {
                suggestMatches := dropCol.FindStringSubmatch(result.word)

                if len(suggestMatches) != 3 {
                    panic(errors.New("Can't parse hunspell output"))
                }

                response := &hunspell_parser.SuggestResponse{
                    Line: result.line,
                    Word: suggestMatches[1],
                    Alternative: suggestMatches[2],
                }

                resp, err := json.Marshal(response)

                if err != nil {
                    panic(err)
                }

                fmt.Println(string(resp))
            }
        }
    }
}

func parseHunspellOutput(scanner *bufio.Scanner) map[string]*TypeResult {
    line := 1;
    types := make(map[string]*TypeResult)

    scanner.Scan()

    for scanner.Scan() {
        text := scanner.Text()

        if text == "" {
            line++;
        } else {
            resultType := text[0:1]

            typeResult, ok := types[resultType]

            if !ok {
                typeResult = &TypeResult{}
                types[resultType] = typeResult
            }

            typeResult.results = append(typeResult.results, Result{line: line, word: strings.Trim(text[1:], " ")})
        }
    }

    return types
}
