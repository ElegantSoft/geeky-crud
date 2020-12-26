package crud

import (
	"gorm.io/gorm"
	"log"
)

type Crud struct {
	db *gorm.DB
}

func (c *Crud) Find(query findQuery, model interface{}) {
	trx := c.db
	for k, v := range query.q {
		operator, value := getOperatorAndValue(v)
		condition := k + " " + operator + " ?"
		log.Printf("cond -> %+v & value -> %+v", condition, value)
		trx = trx.Where(condition, value)
	}
	trx.Find(model)
}

func NewCrudService(db *gorm.DB) *Crud {
	return &Crud{db: db}
}
