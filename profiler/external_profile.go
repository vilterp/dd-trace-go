package profiler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/internal/log"
)

func (p *profiler) collectExternalProfile() ([]byte, error) {
	log.Info("requesting profile")
	// TODO: make URL configurable
	resp, err := http.Get("http://localhost:19465/profile?duration=10")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Info("got profile. err:", err)
	return body, err
}
