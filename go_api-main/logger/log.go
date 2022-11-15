package logger

/**
 * Copyright by NAB.
 * Creator: Le Cong Nghi
 * Date: 31/10/2022
 * Time: 4:41 PM
 */
import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	DebugLogger   *log.Logger
	letters       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	InfoLogger = log.New(file, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "[WARNING]: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file, "[DEBUG]: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(b))
	return string(b)
}
