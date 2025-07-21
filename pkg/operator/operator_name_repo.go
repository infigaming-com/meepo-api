package operator

import (
	"context"
	"strconv"
	"sync"

	"github.com/infigaming-com/meepo-api/pkg/util"
	user "github.com/infigaming-com/meepo-api/user/service/v1"
	"github.com/samber/lo"
)

type OperatorNameRepo interface {
	SetOperatorIds(ctx context.Context, operatorIds []int64) error
	GetOperatorName(operatorId int64, isSystemOperator bool) string
}

type operatorNameRepo struct {
	operatorNameMap     map[int64]string
	operatorNameMapLock sync.RWMutex
	userClient          user.UserClient
}

func NewOperatorNameRepo(userClient user.UserClient) OperatorNameRepo {
	return &operatorNameRepo{
		operatorNameMap: make(map[int64]string),
		userClient:      userClient,
	}
}

func (o *operatorNameRepo) SetOperatorIds(ctx context.Context, operatorIds []int64) error {
	o.operatorNameMapLock.Lock()
	defer o.operatorNameMapLock.Unlock()

	resp, err := o.userClient.GetOperatorsByIds(
		ctx,
		&user.GetOperatorsByIdsRequest{
			OperatorIds: operatorIds,
		},
	)
	if err != nil {
		return err
	}

	o.operatorNameMap = lo.SliceToMap(
		resp.Operators,
		func(o *user.GetOperatorsByIdsResponse_Operator) (int64, string) {
			return o.OperatorId, o.Name
		},
	)

	return nil
}

func (o *operatorNameRepo) GetOperatorName(operatorId int64, isSystemOperator bool) string {
	o.operatorNameMapLock.RLock()
	defer o.operatorNameMapLock.RUnlock()

	if !isSystemOperator {
		if operatorId == 0 {
			return ""
		}
		name, ok := o.operatorNameMap[operatorId]
		if ok {
			return name
		}
		return strconv.FormatInt(operatorId, 10)
	}

	// system operator
	name, ok := o.operatorNameMap[operatorId]
	if ok {
		return name
	}
	return util.OperatorTypeSystem
}
