/*
	Ryan Robinson, 2021

	getCertExpiry is a go command line tool to find the expiration dates of given server certs and check if they are expired

	TODO: Fix error message when mixed flag arguments are entered
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func openCert(filename string) (*x509.Certificate, error) {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(d))
	if block == nil {
		return nil, err
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

//Takes a time and returns true if the time has past, false otherwise
func isTimePast(t time.Time) bool {
	delta := time.Since(t)
	return delta > 0
}

//Connects to address and returns 0 if the cert is valid and 1 if it is expired in addition to the cert expiration date. If the server does not support SSL certificates, return 3 and an error.
func getCertExpiry(address string, skipVerify bool, cert tls.Certificate) (int, string, error) {
	conf := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: skipVerify,
	}
	conn, err := tls.Dial("tcp", address, conf)
	if err != nil {
		return 3, "", err
	}
	expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
	if isTimePast(expiry) {
		return 1, expiry.Format("2006-01-02 15:04:05"), nil
	}
	return 0, expiry.Format("2006-01-02 15:04:05"), nil
}

//Parses the first argument for the address and then looks for flags. Currently the only flag is the "insecure" flag which allows for insecure tls connections
func parseArgs(args []string) (string, bool, *tls.Certificate, error) {
	if len(args) < 2 {
		return "", false, nil, errors.New("args error: no args, refer to README for arg format")
	}

	subAdr := flag.NewFlagSet("adr", flag.ExitOnError)

	secureFlag := subAdr.Bool("insecure", false, "Allow insecure tls connections")
	certFile := subAdr.String("cert", "ryan.crt", "A PEM eoncoded certificate file.")
	keyFile := subAdr.String("key", "ryan.key", "A PEM encoded private key file.")

	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}
	subAdr.Parse(os.Args[2:])

	return args[1], *secureFlag, &cert, nil
}

//Exits with status 0 if cert valid and supported, 1 if expired, 3 if not supported, 4 if untrusted (TODO), and 5 if there are no args
func main() {
	args, secure, tlscert, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	status, expiry, err := getCertExpiry(args, secure, *tlscert)

	if err != nil {
		if err.Error() == "x509: certificate signed by unknown authority" && !secure {
			fmt.Println(err.Error())
			os.Exit(4)
		}
		fmt.Printf("Server doesn't support SSL certificate err: " + err.Error())
	}
	fmt.Println(expiry)
	os.Exit(status)
}
