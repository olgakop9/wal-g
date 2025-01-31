package lzma

import (
	"io"
	"io/ioutil"

	"github.com/ulikunitz/xz/lzma"
	"github.com/wal-g/wal-g/internal/compression/computils"
)

type Decompressor struct{}

func (decompressor Decompressor) Decompress(src io.Reader) (io.ReadCloser, error) {
	lzReader, err := lzma.NewReader(computils.NewUntilEOFReader(src))
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(lzReader), nil
}

func (decompressor Decompressor) FileExtension() string {
	return FileExtension
}
