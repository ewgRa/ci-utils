package hunspell_parser

import (
	"bufio"
	"os"
	"encoding/json"
)

// ReadHunspellParserResponse read hunspell parser response to SuggestResponse struct, that can be used in third-party golang scripts to parse it and generate for example json for
// comment via github API
func ReadHunspellParserResponse(fileName string) []*SuggestResponse {
	var response []*SuggestResponse

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var resp *SuggestResponse

		line := scanner.Bytes()
		err := json.Unmarshal(line, &resp)

		if err != nil {
			panic(err)
		}

		response = append(response, resp)
	}

	if err := scanner.Err(); err != nil {
		panic(err)

	}

	return response
}