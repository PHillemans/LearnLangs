package cmd

import (
	"log"
	"os"
)

var (
  WarningLogger *log.Logger
  InfoLogger *log.Logger
  ErrorLogger *log.Logger
)

func init() {
  WarningLogger = log.New(os.Stderr, "Warning: ", log.Ldate|log.Ltime)
  InfoLogger = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime)
  ErrorLogger = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime)
}
