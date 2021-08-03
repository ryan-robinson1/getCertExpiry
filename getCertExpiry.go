/*
	Ryan Robinson, 2021

	getCertExpiry is a go command line tool to find the expiration dates of given server certs and check if they are expired
*/

package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"time"
)

//Takes a time and returns true if the time has past, false otherwise
func isTimePast(t time.Time) bool {
	delta := time.Since(t)
	return delta > 0
}

//Connects to address and returns 0 if the cert is valid and 1 if it is expired in addition to the cert expiration date. If the server does not support SSL certificates, return 3 and an error.
func getCertExpiry(address string) (int, string, error) {
	conn, err := tls.Dial("tcp", address, nil)
	if err != nil {
		return 3, "", err
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	if isTimePast(expiry) {
		return 1, expiry.Format("2006-01-02 15:04:05"), nil
	}
	return 0, expiry.Format("2006-01-02 15:04:05"), nil

}

func parseArgs(args []string) (string, error) {
	if len(args) != 2 {
		return "", errors.New("args error: args formatted incorrectly, refer to README for correct formatting")
	}
	return args[1], nil
}

//Exits with status 0 if cert valid and supported, 1 if expired, 3 if not supported, 4 if untrusted, and 5 if the args are incorrectly formatted
func main() {
	args, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	status, expiry, err := getCertExpiry(args)
	if err != nil {
		fmt.Printf("Server doesn't support SSL certificate err: " + err.Error())
	}
	fmt.Println(expiry)
	os.Exit(status)
}
