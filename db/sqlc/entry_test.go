package db

import (
	"context"
	"github.com/TheAlchemistKE/BankBackend/util"
	"github.com/stretchr/testify/require"
	"testing"
)


func createRandomEntry(t *testing.T) (entry Entry) {
	arg := CreateAccountParams{
		Owner: util.RandomOwnerName(),
		Balance: util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	newEntry := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomAmount(),
	}
	entry, err1 := testQueries.CreateEntry(context.Background(), newEntry)
	require.NoError(t, err1)
	require.Equal(t, newEntry.AccountID, entry.AccountID)
	require.Equal(t, newEntry.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return
}

func TestCreateTestEntry(t *testing.T) {
	createRandomEntry(t)
}


func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry1.AccountID, entry.AccountID)
	require.Equal(t, entry1.Amount, entry.Amount)

	require.NotZero(t, entry.ID)

}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		AccountID: 37,
		Limit:     1,
		Offset:    0,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 1)

	for _, entry := range entries {
		require.NotEmpty(t, entry)

	}

}