package bundle

import "io"

var _ io.ReadCloser = (*Reader)(nil)

type Reader struct {
	reader io.Reader
}

func (r *Reader) Read(data []byte) (n int, err error) {
	return r.Read(data)
}

func (r *Reader) Close() error {
	_, err := io.Copy(io.Discard, r.reader)

	return err
}
