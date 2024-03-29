package methods

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

func Get(url string) (respBody string, err error) {
	resp, err := http.Get(url)
    if err != nil {
        return
    }

    b, err := io.ReadAll(resp.Body)
    if err != nil {
        return
    }

    respBody = string(b)
    return
}

func Post(url string, headers []string, body string) (respBody string, err error) {
    reqBody := bytes.NewBuffer([]byte(body))

    
    client := &http.Client{}
    req, err := http.NewRequest("POST", url, reqBody)
    if err != nil {
        return
    }

    for _, header := range headers {
        parts := strings.Split(header, ":")
        if len(parts) >= 2 {
            k := strings.TrimSpace(parts[0])
            v := strings.TrimSpace(parts[1])

            req.Header.Add(k, v)
        }
    }

    resp, err := client.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close() 
    
    b, err := io.ReadAll(resp.Body)
    if err != nil {
        return
    }

    respBody = string(b)
    return
}
