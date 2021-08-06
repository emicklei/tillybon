package main

import (
	"bytes"
	"context"
	"log"

	"bazil.org/fuse"
)

// File implements both Node and Handle for the hello file.
type File struct{}

const greeting = "hello, world\n"

func (File) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = 2
	a.Mode = 0444 // -rw-rwxr--
	a.Size = uint64(len(greeting))
	return nil
}

func (File) ReadAll(ctx context.Context) ([]byte, error) {
	return []byte(greeting), nil
}

func (File) Write(ctx context.Context, req *fuse.WriteRequest, resp *fuse.WriteResponse) error {
	var buf bytes.Buffer
	n, err := buf.Write(req.Data)
	if err != nil {
		return err
	}
	resp.Size = n
	log.Println(req.Offset, n, buf.String(), req.Data)
	return nil
}
