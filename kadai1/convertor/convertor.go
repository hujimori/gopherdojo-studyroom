package convertor

import (
	"image"
	"os"
)

const (
	JPG  = "jpg"
	JPEG = "jpeg"
	PNG  = "png"
)

type convertor struct {
	srcFile          *os.File
	dstFile          *os.File
	srcFileExtension string
	dstFileExtension string

	image image.Image
}

func New(srcFile *os.File, dstFile *os.File) *convertor {

	c := &convertor{
		srcFile: srcFile,
		dstFile: dstFile,
	}

	return c
}

// TODO 画像デコード
func Start() (image.Image, error) {

}

// TODO ファイルパスチェック
func isExistFile(filepath string) {

}
