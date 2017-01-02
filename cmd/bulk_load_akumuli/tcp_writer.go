package main

import (
	"net"
	"time"
)

// LineProtocolWriter is the interface used to write Akumuli bulk data.
type LineProtocolWriter interface {
	// WriteLineProtocol writes the given byte slice containing bulk data
	// to an implementation-specific remote server.
	// Returns the latency, in nanoseconds, of executing the write against the remote server and applicable errors.
	// Implementers must return errors returned by the underlying transport but are free to return
	// other, context-specific errors.
	WriteLineProtocol([]byte) (latencyNs int64, err error)
}

// TCPWriter is a Writer that writes to an Akumuli TCP server.
type TCPWriter struct {
	client *net.TCPConn
}

// NewTCPWriter returns a new TCPWriter from the supplied TCPWriterConfig.
func NewTCPWriter(host string) (LineProtocolWriter, error) {
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return nil, err
	}
	c, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	return &TCPWriter{
		client: c,
	}, err
}

// WriteLineProtocol writes the given byte slice to the TCP socket
// It returns the latency in nanoseconds and any error received while sending the data over the TCP
func (w *TCPWriter) WriteLineProtocol(body []byte) (int64, error) {
	start := time.Now()
	_, err := w.client.Write(body)
	lat := time.Since(start).Nanoseconds()
	return lat, err
}
