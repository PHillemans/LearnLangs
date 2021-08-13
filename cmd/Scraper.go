package cmd

import "strconv"

func RetreiveWordsForLang(lang string) TranslationCollection {
  var collection TranslationCollection

  c := colly.NewCollector()

  c.OnHTML("tr", func(e *colly.HTMLElement){
    var commonWordTranslation Translation
    e.ForEach("td", func(cijfer int, e *colly.HTMLElement){
      // Parsing words table from 1000mostcommonwords
      switch(cijfer) {
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
        WarningLogger.Println("Something is wrong here:", e.Text)
      }
    })
    collection = append(collection, commonWordTranslation)
  })

  c.OnRequest(func(r *colly.Request) {
    InfoLogger.Println("Visiting and getting words from:", r.URL)
  })

  c.Visit("https://1000mostcommonwords.com/1000-most-common-"+lang+"-words/")
  return collection
}
