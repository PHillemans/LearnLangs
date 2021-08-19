package types

type Translation struct {
  Order int
  Original string
  Translation string
}

type TranslationCollection []Translation

type TranslationMsg TranslationCollection

