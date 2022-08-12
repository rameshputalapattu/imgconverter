package imageconverter

import (
	"errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/sirupsen/logrus"
)

//Converts every image into RGBA format
func ConvertToRGBA(img image.Image) *image.RGBA {
	var imgRgba *image.RGBA
	switch src := img.(type) {
	case *image.NRGBA64, *image.NRGBA, *image.RGBA64, *image.NYCbCrA, *image.CMYK, *image.YCbCr:
		b := src.Bounds()
		imgRgba = image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(imgRgba, imgRgba.Bounds(), src, b.Min, draw.Src)
	case *image.RGBA:
		return src
	default:
		logrus.Println("Invalid image format")

	}

	return imgRgba
}

//Reads image from a file
func ReadImage(r io.Reader, extn string) (image.Image, error) {

	switch extn {
	case ".png":
		img, err := png.Decode(r)
		return img, err
	case ".jpg", ".JPEG", ".jpeg", ".JPG":
		img, err := jpeg.Decode(r)
		return img, err
	default:
		return nil, errors.New("not a valid extension")
	}

}

//Writes image to a file
func WriteImage(img *image.RGBA, w io.Writer, extn string) error {

	switch extn {
	case ".png":
		err := png.Encode(w, img)
		return err
	case ".jpg", ".JPEG", ".jpeg", ".JPG":
		err := jpeg.Encode(w, img, &jpeg.Options{Quality: 100})
		return err
	default:
		return errors.New("not a valid extension")
	}
}
