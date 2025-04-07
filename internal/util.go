package internal

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	localCertFile = "obsidian-local-rest-api.crt"
)

func Call(path string) (io.ReadCloser, error) {
	fmt.Printf("Calling %s ...\n", path)
	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := os.ReadFile(localCertFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", localCertFile, err)
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	url := "https://127.0.0.1:27124" + path
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer f76d3fa69bf1e7d1eb568d21ba839ffa7e5bfae7e5b18efe2bd224c1b8e466c2")
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	return res.Body, nil
}

func ToDict(reader io.ReadCloser) (map[string]interface{}, error) {
	var result map[string]interface{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
