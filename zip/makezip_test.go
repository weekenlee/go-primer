package zip

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	err := Zip("demo.zip", []string{
		"/Users/liweijian/Code/go/zip/makezip.go",
		"/Users/liweijian/Code/go/zip/makezip_test.go",
	})
	if err != nil {
		fmt.Println(err)
	}

}

func TestUnzip(t *testing.T) {
	err := Unzip("demo.zip")
	if err != nil {
		fmt.Println(err)
	}
}
