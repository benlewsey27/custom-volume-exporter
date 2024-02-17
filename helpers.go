package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type VolumeData struct {
	Namespace string
	PersistentVolumeClaim string
	PersistentVolume string
	Usage int
}

func sendResponse(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	w.Write([]byte(body))
}

func getAllFiles(dirPath string) ([]string, error) {
	var files []string

	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fullPath := fmt.Sprintf("%s/%s", dirPath, entry.Name())

		if entry.IsDir() {
			subdirFiles, err := getAllFiles(fullPath)
			if err != nil {
				return nil, err
			}
			files = append(files, subdirFiles...)
		} else {
			files = append(files, fullPath)
		}
	}

	return files, nil
}

func getVolumeSize(filepath string) (int, error) {
	totalSize := 0
	
	files, err := getAllFiles(filepath)
	if err != nil {
		return 0, err
	}

	for _, file := range files{
		fileInfo, err := os.Stat(file)
    if err != nil {
			return 0, err
    }

		totalSize = totalSize + int(fileInfo.Size())
	}
	
	return totalSize, nil
}

func getVolumeData(filepath string) ([]VolumeData, error) {
	var out []VolumeData

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filename := file.Name()

		namespace := strings.Split(filename, "-")[0]
		persistentVolumeClaim := strings.ReplaceAll(strings.Split(filename, "-pvc")[0], fmt.Sprintf("%s-", namespace), "")
		persistentVolume := fmt.Sprintf("pvc-%s", strings.Split(filename, "pvc-")[1])
		usage, err := getVolumeSize(fmt.Sprintf("%s/%s", filepath, filename))
		if err != nil {
			return nil, err
		}

		tmp := VolumeData{Namespace: namespace, PersistentVolume: persistentVolume, PersistentVolumeClaim: persistentVolumeClaim, Usage: usage}
		out = append(out, tmp)
	}

	return out, nil
}
