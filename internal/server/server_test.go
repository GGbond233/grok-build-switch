package server

import "testing"

func TestListenRejectsInvalidPreferredPort(t *testing.T) {
	server := &Server{}
	if _, _, err := server.Listen(70000); err == nil {
		t.Fatal("Listen() accepted an invalid preferred port")
	}
}
