package configuration

import (
	"gopkg.in/yaml.v3"
	"io"
	"time"
)

// ParseRequests takes Reader, and uses yaml library to decode file into Requests struct
func ParseRequests(file io.Reader) (*Requests, error) {
	dec := yaml.NewDecoder(file)
	dec.KnownFields(true)
	r := &Requests{}
	err := dec.Decode(r)
	if err != nil {
		return nil, err
	}
	r.Timeout, err = time.ParseDuration(r.RawTimeout)
	if err != nil {
		return nil, err
	}
	//r.DataFilter.Delay, err = time.ParseDuration(r.DataFilter.RawDelay)
	//if err != nil {
	//	return nil, err
	//}
	return r, nil
}
