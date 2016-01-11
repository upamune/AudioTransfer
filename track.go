package main

import (
	"fmt"
	"github.com/wtolson/go-taglib"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Track info
type Track struct {
	Title  string
	Artist string
	Album  string
	Ext    string
	Number int
	Path   string
}

// NewTrack returns Track and error
func NewTrack(fileName string) (Track, error) {
	f, err := taglib.Read(fileName)
	if err != nil {
		log.Fatal("can't " + fileName + " file read")
		return Track{}, err
	}

	fileAbsPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal("can't get " + fileName + "absolute file path")
		return Track{}, err
	}

	t := Track{
		Title:  f.Title(),
		Artist: f.Artist(),
		Album:  f.Album(),
		Number: f.Track(),
		Ext:    filepath.Ext(fileName),
		Path:   fileAbsPath,
	}

	return t, nil
}

// TransferTo transfer track to dir
func (t Track) TransferTo(dir string) error {
	dst := assemblePath(dir, t)
	if err := os.MkdirAll(filepath.Dir(dst), 0777); err != nil {
		log.Fatal("can't create " + filepath.Dir(dst) + "dir")
		return err
	}
	if err := copy(t.Path, dst); err != nil {
		log.Fatal("can't copy " + t.Path + " to " + dst)
		return err
	}

	return nil
}

func assemblePath(dir string, t Track) string {
	return fmt.Sprintf("%s/%s/%s/%02d-%s%s", dir, t.Artist, t.Album, t.Number, t.Title, t.Ext)
}

func copy(src, dst string) error {
	s, err := os.Open(src)
	if err != nil {
		return err
	}

	defer s.Close()
	d, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}
