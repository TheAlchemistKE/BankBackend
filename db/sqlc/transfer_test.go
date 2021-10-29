package db

import (
	"context"
	"github.com/TheAlchemistKE/BankBackend/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomTransfer(t *testing.T) (transfer Transfer){
	arg := CreateTransferParams{
		FromAccountID: 2,
		ToAccountID:   3,
		Amount:        util.RandomAmount(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	arg := createRandomTransfer(t)
	transfer, err := testQueries.GetTransfer(context.Background(), arg.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		FromAccountID: 2,
		ToAccountID:   3,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

	}
}
