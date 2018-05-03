package links_checker

import (
    "net/http"
    "crypto/tls"
)

// Checker help you to check is link have expected response code
type Checker struct {
    expectedCodesList expectedCodes
    client            *http.Client
}

func NewChecker(expectedCodesFile string) *Checker {
    transCfg := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
    }

    return &Checker{expectedCodesList: readExpectedCodes(expectedCodesFile), client: &http.Client{Transport: transCfg}}
}

func (c *Checker) Check(link string) (bool, int, []int) {
    req, err := http.NewRequest("GET", link, nil)

    if err != nil {
        panic(err)
    }

    req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

    resp, err := c.client.Do(req)

    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    expectedCodes := []int{200}
    expectedCodes = append(expectedCodes, c.expectedCodesList.getList(link)...)

    return resp.StatusCode == 200 || c.expectedCodesList.has(link, resp.StatusCode), resp.StatusCode, expectedCodes
}
