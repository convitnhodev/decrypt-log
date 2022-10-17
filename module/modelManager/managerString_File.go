package modelManager

import "os"

type ManagerCrud struct {
	DataString string `form:"data" json:"data" bidding:"required"`
	DataFile   os.File
}
