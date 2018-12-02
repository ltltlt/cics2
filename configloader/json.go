package configloader

// This file include loader which load json config

import (
	"encoding/json"
	"io/ioutil"
)

type JSON struct{}

func (loader *JSON) Load(file string) error {
	bts, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bts, &config); err != nil {
		return err
	}
	return nil
}
