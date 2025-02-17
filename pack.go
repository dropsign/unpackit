package packutil

// import (
// 	"archive/tar"
// 	"errors"
// 	"io"
// 	"os"
// 	"strings"

// 	"github.com/ulikunitz/xz"
// )

// type Filer interface {
// 	AddFile(filename string) error
// }

// type TarWriter struct {
// 	tarWriter         *tar.Writer
// 	compressionWriter io.WriteCloser
// }

// func (tw *TarWriter) WriteFile(file os.FileInfo) error {
// 	f, err := os.Open(file.Name())
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	header := &tar.Header{
// 		Name: filename,
// 	}

// 	err = tw.tarWriter.WriteHeader(header)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = io.Copy(tw.tarWriter, f)

// 	return err
// }

// func (tw *TarWriter) WriteBytes(header *tar.Header, p []byte) (int, error) {
// 	if err := tw.tarWriter.WriteHeader(header); err != nil {
// 		return 0, err
// 	}

// 	return tw.tarWriter.Write(p)
// }

// func (tw *TarWriter) Close() error {
// 	err := tw.tarWriter.Close()

// 	if tw.compressionWriter != nil {
// 		err = errors.Join(tw.compressionWriter.Close())
// 	}

// 	return err
// }

// func NewTarWriter(writer io.Writer, format string) (*TarWriter, error) {
// 	var (
// 		w   io.WriteCloser
// 		cw  io.WriteCloser
// 		err error
// 	)

// 	switch strings.ToLower(format) {
// 	case "xz":
// 		cw, err = xz.NewWriter(writer)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if cw != nil {
// 		w = tar.NewWriter(cw)
// 	} else {
// 		w = tar.NewWriter(writer)
// 	}

// 	reader := &TarWriter{
// 		tarWriter:         w,
// 		compressionWriter: cw,
// 	}

// 	return reader, nil
// }
