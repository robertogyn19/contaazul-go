package contaazul

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	UnmarshalResponseError = errors.New("could not unmarshal response")
)

type RequestParams struct {
	cli              *Client
	method           string
	url              string
	params           []byte
	target           interface{}
	expectedHttpCode []int
	compress         bool
}

func executeApiRequest(erp RequestParams) error {
	cli := erp.cli
	params := erp.params

	var buffer bytes.Buffer
	if params == nil {
		params = []byte{}
	}

	if erp.compress {
		g := gzip.NewWriter(&buffer)
		_, err := g.Write(params)
		if err != nil {
			return err
		}

		err = g.Close()
		if err != nil {
			return err
		}
	} else {
		buf := bytes.NewBuffer(params)
		buffer = *buf
	}

	req, err := http.NewRequest(erp.method, erp.url, &buffer)

	if err == nil && erp.compress {
		req.Header.Set("Accept-Encoding", "gzip")
		req.Header.Set("Content-Encoding", "gzip")
	}

	if err != nil {
		log.Printf("could not create request: %v", err)
		return err
	}

	err = cli.doRequest(req, erp.expectedHttpCode, erp.target, false)
	return err
}

func (cli *Client) doRequest(req *http.Request, status []int, target interface{}, refreshed bool) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cli.accessToken))
	resp, err := cli.client.Do(req)

	if err != nil {
		body := ""
		if resp != nil && resp.Body != nil {
			b, _ := ioutil.ReadAll(resp.Body)
			body = string(b)
		}
		log.Printf("could not execute request, body: %s, error: %v", body, err)
		return err
	}

	statusCode := resp.StatusCode
	if checkStatus(status, statusCode) {
		if target != nil {
			encoding := resp.Header.Get("Content-Encoding")
			var reader io.Reader

			if encoding == "gzip" {
				gz, err := gzip.NewReader(resp.Body)
				if err != nil {
					return UnmarshalResponseError
				}

				defer gz.Close()

				rawData, err := ioutil.ReadAll(gz)

				if err != nil {
					return UnmarshalResponseError
				}

				reader = bytes.NewBuffer(rawData)
			} else {
				defer resp.Body.Close()
				reader = resp.Body
			}

			body, err := unmarshalBody(reader, target)

			if err != nil {
				log.Printf("(%d) could not unmarshal response body: %s, error: %v", statusCode, body, err)
				return UnmarshalResponseError
			}
		}
	} else if statusCode == http.StatusUnauthorized {
		if refreshed {
			msg := "unauthorized request even with a refreshed token"
			log.Println(msg)
			return fmt.Errorf(msg)
		} else {
			err = refreshToken(cli)
			if err != nil {
				return err
			}

			return cli.doRequest(req, status, target, true)
		}
	} else {
		msg := fmt.Sprintf("unexpected response status code, expected: %v, actual: %d", status, statusCode)
		log.Println(msg)
		err = fmt.Errorf(msg)
	}

	return err
}

func unmarshalBody(reader io.Reader, target interface{}) ([]byte, error) {
	body, err := ioutil.ReadAll(reader)

	if err != nil {
		log.Printf("could not unmarshal body: %s, error: %v", body, err)
		return body, err
	}

	return body, json.Unmarshal(body, target)
}

func checkStatus(codes []int, status int) bool {
	for _, s := range codes {
		if s == status {
			return true
		}
	}

	return false
}
