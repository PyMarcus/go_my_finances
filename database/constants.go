package database

import "fmt"

const USER string = "marcusv"
const PASSWORD string = "123marcus"

var URL string = fmt.Sprintf(
	"mongodb+srv://%s:%s@cluster0.byzxgxu.mongodb.net", USER, PASSWORD,
)
