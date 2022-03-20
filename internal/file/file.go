package file

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"regexp"
)

func Write(filename string, data interface{}) error {
	var (
		file []byte
		err  error
	)
	if err = validate(filename); err != nil {
		return err
	}
	if file, err = json.MarshalIndent(data, "", " "); err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, file, 0644); err != nil {
		return err
	}
	return nil
}

func validate(filename string) error {
	var (
		m bool
	)
	r, _ := regexp.Compile("^.+\\.json$")
	if m = r.Match([]byte(filename)); !m {
		return errors.New("filename must not be empty and must end in '.json'")
	}
	return nil
}
