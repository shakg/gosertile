package optionsFromArguments
import (
	"errors"

	cla "../shared/commandLineArguments"
)
func getOptionsFromArgs() (cla.CommandLineArguments,error) {
	if len(os.Args) >= 3 {
		retVal := CommandLineArguments{}
		retVal.Port = os.Args[1]
		retVal.MapsFolder = os.Args[2]
		return retVal,nil
	}
	return nil, errors.New("Not enough command line arguments. Usage is Usage ./sertile <port> <map-folder-path>")
}