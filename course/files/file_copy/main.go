package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

const (
	BUFFER_SIZE = 4096 //default size of bufio
)

func FileCopy_io_copy(desFilePath, oriFilePath string) error {
	oldFile, err := os.Open(oriFilePath)

	if err != nil {
		return err
	}
	defer oldFile.Close()

	newFile, err := os.Create(desFilePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, oldFile)
	if err != nil {
		return err
	}

	err = newFile.Sync()

	return err
}

func FileCopy_ioutil_WriteFile(desFilePath, oriFilePath string) error {
	oriContent, err := os.ReadFile(oriFilePath)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(desFilePath, oriContent, 0666)
	return err
}

func FileCopy_os_Read_os_Write(desFilePath, oriFilePath string) error {
	oldFile, err := os.Open(oriFilePath)

	if err != nil {
		return err
	}
	defer oldFile.Close()

	newFile, err := os.Create(desFilePath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	buf := make([]byte, BUFFER_SIZE)
	for {
		n, err := oldFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := newFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("File Copy it's running")
	if err := FileCopy_io_copy(fmt.Sprintf("./new_%d_test.txt", rand.Intn(1000)), "../test.txt"); err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("File Copy it's end")
}
