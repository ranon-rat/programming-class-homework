package data

func MakeApi(apiChan chan Api) error {

	api := make(Api)
	locationsChan := make(chan []string)
	// obtiene las localizaciones
	go GetLocations(locationsChan)

	locations := <-locationsChan

	// obtiene la comida de cada localizacion
	for _, location := range locations {

		foodChan := make(chan []Food)
		go GetFoods(location, foodChan)

		api[location] = <-foodChan

	}
	apiChan <- api
	//envia el json
	return nil
}
