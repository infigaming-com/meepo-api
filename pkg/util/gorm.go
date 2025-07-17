package util

import (
	"github.com/infigaming-com/meepo-api/common"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func BuildOperatorContextQuery(query *gorm.DB, operatorContext *common.OperatorContext) *gorm.DB {
	if operatorContext == nil {
		return query
	}
	if operatorContext.OperatorId != 0 {
		query = query.Where("operator_id = ?", operatorContext.OperatorId)
	}
	if operatorContext.CompanyOperatorId != 0 {
		query = query.Where("company_operator_id = ?", operatorContext.CompanyOperatorId)
	}
	if operatorContext.RetailerOperatorId != 0 {
		query = query.Where("retailer_operator_id = ?", operatorContext.RetailerOperatorId)
	}
	if operatorContext.SystemOperatorId != 0 {
		query = query.Where("system_operator_id = ?", operatorContext.SystemOperatorId)
	}
	return query
}

func BuildOperatorContextFiltersQuery(query *gorm.DB, operatorContextFilters *common.OperatorContextFilters) *gorm.DB {
	if operatorContextFilters == nil {
		return query
	}
	if len(operatorContextFilters.OperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(operatorContextFilters.OperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.OperatorId, context.CompanyOperatorId, context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(operator_id, company_operator_id, retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	} else if len(operatorContextFilters.CompanyOperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(operatorContextFilters.CompanyOperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.CompanyOperatorId, context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(company_operator_id, retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	} else if len(operatorContextFilters.RetailerOperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(operatorContextFilters.RetailerOperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	}
	return query
}
