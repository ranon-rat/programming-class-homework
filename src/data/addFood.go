package data

import "log"

func AddFood(addThisFood Food) {
	q := `INSERT INTO comida(
		costo,
		deDondeEs,
		nombre,
		image_url) VALUES($1,$2,$3,$4);`
	db := getConnection()
	stm, err := db.Prepare(q)
	if err != nil {
		log.Println("error en linea 24 add food data")
		log.Println(err)
		return
	}
	if _, err := stm.Exec(addThisFood.Cost, addThisFood.WhereItIs, addThisFood.Name, addThisFood.ImageURL); err != nil {
		log.Println(err)
	}

}
