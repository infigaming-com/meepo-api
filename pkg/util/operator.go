package util

import (
	"fmt"
)

const (
	// Operator type: operator, company, retailer, system
	OperatorTypeOperator = "operator"
	OperatorTypeCompany  = "company"
	OperatorTypeRetailer = "retailer"
	OperatorTypeSystem   = "system"

	//Operator mode: individual, co-opearation, internal-co-opearation
	OperatorModeIndividual           = "individual"
	OperatorModeCoOpearation         = "co-opearation"
	OperatorModeInternalCoOpearation = "internal-co-opearation"

	// Operator status: pending, live, suspended, request_to_close, closed
	OperatorStatusPending        = "pending"
	OperatorStatusLive           = "live"
	OperatorStatusSuspended      = "suspended"
	OperatorStatusRequestToClose = "request_to_close"
	OperatorStatusClosed         = "closed"
	OperatorStatusMaintain       = "maintain"

	// Operator Admin Role
	OperatorAdminRoleId         = 4
	CompanyOperatorAdminRoleId  = 3
	RetailerOperatorAdminRoleId = 2
	SystemOperatorAdminRoleId   = 1

	// Operator Player Role
	PlayerRoleId = 0

	// Affiliate Role
	AffiliateRoleId = -1

	// Operator Deduction Order
	OperatorDeductionOrderCashFirst  = "cash_first"  // Cash first, then bonus
	OperatorDeductionOrderBonusFirst = "bonus_first" // Bonus first, then cash
	OperatorDeductionOrderMixed      = "mixed"       // Mixed, cash and bonus are used by ratio

	// Operator template name
	OperatorTemplateNameMobileOnly    = "mobile-only"
	OperatorTemplateNameMobileDesktop = "mobile-desktop"
)

// BuildOperatorHierarchy builds a complete operator hierarchy path based on operatorId and parentIds
// Returns a slice containing the complete path from the bottom-level operatorId to systemOperatorId
// Always returns a list of length 4, with leading zeros filled based on the hierarchy level
// Hierarchy order: operator-company-retailer-system
func BuildOperatorHierarchy(operatorId int64, parentOperatorIds []int64) ([]int64, string, error) {
	// Check if parentIds length is valid (cannot exceed 3)
	if len(parentOperatorIds) > 3 {
		return nil, "", fmt.Errorf("parentIds length cannot exceed 3, got %d", len(parentOperatorIds))
	}

	// Always return a list of length 4
	result := make([]int64, 4)

	var operatorType string

	// Fill the result based on parentIds length
	switch len(parentOperatorIds) {
	case 0: // System level - operatorId at the end (index 3)
		result[0] = 0
		result[1] = 0
		result[2] = 0
		result[3] = operatorId
		operatorType = OperatorTypeSystem

	case 1: // Retailer level - operatorId at third position (index 2)
		result[0] = 0
		result[1] = 0
		result[2] = operatorId
		result[3] = parentOperatorIds[0]
		operatorType = OperatorTypeRetailer

	case 2: // Company level - operatorId at second position (index 1)
		result[0] = 0
		result[1] = operatorId
		result[2] = parentOperatorIds[0]
		result[3] = parentOperatorIds[1]
		operatorType = OperatorTypeCompany

	case 3: // Operator level - operatorId at first position (index 0)
		result[0] = operatorId
		result[1] = parentOperatorIds[0]
		result[2] = parentOperatorIds[1]
		result[3] = parentOperatorIds[2]
		operatorType = OperatorTypeOperator
	}

	return result, operatorType, nil
}

// GetActualOperatorId finds the actual operatorId from a list by finding the last non-zero ID
// Returns 0 if all elements are zero
func GetActualOperatorId(operatorList []int64) int64 {
	if len(operatorList) == 0 {
		return 0
	}

	// Find the last non-zero ID from the beginning
	for i := range operatorList {
		if operatorList[i] != 0 {
			return operatorList[i]
		}
	}

	// If all are zero, return 0
	return 0
}
