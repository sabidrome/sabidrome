package extractor

import (
	"archive/zip"
	"errors"
)

func OpenEpub(path string) (*Epub, error) {

	var epub *Epub = new(Epub)
	var err error

	epub.ReadCloser, err = zip.OpenReader(path)
	if err != nil {
		return nil, errors.New("open zip file: " + err.Error())
	}

	c, err := epub.Container()
	if err != nil {
		return nil, errors.New("open container file: " + err.Error())
	}
}
