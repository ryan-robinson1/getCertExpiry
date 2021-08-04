/*
	Ryan Robinson, 2021

	getCertExpiry is a go command line tool to find the expiration dates of given server certs and check if they are expired

	TODO: Fix error message when mixed flag arguments are entered
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//Takes a time and returns true if the time has past, false otherwise
func isTimePast(t time.Time) bool {
	delta := time.Since(t)
	return delta > 0
}

//Connects to address and returns 0 if the cert is valid and 1 if it is expired in addition to the cert expiration date. If the server does not support SSL certificates, return 3 and an error.
func getCertExpiry(address string, skipVerify bool, certflag bool, certs []string) (int, string, error) {
	conf := &tls.Config{
		InsecureSkipVerify: skipVerify,
	}
	if certflag {
		cert, err := tls.LoadX509KeyPair(certs[3], certs[2])
		if err != nil {
			log.Fatal(err)
		}
		caCert, err := ioutil.ReadFile(certs[1])
		if err != nil {
			log.Fatal(err)
		}
		users, err := ioutil.ReadFile(certs[0])
		if err != nil {
			log.Fatal(err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		caCertPool.AppendCertsFromPEM(users)

		conf = &tls.Config{
			Certificates:       []tls.Certificate{cert},
			RootCAs:            caCertPool,
			ClientCAs:          caCertPool,
			InsecureSkipVerify: skipVerify,
		}

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
func parseArgs(args []string) (string, bool, bool, []string, error) {
	if len(args) < 2 {
		return "", false, false, nil, errors.New("args error: no args, refer to README for arg format")
	}

	subAdr := flag.NewFlagSet("adr", flag.ExitOnError)
	secureFlag := subAdr.Bool("insecure", false, "Allow insecure tls connections")
	certFlag := subAdr.Bool("certs", false, "Connection requires certs")

	subAdr.Parse(os.Args[2:])

	//if cert flag, set the last 4 args to be certs and return an array of those certs
	if *certFlag {
		if len(os.Args) == 7 || *secureFlag && len(os.Args) == 8 {
			var certs []string
			for i := 1; i <= 4; i++ {
				certs = append(certs, os.Args[len(os.Args)-i])
			}
			return args[1], *secureFlag, *certFlag, certs, nil
		}
		return "", false, false, nil, errors.New("args error: cert args incorrect, refer to README for arg format")
	}
	return args[1], *secureFlag, *certFlag, nil, nil
}

//Exits with status 0 if cert valid and supported, 1 if expired, 3 if not supported, 4 if untrusted (TODO), and 5 if there are no args
func main() {
	args, secure, certFlag, certs, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	status, expiry, err := getCertExpiry(args, secure, certFlag, certs)

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
