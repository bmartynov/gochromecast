package service

import (
	"io"
	"os"
	"github.com/anacrolix/torrent"
)

type SeekAbleContent interface {
	io.Closer
	io.ReadSeeker
}

type FileEntry struct {
	File *torrent.File
	*torrent.Reader
}

func (f FileEntry) Seek(offset int64, whence int) (int64, error) {
	return f.Reader.Seek(offset + f.File.Offset(), whence)
}

func NewReader(t torrent.Torrent, tf torrent.File) (SeekAbleContent, error) {
	var readahead = tf.Length() / 100

	reader := t.NewReader()
	reader.SetReadahead(readahead)
	reader.SetResponsive()
	_, err := reader.Seek(tf.Offset(), os.SEEK_SET)

	return &FileEntry{File: &tf, Reader: reader}, err
}
