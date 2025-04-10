package commands

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func (state *State) CallObsidian(path string, header map[string]string) (io.ReadCloser, error) {
	// fmt.Printf("Calling %s ...\n", path)
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
	req.Header.Set("Authorization", "Bearer "+state.ApiKey)
	for key, value := range header {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	return res.Body, nil
}
