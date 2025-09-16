package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type fileLogger struct {
	file *os.File
	once *sync.Once
}

func createFileLogger(path string) fileLogger {
	dir, err := os.MkdirTemp(path, "log")
	if err != nil {
		panic(err)
	}
	logFile, err := os.CreateTemp(dir, "log")
	if err != nil {
		log.Fatal(err)
	}
	return fileLogger{
		logFile,
		&sync.Once{},
	}
}

func (f fileLogger) logging(s string) error {
	logFile, err := os.OpenFile(f.file.Name(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	logger := log.New(logFile, "", log.Ldate|log.Ltime)
	// logger.SetFlags(log.Ldate | log.Ltime)
	logger.Println(s)
	defer f.file.Close()
	return err
}

func (f fileLogger) close() (err error) {
	f.once.Do(func ()  {
		err = f.file.Close()		
	})
	return err
}

func main() {
	logFile := createFileLogger(".")
	// err := logFile.logging("Hello")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(logFile.file.Name())
	defer logFile.file.Close()

}
