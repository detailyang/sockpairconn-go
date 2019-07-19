package sockpairconn

import "testing"

func TestSockPairConn(t *testing.T) {
	sp0, sp1, err := NewSocketPairConn()
	if err != nil {
		t.Fatal(err)
	}

	data := "Hello World"
	n, err := sp1.Write([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	if n != len(data) {
		t.Fatalf("expect got %d length but only %d", len(data), n)
	}

	z := make([]byte, n)
	sp0.Read(z)
	if string(z) != data {
		t.Fatalf("expect read %q but got %q", data, string(z))
	}
}
