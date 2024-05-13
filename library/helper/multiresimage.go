package helper

import (
	"net/url"
	"path"

	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func MultiResImages(media *model.Media) *model.MultiResImage {
	conf := config.NewConfig()
	multiResImage := new(model.MultiResImage)

	if media == nil {
		return multiResImage
	}
	imgURL, err := url.Parse(conf.GetString("image.url"))
	if err != nil {
		return multiResImage
	}
	imgURL.Path = path.Join(imgURL.Path, conf.GetString("image.bucket"), media.Filename)

	multiResImage.MediaUUID = media.UUID
	multiResImage.Original = imgURL.String()

	q := imgURL.Query()

	q.Set("w", "300") // thumbnail width
	imgURL.RawQuery = q.Encode()
	multiResImage.Thumbnail = imgURL.String()

	q.Set("w", "700")
	imgURL.RawQuery = q.Encode()
	multiResImage.Desktop = imgURL.String()

	q.Set("w", "512")
	imgURL.RawQuery = q.Encode()
	multiResImage.Mobile = imgURL.String()

	return multiResImage
}
