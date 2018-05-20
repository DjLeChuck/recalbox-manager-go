package routes

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/kataras/iris"

	"github.com/spf13/viper"

	"github.com/djlechuck/recalbox-manager/structs"
	"github.com/djlechuck/recalbox-manager/utils/bios"
	"github.com/djlechuck/recalbox-manager/utils/errors"
	"github.com/djlechuck/recalbox-manager/utils/md5"
)

// GetBiosHandler handles the GET requests on /bios.
func GetBiosHandler(ctx iris.Context) {
	biosPath := viper.GetString("recalbox.bios.filesPath")
	files, err := ioutil.ReadDir(biosPath)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	md5File := viper.GetString("recalbox.bios.md5FilePath")
	biosList, err := bios.GetList(md5File)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	for _, file := range files {
		// Exclude directories and .txt files
		if file.IsDir() || filepath.Ext(file.Name()) == ".txt" {
			continue
		}

		// Init BIOS file and check MD5
		for k, b := range biosList {
			if b.Name == file.Name() {
				fileMd5, err := md5.GetFileMd5(biosPath + file.Name())

				if err != nil {
					ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
					ctx.StatusCode(500)

					return
				}

				biosList[k].IsPresent = true
				biosList[k].IsValid = b.CheckValidity(fileMd5)
			}
		}
	}

	ctx.ViewData("PageTitle", ctx.Translate("Bios.Title"))

	ctx.ViewData("BiosPath", biosPath)
	ctx.ViewData("BiosList", biosList)
	ctx.ViewData("Tr", map[string]interface{}{
		"Text1":         ctx.Translate("Bios.Text1"),
		"Text2":         ctx.Translate("Bios.Text2"),
		"Text3":         ctx.Translate("Bios.Text3"),
		"UploadBiosBtn": ctx.Translate("Bios.UploadBtn"),
		"TableHeader": map[string]string{
			"Bios":   ctx.Translate("BIOS"),
			"Md5":    ctx.Translate("MD5 attendu"),
			"Valid":  ctx.Translate("Valide"),
			"Action": ctx.Translate("Action"),
		},
	})

	ctx.View("views/bios.pug")
}

// GetBiosCheckHandler handles the GET requests on /bios/check.
func GetBiosCheckHandler(ctx iris.Context) {
	biosPath := viper.GetString("recalbox.bios.filesPath")
	md5File := viper.GetString("recalbox.bios.md5FilePath")
	biosList, err := bios.GetList(md5File)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	fileName := ctx.URLParam("file")
	biosFile := structs.BiosFile{}

	// Init BIOS file and check MD5
	for k, b := range biosList {
		if b.Name == fileName {
			biosFile = biosList[k]
			fileMd5, err := md5.GetFileMd5(biosPath + fileName)

			if err != nil {
				ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
				ctx.StatusCode(500)

				return
			}

			biosFile.IsPresent = true
			biosFile.IsValid = b.CheckValidity(fileMd5)
		}
	}

	ctx.JSON(iris.Map{"success": true, "data": biosFile})
}

// PostBiosUploadHandler handles the POST requests on /bios/upload.
func PostBiosUploadHandler(ctx iris.Context) {
	biosPath := viper.GetString("recalbox.bios.filesPath")
	file, info, err := ctx.FormFile("file")

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	defer file.Close()
	fname := info.Filename

	// Create a file with the same name
	out, err := os.OpenFile(biosPath+fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		ctx.Values().Set("error", errors.FormatErrorForLog(ctx, err.(error)))
		ctx.StatusCode(500)

		return
	}

	ctx.JSON(iris.Map{"success": true})
}
