package ws

import (
	"context"
	"strings"
	// "time"
)

type Client struct {
	conn *Conn
}

func (c *Client) Close() {
	if c == nil || c.conn == nil {
		return
	}

	c.conn.Close()
}

func (c *Client) WaitServerConnect(ctx context.Context, namespace string) (*NSConn, error) {
	return c.conn.WaitConnect(ctx, namespace)
}

func (c *Client) Connect(ctx context.Context, namespace string) (*NSConn, error) {
	return c.conn.Connect(ctx, namespace)
}

type Dialer func(ctx context.Context, url string) (Socket, error)

// Dial establish a new websocket client.
// Context "ctx" is used for handshake timeout.
func Dial(dial Dialer, ctx context.Context, url string, connHandler ConnHandler) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if !strings.HasPrefix(url, "ws://") && !strings.HasPrefix(url, "wss://") {
		url = "ws://" + url
	}

	underline, err := dial(ctx, url)
	if err != nil {
		return nil, err
	}

	if connHandler == nil {
		connHandler = Namespaces{}
	}

	c := newConn(underline, connHandler.getNamespaces())
	readTimeout, writeTimeout := getTimeouts(connHandler)
	c.readTimeout = readTimeout
	c.writeTimeout = writeTimeout

	if err = c.clientAck(); err != nil {
		return nil, err
	}

	go c.startReader()

	return &Client{conn: c}, nil
}
