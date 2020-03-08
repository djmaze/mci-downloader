package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type WADMClient struct {
	IP   string
	Port uint
}

func WADM(ip string, port uint) WADMClient {
	return WADMClient{IP: ip, Port: port}
}

func (wadm WADMClient) getTracks() []byte {
	url := fmt.Sprintf("http://%s:%d", wadm.IP, wadm.Port)
	payload := bytes.NewBuffer(
		[]byte("<requestplayabledata><nodeid>385875968</nodeid><numelem>0</numelem><fromindex>0</fromindex></requestplayabledata>\r\n"),
	)
	resp, err := http.Post(url, "text/xml", payload)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (wadm WADMClient) downloadTrack(url string) (error, []byte) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err, nil
		}
		return nil, body
	} else {
		return errors.New("Got response " + resp.Status), nil
	}
}
