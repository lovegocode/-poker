package ranges
type Data struct{
	Cards Card `json:"cards"`
	Tables Table `json:"table"`
}
type Card struct{
   Hand [] string `json:"hand"`
   Position int   `json:"position"`
  Public []int  `json:"publicCards"`
}
type Table struct {
  Person int `json:"person"`
  Action [] int `json:"action"`
}
