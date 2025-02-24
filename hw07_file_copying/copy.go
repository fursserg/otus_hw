package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var (
		total, over int
		writeOffset int64
	)

	source, err := os.Open(fromPath) // открываем файл (не забыть про err!)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	fileInfo, err := source.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	defer func() {
		if closeErr := source.Close(); closeErr != nil {
			fmt.Printf("failed to close source file: %v\n", closeErr)
		}
	}()

	if offset > fileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}

	destination, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func() {
		if closeErr := destination.Close(); closeErr != nil {
			fmt.Printf("failed to close destination file: %v\n", closeErr)
		}
	}()

	if limit == 0 || limit > fileInfo.Size() || limit+offset > fileInfo.Size() {
		limit = fileInfo.Size() - offset
	}

	step := fileInfo.Size() / 100
	pb := progressBar{BarLength: 50, Total: limit}

	for {
		// буфер перезатираю в каждом цикле, чтобы в памяти не держать файл целиком
		buf := make([]byte, step)
		read, readErr := source.ReadAt(buf, offset)
		total += read

		if total > int(limit) {
			over = total - int(limit)
		}

		if readErr != nil && readErr != io.EOF {
			return fmt.Errorf("failed to read: %w", readErr)
		}

		written, err := destination.WriteAt(buf[:read-over], writeOffset)
		if err != nil {
			return fmt.Errorf("failed to write: %w", err)
		}

		pb.Add(written)

		if readErr == io.EOF || total >= int(limit) {
			break
		}

		offset += int64(read)
		writeOffset += int64(read)
	}

	pb.End()

	return nil
}
