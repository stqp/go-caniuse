package datasrc

import (
	"io"
	"io/ioutil"
	l "log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/stqp/go-caniuse/pkg/utils"
)

var (
	Data     []byte
	Filepath string
)

func init() {

	Filepath = filepath.Join(utils.CacheDir(), "data.json")

	data, err := Update(false)
	if err != nil {
		l.Fatal("Failed to update data source file")
		return
	}
	Data = data
}

func Update(force bool) (data []byte, err error) {

	exists, err := utils.Exists(Filepath)
	if exists && force {
		if err := os.Remove(Filepath); err != nil {
			l.Fatal("Failed to delete old data source file")
			l.Fatal(err)
			return nil, err
		}
	}

	if !exists || force {
		if err := downloadAndReplace(); err != nil {
			l.Fatal(err)
			return nil, err
		}
	}
	return GetData()
}

func downloadAndReplace() (err error) {

	resp, err := http.Get("https://raw.githubusercontent.com/Fyrd/caniuse/master/fulldata-json/data-2.0.json")
	if err != nil {
		l.Fatal("Failed to download data source file")
		l.Fatal(err)
		return err
	}
	defer resp.Body.Close()

	if err = os.MkdirAll(filepath.Dir(Filepath), 0700); err != nil {
		return err
	}

	out, err := os.Create(Filepath)
	if err != nil {
		l.Fatal("Failed to create data source file")
		l.Fatal(err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		l.Fatal("Failed to write data source file")
		l.Fatal(err)
		return err
	}
	return nil
}

func GetData() (data []byte, err error) {
	data, err = ioutil.ReadFile(Filepath)
	if err != nil {
		l.Fatal("Failed to read source data")
		l.Fatal(err)
		return nil, err
	}
	Data = data
	return data, nil
}
