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

type Writer struct {
    w io.Writer
}

func NewWriter(w io.Writer) *Writer {
    return &Writer{w: w}
}


func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
	switch statusCode {
		case OK:
			_, err := w.w.Write([]byte("HTTP/1.1 200 OK\r\n"))
			return err
		case BAD_REQUEST:
			_, err := w.w.Write([]byte("HTTP/1.1 400 Bad Request\r\n"))
			return err
		case INTERNAL_SERVER_ERROR:
			_, err := w.w.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n"))
			return err
		default:
			_, err := w.w.Write([]byte(fmt.Sprintf("HTTP/1.1 %d \r\n", statusCode)))
			return err
	}
}

func (w *Writer) GetDefaultHeaders(contentLen int) headers.Headers {
	headers := headers.Headers{}

	headers.Set("Content-Type", "text/html")
	headers.Set("Content-Length", strconv.Itoa(contentLen))
	headers.Set("Connection", "close")

	return headers
}

func (w *Writer) WriteHeaders(headers headers.Headers) error {
	for key,val := range headers {
		_, err := w.w.Write([]byte(key + ":" + val + "\r\n"));
		if err != nil {
			return err
		}
	}
	_, err := w.w.Write([]byte("\r\n"))
	return err
}

func (w *Writer) WriteBody(data []byte) error {
	_, err := w.w.Write(data);
	return err
}