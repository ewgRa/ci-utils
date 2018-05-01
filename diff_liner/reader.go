package diff_liner

import (
	"bufio"
	"os"
	"encoding/json"
)

type LinerResponse struct {
	lines map[string]map[int]int
}

func NewLinerResponse () *LinerResponse {
	return &LinerResponse{lines: make(map[string]map[int]int, 0)}
}

func (lr *LinerResponse) Add(resp *DiffLinerResponse) {
	if _, ok := lr.lines[resp.File]; !ok {
		lr.lines[resp.File] = make(map[int]int, 0)
	}

	lr.lines[resp.File][resp.Line] = resp.DiffLine
}

func (lr *LinerResponse) GetDiffLine(file string, line int) int {
	if _, ok := lr.lines[file]; !ok {
		return 0
	}

	if _, ok := lr.lines[file][line]; !ok {
		return 0
	}

	return lr.lines[file][line]
}

// ReadLinerResponse read liner response to LinerResponse struct, that can be used in third-party golang scripts to parse it and generate for example json for
// comment via github API
func ReadLinerResponse(fileName string) *LinerResponse {
	linerResponse := NewLinerResponse()

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var resp *DiffLinerResponse

		line := scanner.Bytes()
		err := json.Unmarshal(line, &resp)

		if err != nil {
			panic(err)
		}

		linerResponse.Add(resp)
	}

	if err := scanner.Err(); err != nil {
		panic(err)

	}

	return linerResponse
}