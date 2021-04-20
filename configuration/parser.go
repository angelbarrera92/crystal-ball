package configuration

import (
	"gopkg.in/yaml.v3"
	"io"
)

func ParseFeeds(file io.Reader) (*Feeds, error) {
	dec := yaml.NewDecoder(file)
	dec.KnownFields(true)
	f := &Feeds{}
	err := dec.Decode(f)
	return f, err
}
