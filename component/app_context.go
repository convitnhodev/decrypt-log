package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	GetKeyText() string
	GetCommonIV() string
}

type appCtx struct {
	db       *gorm.DB
	keyText  string
	commonIV string
}

func NewAppContext(database *gorm.DB, keyText string, commonIV string) *appCtx {
	return &appCtx{database, keyText, commonIV}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) GetKeyText() string {
	return ctx.keyText
}

func (ctx *appCtx) GetCommonIV() string {
	return ctx.commonIV
}
