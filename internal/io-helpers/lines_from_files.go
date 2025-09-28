package iohelpers

import (
	"io"
	"strings"
)

func GetLinesChannel(f io.ReadCloser) <-chan string {

	chunks := make([]byte,8)
	str := ""

	channel := make(chan string);

	go func() {

		defer f.Close()
		defer close(channel)


		for {
			n, err := f.Read(chunks);
			if err != nil || err == io.EOF {
				channel <- str
				return
			}

			parts := strings.Split(string(chunks[:n]), "\n");
			if len(parts) > 1 {
				str += parts[0];
				channel <- str
				str = ""
				str += strings.Join(parts[1:], "")
				continue
			}

			str += strings.Join(parts[:], "")
		}
	}()

	return channel
}