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
		operatorValues := getOperatorAndValue(v)
		for _, ov := range operatorValues {
			condition := k + " " + ov[0].(string) + " ?"
			log.Printf("cond -> %+v & value -> %+v", condition, ov[1])
			trx = trx.Where(condition, ov[1])
		}
	}
	trx.Find(model)
}

func NewCrudService(db *gorm.DB) *Crud {
	return &Crud{db: db}
}
