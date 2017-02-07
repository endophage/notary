package remoteks

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewRemoteKS(t *testing.T) {
	_, err := NewRemoteKS("testserver")
	require.NoError(t, err)
}