package response

import (
	"fmt"
	"io"
	"strconv"

	"github.com/vijayvenkatj/http-protocol/internal/headers"
)

type StatusCode int

const (
	OK                    StatusCode = 200
	BAD_REQUEST           StatusCode = 400
	INTERNAL_SERVER_ERROR StatusCode = 500
)

func WriteStatusLine(w io.Writer, statusCode StatusCode) error {
	switch statusCode {
		case OK:
			_, err := w.Write([]byte("HTTP/1.1 200 OK\r\n"))
			return err
		case BAD_REQUEST:
			_, err := w.Write([]byte("HTTP/1.1 400 Bad Request\r\n"))
			return err
		case INTERNAL_SERVER_ERROR:
			_, err := w.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n"))
			return err
		default:
			_, err := w.Write([]byte(fmt.Sprintf("HTTP/1.1 %d \r\n", statusCode)))
			return err
	}
}

func GetDefaultHeaders(contentLen int) headers.Headers {
	headers := headers.Headers{}

	headers.Set("Content-Type", "text/html")
	headers.Set("Content-Length", strconv.Itoa(contentLen))
	headers.Set("Connection", "close")

	return headers
}

func WriteHeaders(w io.Writer, headers headers.Headers) error {
	for key,val := range headers {
		_, err := w.Write([]byte(key + ":" + val + "\r\n"));
		if err != nil {
			return err
		}
	}
	_, err := w.Write([]byte("\r\n"))
	return err
}