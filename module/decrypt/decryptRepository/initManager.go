package decryptRepository

import (
	"fmt"
	"io"
	"os"
)

type ManagerString struct {
	Source string
	Des    string
}

func PassSource(input string) *ManagerString {
	return &ManagerString{input, input}
}

func ConvertFileToString(fileName string) string {

	result := ""

	file, err := os.Open(fmt.Sprintf("./static/%s", fileName))
	if err != nil {
		fmt.Println("Error opening file!!!")
	}
	defer file.Close()

	// declare chunk size
	const maxSz = 4

	// create buffer
	b := make([]byte, maxSz)

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		result += string(b[:readTotal])

	}
	return result
}
