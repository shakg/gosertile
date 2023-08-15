package terrainHandler
import (
	"net/http"
	"fmt"
	"filepath"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
	omp "./../shared/offlineMapsParameters"
	ofa "./../utils/optionsFromArguments.go"

)

func terrainHandler(w http.ResponseWriter, r *http.Request) {
	options,err := ofa.getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	vars := mux.Vars(r)
	z, _ := strconv.Atoi(vars["z"])
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(vars["y"])

	offlineMapParameters := omp.ConstructOfflineMapParameters(options)

	terrainPath := filepath.Join(offlineMapParameters.rootFolder, "/terrain/", fmt.Sprintf("%d", z), fmt.Sprintf("%d", x), fmt.Sprintf("%d.terrain", y))
	_, err := os.Stat(terrainPath)
	if err != nil{
		// Override terrainPath to default.
		terrainPath = filepath.Join(offlineMapParameters.rootFolder, "/terrain/", "default.terrain")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Disposition", "attachment; filename=" + terrainPath)
	http.ServeFile(w, r, terrainPath)
}
