package cmd

import (
	"encoding/json"
	"io/ioutil"
)

func ReadDataFile(name string) TranslationCollection  {
  file, err := ioutil.ReadFile(name+".json")
  if err != nil {
    return nil
  }
  var translations TranslationCollection
  InfoLogger.Println("Fetching internal file data")
  err = json.Unmarshal(file, &translations)
  if err != nil {
    return nil
  }

  return translations
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
