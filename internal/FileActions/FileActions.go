package FileActions

import (
	"encoding/json"
	"io/ioutil"

  . "github.com/PHillemans/LearnLangs/internal/Types"
)

func ReadDataFile(name string) TranslationCollection  {
  file, err := ioutil.ReadFile("languages/"+name+".json")
  if err != nil {
    return nil
  }
  var translations TranslationCollection
  err = json.Unmarshal(file, &translations)
  if err != nil {
    return nil
  }

  return translations
}

func WriteToDataFile(w TranslationCollection, fileName string) error {
  file, err := json.MarshalIndent(w, "", "")
  if err != nil {
    return err
  }

  err = ioutil.WriteFile("languages/"+fileName+".json", file, 0644)
  if err != nil {
    return err
  } else {
    return nil
  }
}
