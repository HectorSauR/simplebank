package db

import (
	"context"
	"testing"
	"time"

	"github.com/HectorSauR/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accountId int64, amount int64) Entry {

	if amount == 0 {
		amount = util.RandomMoney()
	}

	args := CreateEntryParams{
		AccountID: accountId,
		Amount:    amount,
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, accountId, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account1 := createRandomAccount(t)

	createRandomEntry(t, account1.ID, 0)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := createRandomEntry(t, account1.ID, 0)

	args := GetEntryParams{
		ID:        entry1.ID,
		AccountID: entry1.AccountID,
	}

	entry2, err := testQueries.GetEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account1 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, account1.ID, 0)
	}

	args := ListEntryParams{
		AccountID: account1.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.Equal(t, entry.AccountID, account1.ID)
		require.NotEmpty(t, entry)
	}
}
