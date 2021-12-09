package util

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

var (
	defaultEnvFile = ".env.local"
)

// LoadEnvForced loads .env.local from current dir into go env, will
// overwrite any values set in the environment prior to parsing file
func LoadEnvForced() error {
	return LoadEnvFileForced(defaultEnvFile)
}

// LoadEnvForced loads file at `path` into go env, will
// overwrite any values set in the environment prior to parsing file
func LoadEnvFileForced(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parseIntoEnv(data, true)
	return nil

}

// LoadEnv loads .env.local from current dir into go env, does not
// load a key if it already exists in the env
func LoadEnv() error {
	return LoadEnvFile(defaultEnvFile)
}

// LoadEnvFile loads file at `path` into go env, does not
// load a key if it already exists in the env
func LoadEnvFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	parseIntoEnv(data, false)
	return nil

}

func parseIntoEnv(data []byte, overwrite bool) {
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, "=")
		if len(s) != 2 {
			continue
		}

		k := s[0]
		v := s[1]

		if overwrite {
			os.Setenv(k, v)
			continue
		}

		if os.Getenv(k) == "" {
			os.Setenv(k, v)
		}
	}
}
