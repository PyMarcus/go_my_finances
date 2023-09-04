package database

import "fmt"

const USER string = ""
const PASSWORD string = ""

var URL string = fmt.Sprintf(
	"mongodb+srv://%s:%s@cluster0.byzxgxu.mongodb.net", USER, PASSWORD,
)
