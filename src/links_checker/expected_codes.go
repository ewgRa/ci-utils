package links_checker

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

type expectedCodes map[string][]int

func (ec expectedCodes) has(link string, code int) bool {
    if codes, ok := ec[link]; ok {
        for _, expectedCode := range codes {
            if expectedCode == code {
                return true
            }
        }
    }

    return false
}

func (ec expectedCodes) getList(link string) []int {
    if codes, ok := ec[link]; ok {
        return codes
    }

    return []int{}
}

func readExpectedCodes(fileName string) expectedCodes {
    res := make(expectedCodes, 0)

    file, err := os.Open(fileName)

    if err != nil {
        panic(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        data := strings.SplitN(line, ",", 2)
        code, err := strconv.Atoi(data[0])

        if err != nil {
            panic(err)
        }

        link := data[1]

        if _, ok := res[link]; !ok {
            res[link] = []int{}
        }

        res[link] = append(res[link], code)
    }

    return res

}