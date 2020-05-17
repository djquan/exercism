package paasio

import (
	"io"
	"sync"
)

type counter struct {
	nbytes int64
	nops   int
	reader io.Reader
	writer io.Writer
	mutex  sync.Mutex
}

func (c *counter) Read(p []byte) (n int, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if n, err = c.reader.Read(p); err != nil {
		return 0, err
	}

	c.nops++
	c.nbytes += int64(n)
	return
}

func (c *counter) ReadCount() (n int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.nbytes, c.nops
}

func (c *counter) Write(p []byte) (n int, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if n, err = c.writer.Write(p); err != nil {
		return 0, err
	}

	c.nops++
	c.nbytes += int64(n)
	return
}

func (c *counter) WriteCount() (n int64, nops int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.nbytes, c.nops
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &counter{
		writer: writer,
	}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &counter{
		reader: rw,
		writer: rw,
	}
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &counter{
		reader: r,
	}
}
