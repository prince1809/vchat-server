package utils

import (
	"fmt"
	"github.com/mattermost/go-i18n/i18n"
	"github.com/prince1809/vchat-server/model"
	"github.com/prince1809/vchat-server/utils/fileutils"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var T i18n.TranslateFunc
var TDefault i18n.TranslateFunc
var locales map[string]string = make(map[string]string)
var setting model.LocalizationSettings

// this functions load translations from filesystem if they are not
// loaded already and assigns english while loading server config
func TranslationsPreInit() error {
	if T != nil {
		return nil
	}

	// Set T even if we fail to load the translations. Lots of shutdown handling code will
	// segfault trying to handle the error, and the untranslated IDs are strictly better.
	T = TfuncWithFallback("en")
	TDefault = TfuncWithFallback("en")
	return InitTranslationsWithDir("i18n")
}

func InitTranslationsWithDir(dir string) error {
	i18nDirectory, found := fileutils.FindDir(dir)
	if !found {
		return fmt.Errorf("Unable to find i18n directory")
	}

	files, _ := ioutil.ReadDir(i18nDirectory)
	for _, f := range files {
		if filepath.Ext(f.Name()) == ".json" {
			filename := f.Name()
			locales[strings.Split(filename, ".")[0]] = filepath.Join(i18nDirectory, filename)

			if err := i18n.LoadTranslationFile(filepath.Join(i18nDirectory, filename)); err != nil {
				return err
			}
		}
	}

	return nil
}

func TfuncWithFallback(pref string) i18n.TranslateFunc {
	t, _ := i18n.Tfunc(pref)
	return func(translationID string, args ...interface{}) string {
		if translated := t(translationID, args...); translated != translationID {
			return translated
		}

		t, _ := i18n.Tfunc(model.DEFAULT_LOCALE)
		return t(translationID, args...)
	}
}
