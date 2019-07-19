package sockpairconn

import (
	"net"
	"syscall"
	"time"
)

// SocketPairConn represents a connection by socketpair.
// It implements the net.Conn interface.
type SocketPairConn struct {
	fd   int
	conn net.Conn
}

// NewSocketPairConn returns the socketpairt connection.
func NewSocketPairConn() (*SocketPairConn, *SocketPairConn, error) {
	fds, err := syscall.Socketpair(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, nil, err
	}

	pc0, err := newFileConn(fds[0])
	if err != nil {
		return nil, nil, err
	}

	pc1, err := newFileConn(fds[1])
	if err != nil {
		pc0.Close()
		return nil, nil, err
	}

	return &SocketPairConn{fd: fds[0], conn: pc0}, &SocketPairConn{fd: fds[1], conn: pc1}, nil
}

// Read reads data from the connection.
// Read can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetReadDeadline.
func (pc *SocketPairConn) Read(b []byte) (n int, err error) {
	return pc.conn.Read(b)
}

// Write writes data to the connection.
// Write can be made to time out and return an Error with Timeout() == true
// after a fixed time limit; see SetDeadline and SetWriteDeadline.
func (pc *SocketPairConn) Write(b []byte) (n int, err error) {
	return pc.conn.Write(b)
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (pc *SocketPairConn) Close() error {
	return pc.conn.Close()
}

// LocalAddr returns the local network address.
func (pc *SocketPairConn) LocalAddr() net.Addr {
	return pc.conn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (pc *SocketPairConn) RemoteAddr() net.Addr {
	return pc.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail with a timeout (see type Error) instead of
// blocking. The deadline applies to all future and pending
// I/O, not just the immediately following call to Read or
// Write. After a deadline has been exceeded, the connection
// can be refreshed by setting a deadline in the future.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (pc *SocketPairConn) SetDeadline(t time.Time) error {
	return pc.conn.SetDeadline(t)
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (pc *SocketPairConn) SetReadDeadline(t time.Time) error {
	return pc.conn.SetReadDeadline(t)
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (pc *SocketPairConn) SetWriteDeadline(t time.Time) error {
	return pc.conn.SetWriteDeadline(t)
}
