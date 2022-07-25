// Package go_req
// Time    : 2022/7/24 16:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package go_req

import (
	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func TestExampleGetJson(t *testing.T) {
	jB, err := ExampleGetJson()
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"ID", "Title", "Body", "UserID"}
	table.SetHeader(header)
	table.Append([]string{
		strconv.Itoa(jB.ID), jB.Title, jB.Body, strconv.Itoa(jB.UserID),
	})
	table.Render()

	require.NoError(t, err)
}

func TestExamplePostJson(t *testing.T) {
	jB, err := ExamplePostJson()
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{"ID", "Title", "Body", "UserID"}
	table.SetHeader(header)
	table.Append([]string{
		strconv.Itoa(jB.ID), jB.Title, jB.Body, strconv.Itoa(jB.UserID),
	})
	table.Render()

	require.NoError(t, err)
}
