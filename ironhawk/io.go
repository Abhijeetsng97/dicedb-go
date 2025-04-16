package ironhawk

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/dicedb/dicedb-go/wire"
	"google.golang.org/protobuf/proto"
)

const (
	maxRequestSize = 32 * 1024 * 1024 // 32 MB
	ioBufferSize   = 16 * 1024        // 16 KB
	idleTimeout    = 30 * time.Minute
)

func Read(conn net.Conn) (*wire.Result, error) {
	var result []byte
	reader := bufio.NewReaderSize(conn, ioBufferSize)
	buf := make([]byte, ioBufferSize)

	for {
		n, err := reader.Read(buf)
		if n > 0 {
			if len(result)+n > maxRequestSize {
				return nil, fmt.Errorf("request too large")
			}
			result = append(result, buf[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if n < len(buf) {
			break
		}
	}

	if len(result) == 0 {
		return nil, io.EOF
	}

	r := &wire.Result{}
	if err := proto.Unmarshal(result, r); err != nil {
		return nil, fmt.Errorf("failed to unmarshal command: %w", err)
	}
	return r, nil
}

func Write(conn net.Conn, c *wire.Command) error {
	data, err := proto.Marshal(c)
	if err != nil {
		return err
	}
	_, err = conn.Write(data)
	return err
}
