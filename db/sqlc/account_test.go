package db

import (
	"context"
	"database/sql"
	"github.com/TheAlchemistKE/BankBackend/util"
	"github.com/stretchr/testify/require"
	"testing"
)



func createRandomAccount(t *testing.T) (account Account){
	arg := CreateAccountParams{
		Owner: util.RandomOwnerName(),
		Balance: util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return

}

func TestCreateRandomAccount(t *testing.T) {
	createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	testAcc, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, testAcc)

	require.Equal(t, account1.Owner, testAcc.Owner)
	require.Equal(t, account1.Balance, testAcc.Balance)
	require.Equal(t, account1.Currency, testAcc.Currency)

	require.NotZero(t, testAcc.ID)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomAmount(),
	}
	account1, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account1)

	require.Equal(t, account.Owner, account1.Owner)
	require.Equal(t, account.Currency, account1.Currency)

	require.NotZero(t, account1.ID)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account1, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account1)
}
