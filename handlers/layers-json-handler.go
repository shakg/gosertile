package layerJsonHandler
import (
	"net/http"
	omp "./../shared/offlineMapsParameters"
	ofa "./../utils/optionsFromArguments.go"
)

func layerJsonHandler(w http.ResponseWriter, r *http.Request){

	options,err := ofa.getOptionsFromArgs()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	offlineMapParameters := omp.ConstructOfflineMapParameters(options)
	http.ServeFile(w, r, offlineMapParameters.layerJsonPath)
}
