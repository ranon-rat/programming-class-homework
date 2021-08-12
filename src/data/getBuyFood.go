package data

/**
CREATE TABLE comida(
    id          INTEGER PRIMARY KEY,
    costo       INTEGER,
    deDondeEs TEXT,
    nombre      TEXT,
    image_url   TEXT

);
*/
func GetBuyFood(id int) Food {
	sql := `SELECT * FROM comida WHERE id=$1`
	db := getConnection()
	defer db.Close()
	var food Food
	db.QueryRow(sql, id).Scan(&food.ID, &food.Cost, &food.WhereItIs, &food.Name, &food.ImageURL)
	return food
}
