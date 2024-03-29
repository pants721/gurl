package methods

import (
    "bytes"
    "io"
    "net/http"
    "strings"
)

const (
    MethodGet  = "GET"
    MethodPost = "POST"
    MethodDelete = "DELETE"
    MethodPut = "PUT"
)

func Request(method, url string, headers []string, body string) (respBody string, err error) {
    reqBody := bytes.NewBuffer([]byte(body))

    client := http.DefaultClient
    req, err := http.NewRequest(method, url, reqBody)
    if err != nil {
        return "", err
    }

    req.Header.Add("User-Agent", "gurl")
    req.Header.Add("Accept", "*/*")
    
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
        return "", err
    }
    defer resp.Body.Close()

    b, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    respBody = string(b)
    return respBody, nil
}

func Get(url string, headers []string, body string) (respBody string, err error) {
    return Request(MethodGet, url, headers, body)
}

func Post(url string, headers []string, body string) (respBody string, err error) {
    return Request(MethodPost, url, headers, body)
}

func Delete(url string, headers []string, body string) (respBody string, err error) {
    return Request(MethodDelete, url, headers, body)
}

func Put(url string, headers []string, body string) (respBody string, err error) {
    return Request(MethodPut, url, headers, body)
}

