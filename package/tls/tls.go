package tls

import (
	"fmt"

	"google.golang.org/grpc/credentials"
)

const (
	serverCertFile = "server.crt"
	serverKeyFile  = "server.key"
	clientCertFile = "ca.crt"
)

func GenerateServerTLSCreds(certPath string) (credentials.TransportCredentials, error) {
	certFile := fmt.Sprintf("%s%s", certPath, serverCertFile)
	keyFile := fmt.Sprintf("%s%s", certPath, serverKeyFile)

	return credentials.NewServerTLSFromFile(certFile, keyFile)
}

func GenerateClientTLSCreds(certPath string) (credentials.TransportCredentials, error) {
	certFile := fmt.Sprintf("%s%s", certPath, clientCertFile)
 
	return credentials.NewClientTLSFromFile(certFile, "")
 }
