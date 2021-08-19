package scrape

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"

  . "github.com/PHillemans/LearnLangs/internal/Types"
)

func RetreiveWordsForLang(lang string) (TranslationCollection, error) {
  var collection TranslationCollection

  c := colly.NewCollector()

  c.OnHTML("tr", func(e *colly.HTMLElement){
    var commonWordTranslation Translation

    e.ForEach("td", func(column int, e *colly.HTMLElement){
      // Parsing words table from 1000mostcommonwords
      switch(column) {
      case 0:
        if (e.Text == "Number") {
          commonWordTranslation.Order = 0
        } else {
          commonWordTranslation.Order,_ = strconv.Atoi(e.Text)
        }
      case 1:
        commonWordTranslation.Original = e.Text
      case 2:
        commonWordTranslation.Translation = e.Text
      default:
        fmt.Println("Something is wrong here:", e.Text)
      }
    })
    collection = append(collection, commonWordTranslation)
  })

  url := "https://1000mostcommonwords.com/1000-most-common-"+lang+"-words/"

  if checkIfExists(url) != nil {
    return nil, errors.New("No languages found")
  } else {
    c.Visit(url)
    return collection, nil
  }
}

func checkIfExists(url string) error {
  c := &http.Client{Timeout: 10 * time.Second}
  res, err := c.Get(url)

  if res.StatusCode != 200 || err != nil {
    return errors.New("does not exist")
  }
  return nil
}


type errMsg struct{ err error }
