package qrcode

import (
	"image/jpeg"

	"github.com/EvanXzj/gin-blog/pkg/file"
	"github.com/EvanXzj/gin-blog/pkg/setting"
	"github.com/EvanXzj/gin-blog/pkg/util"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Model  qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

// NewQrCode initialize a qrcode instance
func NewQrCode(url string, w, h int, l qr.ErrorCorrectionLevel, model qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  w,
		Height: h,
		Level:  l,
		Model:  model,
		Ext:    EXT_JPG,
	}
}

// GetQrCodeSavePath get qrcode save path
func GetQrCodeSavePath() string {
	return setting.AppSetting.QrCodeSavePath
}

// GetQrCodeFullPath get qrcode full save path
func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetQrCodeSavePath()
}

// GetQrCodeFullUrl get the full access path
func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodeFullPath() + name
}

// GetQrCodeFileName get qrcode file name
func GetQrCodeFileName(name string) string {
	return util.EncodeMD5(name)
}

// GetQrCodeExt get qrcode ext
func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// Encode generate QR code
func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Model)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
