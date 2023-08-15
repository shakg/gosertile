package offlineMapsParameters
/*
rootFolder = getOptionsFromArgs().MapsFolder
layerJsonPath = fmt.Sprintf("%d", filepath.Join(rootFolder,"terrain/layer.json"))
tilePattern = "tiles/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}"
terrainPattern = "terrain/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.terrain"
*/
import (
	cla "commandLineArguments"
)

type OfflineMapParametersStruct struct{
	rootFolder string
	layerJsonPath string
	tilePattern string
	terrainPattern string
}


func ConstructOfflineMapParameters(cliArgs cla.CommandLineArguments) OfflineMapParametersStruct{
	return OfflineMapParameters :=  OfflineMapParametersStruct{
		rootFolder: cliArgs.MapsFolder,
		layerJsonPath: filepath.Join(cliArgs.MapsFolder,"terrain/layer.json"),
		tilePattern: "tiles/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}",
		terrainPattern: "terrain/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.terrain"
	}
}