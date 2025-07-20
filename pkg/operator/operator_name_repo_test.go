package operator

import (
	"testing"
	"time"

	user "github.com/infigaming-com/meepo-api/user/service/v1"

	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewUserClient() user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9101"),
		grpc.WithTimeout(time.Second*300),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("failed to dial user service: %v", err)
	}
	return user.NewUserClient(conn)
}

func TestOperatorNameRepoWithInitialOperatorIds(t *testing.T) {
	userClient := NewUserClient()

	operatorNameRepo, err := NewOperatorNameRepo(userClient, []int64{0, 1234567891, 133510782607425537})
	require.NoError(t, err)

	assert.Equal(t, operatorNameRepo.GetOperatorName(1234567891, false), "Speedix")
	assert.Equal(t, operatorNameRepo.GetOperatorName(0, true), "system")
}

func TestOperatorNameRepoWithSetOperatorIds(t *testing.T) {
	userClient := NewUserClient()

	operatorNameRepo, err := NewOperatorNameRepo(userClient, []int64{})
	require.NoError(t, err)

	err = operatorNameRepo.SetOperatorIds([]int64{0, 1234567891, 133510782607425537})
	require.NoError(t, err)

	assert.Equal(t, operatorNameRepo.GetOperatorName(1234567891, false), "Speedix")
	assert.Equal(t, operatorNameRepo.GetOperatorName(0, true), "system")
}
