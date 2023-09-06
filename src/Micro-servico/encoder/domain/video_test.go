package domain_test

import (
	"encoder/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)

}
