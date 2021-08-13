package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
  tea "github.com/charmbracelet/bubbletea"
  "github.com/PHillemans/LearnLangs/cmd"
)

type Translation struct {
  Order int
  Original string
  Translation string
}

type TranslationCollection []Translation

var chosenLanguage string

func Execute() {
  var mostCommonWords TranslationCollection

  chosenLanguage = "spanish"

  mostCommonWords = RetreiveWordsForLang(chosenLanguage)

  mostCommonWords.WriteToDataFile(chosenLanguage)
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func (w TranslationCollection) WriteToDataFile(fileName string) error {
  file, err := json.MarshalIndent(w, "", "")
  if err != nil {
    return err
  }

  InfoLogger.Println("Writing words to file")

  err = ioutil.WriteFile(fileName+".json", file, 0644)
  if err != nil {
    return err
  } else {
    return nil
  }
}
