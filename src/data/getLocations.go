package data

import "log"

/*
CREATE TABLE comida(
    id          INTEGER PRIMARY KEY,
    costo       INTEGER
    deDondeEs TEXT,
    nombre      TEXT,
    image_url   TEXT

);
*/
func GetLocations(locationsChan chan []string) error {
	q := `SELECT DISTINCT deDondeEs FROM comida`
	db := getConnection()
	rows, err := db.Query(q)
	if err != nil {
		log.Println("error originado en lalinea 18  en el paquete data fila getLocations.go")
		log.Println(err)
		return err
	}
	locationsArr := []string{}
	for rows.Next() {
		location := ""
		if err = rows.Scan(&location); err != nil {
			log.Println("error originado en la linea 25 del paquete data fila getLocations.go")
			log.Println(err)
			return err
		}
		locationsArr = append(locationsArr, location)
	}
	locationsChan <- locationsArr
	return nil

}
