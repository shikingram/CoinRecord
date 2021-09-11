package httpUtil

import (
	"CoinRecord/global"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func HttpGet(url string, params map[string]string, headers map[string]string) ([]byte, error) {
	c := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return []byte{}, errors.New("new request is fail ")
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", global.GlobalObject.ApiKey)
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	fmt.Printf("Go %s URL : %s \n", http.MethodGet, req.URL.String())
	do, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, errors.Wrap(err, "do req failed.")
	}
	defer do.Body.Close()

	bytesBody, err := ioutil.ReadAll(do.Body)
	if err != nil {
		return []byte{}, errors.New("read body failed")
	}

	fmt.Printf("HttpGet receive : %s \n", string(bytesBody))
	return bytesBody, nil
}
