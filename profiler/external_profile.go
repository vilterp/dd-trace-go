package profiler

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
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

func (p *profiler) collectExternalTypeInferenceProfile() ([]byte, error) {
	log.Info("collectExternalTypeInferenceProfile")
	// TODO: subdirectory?
	entries, err := ioutil.ReadDir("/tmp")
	if err != nil {
		return nil, err
	}
	// TODO: this only captures one txn at a time
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "snoop-txn") && strings.HasSuffix(entry.Name(), ".pb.gz") {
			log.Info("collectExternalTypeInferenceProfile: found matching entry", entry.Name())
			path := "/tmp/" + entry.Name()
			f, err := os.Open(path)
			if err != nil {
				return nil, errors.Wrap(err, "opening file")
			}
			bytes, err := ioutil.ReadAll(f)
			if err != nil {
				return nil, errors.Wrap(err, "reading file")
			}
			if err := os.Remove(path); err != nil {
				return nil, errors.Wrap(err, "deleting file")
			}
			return bytes, nil
		}
	}
	// none found
	return nil, nil
}
