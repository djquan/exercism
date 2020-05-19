package paasio

import (
	"io"
	"sync"
)

type counter struct {
	sync.RWMutex
	nbytes int64
	nops   int
	reader io.Reader
	writer io.Writer
}

func (c *counter) Read(p []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()

	if n, err = c.reader.Read(p); err != nil {
		return 0, err
	}

	c.nops++
	c.nbytes += int64(n)
	return
}

func (c *counter) ReadCount() (n int64, nops int) {
	c.RLock()
	defer c.RUnlock()

	return c.nbytes, c.nops
}

func (c *counter) Write(p []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()

	if n, err = c.writer.Write(p); err != nil {
		return 0, err
	}

	c.nops++
	c.nbytes += int64(n)
	return
}

func (c *counter) WriteCount() (n int64, nops int) {
	c.RLock()
	defer c.RUnlock()

	return c.nbytes, c.nops
}

//NewWriteCounter generates a NewWriteCounter, an object that tracks bytes written
func NewWriteCounter(writer io.Writer) WriteCounter {
	return &counter{
		writer: writer,
	}
}

//NewReadWriteCounter generates a NewReadWriteCounter, an object that tracks bytes written and read
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &counter{
		reader: rw,
		writer: rw,
	}
}

//NewReadCounter generates a NewReadCounter, an object that tracks bytes read.
func NewReadCounter(r io.Reader) ReadCounter {
	return &counter{
		reader: r,
	}
}
