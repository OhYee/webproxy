package proxy

import (
	"context"
	"io"
)

const (
	bufferSize = 512
)

func redirectIO(ctx context.Context, src io.Reader, dst io.Writer) (int, error) {
	if src == nil {
		return 0, nil
	}

	buf := make([]byte, bufferSize)
	total := 0
	for {
		select {
		case <-ctx.Done():
			return total, nil
		default:
			rn, err := src.Read(buf)
			if err != nil {
				return total, err
			}
			written := 0
			for written < rn {
				select {
				case <-ctx.Done():
					return total, nil
				default:
					wn, err := dst.Write(buf[written:rn])
					if err != nil {
						return total, err
					}

					written += wn
					total += wn
				}
			}
		}
	}
}
