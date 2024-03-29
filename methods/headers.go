package methods

import (
	"net/http"
	"time"
)

func GurlDefaultHeaders() []http.Header {
    return []http.Header{
        {"User-Agent": []string{"gurl"}},
        {"Accept": []string{"*/*"}},
        {"Date": []string{time.Now().Format(time.RFC1123)}},
    }
}
