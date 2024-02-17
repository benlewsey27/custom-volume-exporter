package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var filepath string

func serveMetrics(w http.ResponseWriter, r *http.Request) {
	var body bytes.Buffer

	body.WriteString("# HELP pvc_volume_used_bytes Persistent volume size in bytes.\n")
	body.WriteString("# TYPE pvc_volume_used_bytes gauge\n")

	volumeData, err := getVolumeData(filepath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		sendResponse(w, 500, "Internal Server Error\n")
		return
	}

	for _, data := range volumeData {
		tmp := fmt.Sprintf("pvc_volume_used_bytes{name=\"%s\", pvc=\"%s\", namespace=\"%s\"} %d\n", data.PersistentVolumeClaim, data.PersistentVolume, data.Namespace, data.Usage)
		body.WriteString(tmp)
	}

	sendResponse(w, 200, body.String())
}

func main() {
	http.HandleFunc("/metrics", serveMetrics)

	flag.StringVar(&filepath, "filepath", "", "root directory to get volume data")
	flag.Parse()

	if filepath == "" {
		fmt.Println("Error: Missing filepath. Please provide a filepath via the --filepath flag.")
		os.Exit(1)
	}

	fmt.Printf("Running on port 8000...\n")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}
