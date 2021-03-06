package tcp.middleware

import (
	"net"
	"time"
)

type Tcp struct {
	conn    net.Conn
	timeout time.Duration
	//	remtoeAddr string
}

func (tcp *Tcp) RemoteAddr() string {
	return tcp.conn.RemoteAddr().String()
	//	return tcp.remtoeAddr
}

func (tcp *Tcp) Run(c *Context) {
	defer tcp.close()
	defer c.onConnectionClosed(c)

	if c.onConnectionOpen != nil {
		c.onConnectionOpen(c)
		return
	}

	for !c.IsAborted() {
		if _, err := c.Recv(); err != nil {
			c.AbortWithError(err)
			return
		}

		c.onNewMessage(c)
	}
}

func (tcp *Tcp) Recv(cache []byte) (int, error) {
	tcp.conn.SetReadDeadline(time.Now().Add(tcp.timeout))
	return tcp.conn.Read(cache)
}

func (tcp *Tcp) Send(data []byte) (int, error) {
	tcp.conn.SetWriteDeadline(time.Now().Add(tcp.timeout))
	return tcp.conn.Write(data)
}

func (tcp *Tcp) close() error {
	return tcp.conn.Close()
}
