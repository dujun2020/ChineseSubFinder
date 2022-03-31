package chs_cht_changer

import (
	"github.com/allanpk716/ChineseSubFinder/internal/pkg/language"
	"os"
)

func Process(srcSubFileFPath string, desChineseLanguageType int) error {

	// 默认 0 是 简体 ，1 是 繁体
	fBytes, err := os.ReadFile(srcSubFileFPath)
	if err != nil {
		return err
	}

	orgString := string(fBytes)
	outString := ""
	if desChineseLanguageType == 0 {
		// 简体
		outString = language.ChDict.Read(orgString)
	} else {
		// 繁体
		outString = language.ChDict.ReadReverse(orgString)
	}

	err = os.WriteFile(srcSubFileFPath, []byte(outString), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
