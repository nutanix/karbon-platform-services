package io_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"unicode"

	"github.com/go-yaml/yaml"

	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
)

// PrintTable prints table for the given data and header
func PrintTable(data [][]string, header []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func PrettyPrintJSON(i interface{}) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(i); err != nil {
		glog.Fatalf("failed to print json: %s", err.Error())
	}
}

func PrettyJSONStr(i interface{}) string {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		glog.Fatalf("failed to marshal: %s", err.Error())
	}
	return string(b)
}

// PrettyPrintYaml prints yaml
// TODO: this method does not work with yaml within a yaml. Escape characters
// are double escaped
func PrettyPrintYaml(i interface{}) {
	b, err := yaml.Marshal(i)
	if err != nil {
		glog.Fatalf("failed to marshal: %s", err.Error())
	}
	fmt.Println(string(b))
}

// TrimWhitespaces removes whitespaces from the given string
func TrimWhitespaces(input string) string {
	var buf bytes.Buffer

	for _, r := range input {
		if unicode.IsSpace(r) {
			continue
		}
		buf.WriteRune(r)
	}

	return buf.String()
}

// WriteJSON writes the given JSON to the file given by `filePath`
// It first writes the contents to a temp file and then renames it.
func WriteJSON(filePath string, i interface{}) error {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		glog.Fatalf("failed to marshal json: %s", err.Error())
	}

	glog.V(5).Infof("writing %+v to file %s", i, filePath)
	fileName := filepath.Base(filePath)
	dir := filepath.Dir(filePath)

	f, err := ioutil.TempFile(dir, fileName)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	f.Sync()

	err = os.Rename(f.Name(), path.Join(dir, fileName))
	defer os.RemoveAll(f.Name())
	if err != nil {
		return err
	}
	glog.V(5).Infof("successfully written to %v to %s", i, filePath)
	return nil
}
