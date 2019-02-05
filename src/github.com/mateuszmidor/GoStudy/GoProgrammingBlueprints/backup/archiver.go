package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	Archive(src, dest string) error
	DestFmt() string
}

type zipper struct{}

var ZIP Archiver = (*zipper)(nil)

func (z *zipper) DestFmt() string {
	return "%d.zip"
}
func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
		return nil
	})

}
