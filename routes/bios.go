package routes

import (
	"io/ioutil"
	"path/filepath"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils"
)

// GetBiosHandler handles the GET requests on /bios.
func GetBiosHandler(ctx iris.Context) {
	biosPath := viper.GetString("recalbox.bios.filesPath")
	md5File := viper.GetString("recalbox.bios.md5FilePath")
	biosMd5List := utils.GetBiosMD5List(md5File)
	files, err := ioutil.ReadDir(biosPath)

	if err != nil {
		panic(err)
	}

	biosList := []structs.BiosFile{}

	for _, file := range files {
		// Exclude directories and .txt files
		if file.IsDir() || filepath.Ext(file.Name()) == ".txt" {
			continue
		}

		// Init BIOS file and check MD5
		biosFile := structs.BiosFile{
			Name: file.Name(),
			Md5:  biosMd5List[file.Name()],
		}
		fileMd5 := utils.GetFileMd5(biosPath + file.Name())
		biosFile.IsValid = biosFile.CheckValidity(fileMd5)

		// Add to the list
		biosList = append(biosList, biosFile)
	}

	ctx.ViewData("PageTitle", ctx.Translate("Bios.Title"))

	ctx.ViewData("BiosPath", biosPath)
	ctx.ViewData("BiosList", biosList)
	ctx.ViewData("Tr", map[string]interface{}{
		"Text1": ctx.Translate("Bios.Text1"),
		"Text2": ctx.Translate("Bios.Text2"),
		"Text3": ctx.Translate("Bios.Text3"),
		"TableHeader": map[string]string{
			"Bios":   ctx.Translate("BIOS"),
			"Md5":    ctx.Translate("MD5 attendu"),
			"Valid":  ctx.Translate("Valide"),
			"Action": ctx.Translate("Action"),
		},
	})

	ctx.View("views/bios.pug")
}
