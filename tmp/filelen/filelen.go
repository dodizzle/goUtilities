package filelen

import (
	"io"
	"os"
)

func Filelen(fname string) (int, error) {
	file, err := os.Open(fname)
	if err != nil {
		return 0, err
	}
	len := make([]byte, 2048)
	var total int
	for {
		count, err := file.Read(len)
		if err == io.EOF {
			break
		}

		total += count
	}
	return total, nil
}
