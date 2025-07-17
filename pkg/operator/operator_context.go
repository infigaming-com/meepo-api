package operator

import "github.com/infigaming-com/meepo-api/common"

type OperatorContext struct {
	OperatorId         int64
	CompanyOperatorId  int64
	RetailerOperatorId int64
	SystemOperatorId   int64
	RealOperatorId     int64
	OperatorType       string
}

func NewOperatorContext(operatorContext *common.OperatorContext) *OperatorContext {
	return &OperatorContext{
		OperatorId:         operatorContext.OperatorId,
		CompanyOperatorId:  operatorContext.CompanyOperatorId,
		RetailerOperatorId: operatorContext.RetailerOperatorId,
		SystemOperatorId:   operatorContext.SystemOperatorId,
		RealOperatorId:     operatorContext.OperatorId,
		OperatorType:       operatorContext.OperatorType,
	}
}

func NewOperatorContextWithIds(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId int64) *OperatorContext {
	operatorIds := NewOperatorId(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId)
	return NewOperatorContext(operatorIds.GetOperatorContext())
}

func (o *OperatorContext) GetOperatorContext() *OperatorContext {
	return &OperatorContext{
		OperatorId:         o.OperatorId,
		CompanyOperatorId:  o.CompanyOperatorId,
		RetailerOperatorId: o.RetailerOperatorId,
		SystemOperatorId:   o.SystemOperatorId,
	}
}
