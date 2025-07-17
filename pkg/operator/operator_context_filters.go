package operator

import (
	"github.com/infigaming-com/meepo-api/common"
	"github.com/samber/lo"
)

type OperatorContextFilters struct {
	OperatorContexts         []OperatorContext
	CompanyOperatorContexts  []OperatorContext
	RetailerOperatorContexts []OperatorContext
}

func NewOperatorContextFilters(operatorContextFilters *common.OperatorContextFilters) *OperatorContextFilters {
	return &OperatorContextFilters{
		OperatorContexts: lo.Map(operatorContextFilters.OperatorContexts, func(context *common.OperatorContext, _ int) OperatorContext {
			return *NewOperatorContext(context)
		}),
		CompanyOperatorContexts: lo.Map(operatorContextFilters.CompanyOperatorContexts, func(context *common.OperatorContext, _ int) OperatorContext {
			return *NewOperatorContext(context)
		}),
		RetailerOperatorContexts: lo.Map(operatorContextFilters.RetailerOperatorContexts, func(context *common.OperatorContext, _ int) OperatorContext {
			return *NewOperatorContext(context)
		}),
	}
}
