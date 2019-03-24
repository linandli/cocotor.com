package torchlib

import (
	"fmt"
	"cocotor.com/utils/torchlib"
)

func GetPlants() map[string]string {
//http.get(['org', 'plants'])
	res = torchlib.Get("org/plants", nil, 60)
	fmt.Println("r----------->", res)
	return 
}