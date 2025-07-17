package util

import (
	"github.com/infigaming-com/meepo-api/common"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func BuildOperatorContextQuery(query *gorm.DB, RetailerOperatorContexts []*common.OperatorContext, CompanyOperatorContexts []*common.OperatorContext, OperatorContexts []*common.OperatorContext) *gorm.DB {
	if len(OperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(OperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.OperatorId, context.CompanyOperatorId, context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(operator_id, company_operator_id, retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	} else if len(CompanyOperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(CompanyOperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.CompanyOperatorId, context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(company_operator_id, retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	} else if len(RetailerOperatorContexts) > 0 {
		// build a list of operatorIdsList
		operatorIdsList := lo.Map(RetailerOperatorContexts, func(context *common.OperatorContext, _ int) []int64 {
			return []int64{context.RetailerOperatorId, context.SystemOperatorId}
		})
		query = query.Where("(retailer_operator_id, system_operator_id) IN ?", operatorIdsList)
	}
	return query
}
