package crud

import (
	"gorm.io/gorm"
	"log"
)

type Crud struct {
	db *gorm.DB
}

func (c *Crud) Find(query FindQuery, model interface{}) {
	trx := c.db

	// adding where statements for query
	for k, v := range query.Q {
		operatorValues := getOperatorAndValue(v)
		for _, ov := range operatorValues {
			condition := k + " " + ov[0].(string) + " ?"
			log.Printf("cond -> %+v & value -> %+v", condition, ov[1])
			if ov[0].(string) != "" && ov[1] != nil {
				trx = trx.Where(condition, ov[1])
			}
		}
	}

	// load relations
	for _, v := range query.Joins {
		gormRelation := convertSnakeToGormPascal(v)
		trx.Preload(gormRelation)
	}

	// adding selects
	if len(query.Fields) > 0 {
		trx.Select(query.Fields)
	}
	trx.Find(model)
}

func NewCrudService(db *gorm.DB) *Crud {
	return &Crud{db: db}
}
