package server

import (
	"encoding/json"
	"net"
	"radiant/protocol"
)

type Server struct {
	listener net.Listener
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}

	server := &Server{
		listener: listener,
	}

	go server.run()

	return server, nil
}

func (s *Server) run() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			// handle error
		}

		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)

	for {
		var msg protocol.Message
		err := decoder.Decode(&msg)
		if err != nil {
			// handle error
			return
		}

		switch msg.Type {
		case protocol.MessageTypeSendMessage:
			// handle send message
		case protocol.MessageTypeCreateGroup:
			// handle create group
		case protocol.MessageTypeJoinGroup:
			// handle join group
		case protocol.MessageTypeLeaveGroup:
			// handle leave group
		case protocol.MessageTypeGroupMessage:
			// handle group message
		default:
			// handle unknown message type
		}
	}
}
