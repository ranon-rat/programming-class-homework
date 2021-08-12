/**
 * type Food struct {
	ID        int    `json:"food-id"`
	WhereItIs string `json:"where-it-is"`
	ImageURL  string `json:"image-url"`
	Name      string `json:"name"`
	Cost      int    `json:"cost"`
}
 */
export interface Food {
  "food-id"?: number;
  cost: number;
  "where-it-is": string;
  "image-url": string;
  name: string;
}
