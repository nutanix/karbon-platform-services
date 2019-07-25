package io_utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"xi-iot-cli/xi-iot/errutils"

	"github.com/go-yaml/yaml"
)

func Readfile(filepath string) ([]byte, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func ReadJSON(filePath string, i interface{}) error {
	b, err := Readfile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, i)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json. %s", err.Error())
	}
	return nil
}

func ReadYaml(filepath string) ([]interface{}, *errutils.XiErr) {
	res := []interface{}{}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, errutils.NewIOErr(err.Error())
	}
	dec := yaml.NewDecoder(f)
	i := make(map[interface{}]interface{})
	for dec.Decode(&i) == nil {
		res = append(res, i)
		i = make(map[interface{}]interface{})
	}
	return res, nil
}

func Readstdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}
