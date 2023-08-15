package tileHandler
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
func tileHandler(w http.ResponseWriter, r *http.Request) {
	options,err := ofa.getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	vars := mux.Vars(r)
	z, _ := strconv.Atoi(vars["z"])
	x, _ := strconv.Atoi(vars["x"])
	y, _ := strconv.Atoi(strings.TrimSuffix(vars["y"], ".png"))

	offlineMapParameters := omp.ConstructOfflineMapParameters(options)

	tilePath := filepath.Join(offlineMapParameters.rootFolder, "/tiles/", fmt.Sprintf("%d", z), fmt.Sprintf("%d", x), fmt.Sprintf("%d.png", y))
	_, err := os.Stat(tilePath)
	if err != nil{
		tilePath = filepath.Join(offlineMapParameters.rootFolder, "/tiles/", "default.png")
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Disposition", "attachment; filename=" + tilePath)
	http.ServeFile(w,r,tilePath)
}