package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"errors"
	
	"github.com/gorilla/mux"

)

type CommandLineArguments struct{
	Port string
	MapsFolder string
}

type OfflineMapParametersStruct struct{
	rootFolder string
	layerJsonPath string
	tilePattern string
	terrainPattern string
}


func ConstructOfflineMapParameters(cliArgs CommandLineArguments) OfflineMapParametersStruct{
	return OfflineMapParametersStruct{
		rootFolder: cliArgs.MapsFolder,
		layerJsonPath: filepath.Join(cliArgs.MapsFolder,"terrain/layer.json"),
		tilePattern: "tiles/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}",
		terrainPattern: "terrain/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.terrain",
	}
}

func getOptionsFromArgs() (CommandLineArguments,error) {
	if len(os.Args) >= 3 {
		retVal := CommandLineArguments{}
		retVal.Port = os.Args[1]
		retVal.MapsFolder = os.Args[2]
		return retVal,nil
	}
	return CommandLineArguments{}, errors.New("Not enough command line arguments. Usage is Usage ./sertile <port> <map-folder-path>")
}

func tileHandler(w http.ResponseWriter, r *http.Request) {
	options,err := getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	vars := mux.Vars(r)
	z, _ := strconv.Atoi(vars["z"])
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(strings.TrimSuffix(vars["y"], ".png"))

	offlineMapParameters := ConstructOfflineMapParameters(options)

	tilePath := filepath.Join(offlineMapParameters.rootFolder, "/tiles/", fmt.Sprintf("%d", z), fmt.Sprintf("%d", x), fmt.Sprintf("%d.png", y))
	_, osErr := os.Stat(tilePath)
	if osErr != nil{
		tilePath = filepath.Join(offlineMapParameters.rootFolder, "/tiles/", "default.png")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Disposition", "attachment; filename=" + tilePath)
	http.ServeFile(w,r,tilePath)
}

func terrainHandler(w http.ResponseWriter, r *http.Request) {
	options,err := getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	vars := mux.Vars(r)
	z, _ := strconv.Atoi(vars["z"])
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	offlineMapParameters := ConstructOfflineMapParameters(options)

	terrainPath := filepath.Join(offlineMapParameters.rootFolder, "/terrain/", fmt.Sprintf("%d", z), fmt.Sprintf("%d", x), fmt.Sprintf("%d.terrain", y))
	_, osErr := os.Stat(terrainPath)
	if osErr != nil{
		// Override terrainPath to default.
		terrainPath = filepath.Join(offlineMapParameters.rootFolder, "/terrain/", "default.terrain")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Disposition", "attachment; filename=" + terrainPath)
	http.ServeFile(w, r, terrainPath)
}


func layerJsonHandler(w http.ResponseWriter, r *http.Request){

	options,err := getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	offlineMapParameters := ConstructOfflineMapParameters(options)
	http.ServeFile(w, r, offlineMapParameters.layerJsonPath)
}


func main() {

	options,err := getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	offlineMapParameters := ConstructOfflineMapParameters(options)
	r.HandleFunc("/"+offlineMapParameters.tilePattern, tileHandler)
	r.HandleFunc("/"+offlineMapParameters.terrainPattern, terrainHandler)
	r.HandleFunc("/terrain/layer.json", layerJsonHandler)
	http.Handle("/", r)

	port := options.Port
	fmt.Printf("Listening on port %s...\n", port)
	httpErr := http.ListenAndServe(":"+port, nil)
	if httpErr != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
