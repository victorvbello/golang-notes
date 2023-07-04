package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
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

func FileRemove(filePath string) error {
	return os.Remove(filePath)
}

func BenchmarkFileCopy_io_copy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newFile := fmt.Sprintf("./new_%d_test.txt", i)
		if err := FileCopy_io_copy(newFile, "./test.txt"); err != nil {
			b.Error(err)
		}
		if err := FileRemove(newFile); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFileCopy_ioutil_WriteFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newFile := fmt.Sprintf("./new_%d_test.txt", i)
		if err := FileCopy_ioutil_WriteFile(newFile, "./test.txt"); err != nil {
			b.Error(err)
		}
		if err := FileRemove(newFile); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkFileCopy_os_Read_os_Write(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newFile := fmt.Sprintf("./new_%d_test.txt", i)
		if err := FileCopy_os_Read_os_Write(newFile, "./test.txt"); err != nil {
			b.Error(err)
		}
		if err := FileRemove(newFile); err != nil {
			b.Error(err)
		}
	}
}
