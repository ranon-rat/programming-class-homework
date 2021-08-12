package data

import "log"

/*
CREATE TABLE comida(
    id          INTEGER PRIMARY KEY,
    costo       INTEGER,
    deDondeEs TEXT,
    nombre      TEXT,
    image_url   TEXT

);
*/
func GetFoods(locationFood string, foodChan chan []Food) error {
	sql := `SELECT * FROM comida WHERE deDondeEs=$1`
	db := getConnection()
	defer db.Close()
	rows, err := db.Query(sql, locationFood)
	if err != nil {
		log.Println("error en la linea 19 fila data/getFood.go")
		return err
	}
	var foodArr []Food = []Food{}
	for rows.Next() {
		var foodAttr Food
		if err := rows.Scan(&foodAttr.ID, &foodAttr.Cost, &foodAttr.WhereItIs, &foodAttr.Name, &foodAttr.ImageURL); err != nil {
			log.Println("error en la linea 26 fila data/getFood.go")
			return err
		}
		foodArr = append(foodArr, foodAttr)

	}
	foodChan <- foodArr
	return nil
}
