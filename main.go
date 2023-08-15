package main

import (
	"fmt"
	"net/http"
	"os"
	omp "./shared/offlineMapsParameters"
	tih "./handlers/tile-handler.go"
	teh "./handlers/terrain-handler.go"
	ljh "./handlers/layers-json-handler.go"
	ofa "./utils/optionsFromArguments.go"

	"github.com/gorilla/mux"

)

func main() {

	options,err := ofa.getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	r := mux.NewRouter()

	offlineMapParameters := omp.ConstructOfflineMapParameters(options)
	r.HandleFunc("/"+offlineMapParameters.tilePattern, tih.tileHandler)
	r.HandleFunc("/"+offlineMapParameters.terrainPattern, teh.terrainHandler)
	r.HandleFunc("/terrain/layer.json", ljh.layerJsonHandler)
	http.Handle("/", r)

	port := options.Port
	fmt.Printf("Listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
