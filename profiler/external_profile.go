package profiler

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/internal/log"
)

func (p *profiler) collectExternalCPUProfile() ([]byte, error) {
	log.Info("requesting CPU profile")
	// TODO: make URL configurable
	resp, err := http.Get("http://localhost:19465/profile?duration=10")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Info("got CPU profile. err:", err)
	return body, err
}

func (p *profiler) collectExternalAllocProfile() ([]byte, error) {
	log.Info("requesting alloc profile")
	// TODO: make URL configurable
	resp, err := http.Get("http://localhost:19465/allocs_profile?duration=10&sample_rate=0.00001")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Info("got alloc profile. err:", err)
	return body, err
}
