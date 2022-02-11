package file_test

import (
	"encoding/json"
	"testing"

	"github.com/brockweaver/hxu/pkg/file"
	"github.com/stretchr/testify/assert"
)

const CONTENT_LINE = "0008803351                                                  0008804266                                                  Primary   Checking                                6490185                                           Checking                                          Open                                              0000000000360862016-07-29T19:46:41.516-05:00                                               000000000000000                                                                                                                                      000000000000000Y0008803281000.025000000000008198543000000000036796000000000050000UNL                                  "

func TestUnit_IsNotTabDelimted(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	assert.False(t, mdl.IsTabDelimitedFile())
}

func TestUnit_HeaderFileCount(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	assert.Equal(t, 5, mdl.HeaderFieldCount())
}

func TestUnit_ParseNative_NoStrict_NoPreserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, false, false)
	assert.Nil(t, err)
	assert.Equal(t, "8803351", mdl.CustomerId)
}

func TestUnit_ParseNative_NoStrict_Preserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, false, true)
	assert.Nil(t, err)
	assert.Equal(t, "0008803351", mdl.CustomerId)
}

func TestUnit_ParseNative_Strict_NoPreserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "8803351", mdl.CustomerId)
}

func TestUnit_ParseNative_Strict_Preserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, true)
	assert.Nil(t, err)
	assert.Equal(t, "0008803351", mdl.CustomerId)
}

func TestUnit_GetJsonValues_NoPreserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, true)
	assert.Nil(t, err)
	assert.Equal(t, "0008803351", mdl.CustomerId)

	data := mdl.GetJsonStruct(false)
	bytes, err2 := json.Marshal(data)
	assert.Nil(t, err2)
	s := string(bytes)
	assert.Contains(t, s, `"customerId":8803351,`)
}

func TestUnit_GetJsonValues_Preserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, true)
	assert.Nil(t, err)
	assert.Equal(t, "0008803351", mdl.CustomerId)

	data := mdl.GetJsonStruct(true)
	bytes, err2 := json.Marshal(data)
	assert.Nil(t, err2)
	s := string(bytes)
	assert.Contains(t, s, `"customerId":"0008803351",`)
}

func TestUnit_GetValues_NoPreserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, false)
	assert.Nil(t, err)
	assert.Equal(t, "8803351", mdl.CustomerId)
	arr := mdl.GetValues()
	assert.LessOrEqual(t, 2, len(arr))
	assert.Equal(t, "8803351", arr[0])
}

func TestUnit_GetValues_Preserve(t *testing.T) {
	mdl := &file.AccountBalanceFileModel{}
	err := mdl.ParseNative(CONTENT_LINE, true, true)
	assert.Nil(t, err)
	assert.Equal(t, "0008803351", mdl.CustomerId)
	arr := mdl.GetValues()
	assert.LessOrEqual(t, 2, len(arr))
	assert.Equal(t, "0008803351", arr[0])
}
