package owntest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAssert(t *testing.T) {
	// 断⾔两个值相等，否则就失败,并输出提示⽂本
	assert.Equal(t, 123, 123, "they should be equal")
	// 断⾔不相等
	assert.NotEqual(t, 123, 456, "they should not be equal")
	// 断⾔object是nil值
	assert.Nil(t, nil)
	// 断⾔object⾮nil
	if assert.NotNil(t, "not nil") {
		// 断⾔object.Value和某值相等
		assert.Equal(t, "not nil", "not nil")
	}
}

func TestRequire(t *testing.T) {
	// 断⾔两个值相等，否则就失败,并输出提示⽂本
	require.Equal(t, 123, 123, "they should be equal")

	// 断⾔不相等
	require.NotEqual(t, 123, 456, "they should not be equal")
	// 断⾔object是nil值
	require.Nil(t, nil)
}
