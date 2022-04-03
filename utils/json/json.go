package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// the path of directory "data"
var root string

func From(s any) ([]byte, error) {
	return json.Marshal(s)
}

func Parse[T any](b []byte) (t T, err error) {
	err = json.Unmarshal(b, &t)
	return
}

func ParseWithDft(b []byte, t any) error {
	return json.Unmarshal(b, t)
}

func InitConfig(configJson, dftHost, dftPort string) (host, port string) {
	config := struct{ Host, Port, Dir string }{dftHost, dftPort, ""}
	err := ParseWithDft([]byte(configJson), &config)
	if err != nil {
		panic(fmt.Errorf("Config JSON file is invalid:%w", err))
	}
	if config.Dir == "" {
		panic(errors.New("the value of dir in Config JSON file is empty!"))
	}
	root = path.Join(config.Dir, "data")
	if !isExist(root) {
		panic(errors.New("the data dir defined in Config JSON file is not exist in your filesystem:" + config.Dir))
	}
	fmt.Printf("The data dir is:%v\n", root)

	clientJson, err := From(config)
	if err != nil {
		panic(fmt.Errorf("can not marshal clientJson:%w", err))
	}
	write(clientJson, path.Join(config.Dir, "client/src/config.json"), false)

	return config.Host, config.Port
}

func ReadFile[T any](pathSegments ...string) (res T, err error) {
	path := join(pathSegments)
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	return Parse[T](data)
}

func ReadDir[T any](pathSegments ...string) (res []T, err error) {
	err = fs.WalkDir(os.DirFS(join(pathSegments)), ".", func(p string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			err = walkErr
			return walkErr
		}
		if filepath.Ext(p) == ".json" {
			v, err := ReadFile[T](append(pathSegments, p)...)
			if err != nil {
				return err
			}
			res = append(res, v)
		}
		return nil
	})
	return
}

func write(data []byte, filename string, append bool) (err error) {
	if !append && isExist(filename) {
		os.Remove(filename)
	}
	dir := path.Dir(filename)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return
	}

	defer f.Close()

	_, err = f.Write(data)
	return
}

func WriteAppend(data []byte, pathSegments ...string) (err error) {
	return write(data, join(pathSegments), true)
}

func WriteNew(data []byte, pathSegments ...string) (err error) {
	return write(data, join(pathSegments), false)
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func join(pathSegments []string) string {
	return path.Join(root, path.Join(pathSegments...))
}
