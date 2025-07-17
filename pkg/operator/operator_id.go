package operator

import (
	"github.com/infigaming-com/meepo-api/common"
)

const (
	OperatorTypeOperator = "operator"
	OperatorTypeCompany  = "company"
	OperatorTypeRetailer = "retailer"
	OperatorTypeSystem   = "system"
)

type OperatorId struct {
	SystemOperatorId   int64
	CompanyOperatorId  int64
	RetailerOperatorId int64
	OperatorId         int64
	RealOperatorId     int64
	OperatorType       string
}

func NewOperatorId(operatorId, companyOperatorId, retailerOperatorId, systemOperatorId int64) *OperatorId {
	id := &OperatorId{
		OperatorId:         operatorId,
		CompanyOperatorId:  companyOperatorId,
		RetailerOperatorId: retailerOperatorId,
		SystemOperatorId:   systemOperatorId,
	}
	realOperatorId, operatorType := id.GetRealOperatorIdAndType()
	id.RealOperatorId = realOperatorId
	id.OperatorType = operatorType
	return id
}

func NewOperatorIdWithOperatorContext(operatorContext *common.OperatorContext) *OperatorId {
	return &OperatorId{
		OperatorId:         operatorContext.OperatorId,
		CompanyOperatorId:  operatorContext.CompanyOperatorId,
		RetailerOperatorId: operatorContext.RetailerOperatorId,
		SystemOperatorId:   operatorContext.SystemOperatorId,
		RealOperatorId:     operatorContext.OperatorId,
		OperatorType:       operatorContext.OperatorType,
	}
}

func (id *OperatorId) GetRealOperatorIdAndType() (int64, string) {
	if id.OperatorId != 0 {
		return id.OperatorId, OperatorTypeOperator
	} else if id.CompanyOperatorId != 0 {
		return id.CompanyOperatorId, OperatorTypeCompany
	} else if id.RetailerOperatorId != 0 {
		return id.RetailerOperatorId, OperatorTypeRetailer
	} else {
		return id.SystemOperatorId, OperatorTypeSystem
	}
}

func (id *OperatorId) GetOperatorContext() *common.OperatorContext {
	return &common.OperatorContext{
		OperatorId:         id.OperatorId,
		CompanyOperatorId:  id.CompanyOperatorId,
		RetailerOperatorId: id.RetailerOperatorId,
		SystemOperatorId:   id.SystemOperatorId,
		RealOperatorId:     id.RealOperatorId,
		OperatorType:       id.OperatorType,
	}
}

func (id *OperatorId) IsInOperatorContext(operatorContext *OperatorContext) bool {
	switch id.OperatorType {
	case OperatorTypeOperator:
		return id.OperatorId == operatorContext.OperatorId &&
			id.CompanyOperatorId == operatorContext.CompanyOperatorId &&
			id.RetailerOperatorId == operatorContext.RetailerOperatorId &&
			id.SystemOperatorId == operatorContext.SystemOperatorId
	case OperatorTypeCompany:
		return id.CompanyOperatorId == operatorContext.CompanyOperatorId &&
			id.RetailerOperatorId == operatorContext.RetailerOperatorId &&
			id.SystemOperatorId == operatorContext.SystemOperatorId
	case OperatorTypeRetailer:
		return id.RetailerOperatorId == operatorContext.RetailerOperatorId &&
			id.SystemOperatorId == operatorContext.SystemOperatorId &&
			id.OperatorId == operatorContext.OperatorId
	case OperatorTypeSystem:
		return id.SystemOperatorId == operatorContext.SystemOperatorId
	}
	return false
}
