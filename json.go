package testsuite

import (
	"encoding/json"
	"testing"

	"github.com/itchyny/gojq"
	"github.com/stretchr/testify/require"
)

// JQQuery performs a test on a given json output. A jq-like query is performed
// on the given jsonData and the result is compared with the expected output.
// If the output doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func JQQuery(t *testing.T, jsonData []byte, jqQuery string, expected interface{}, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse(jqQuery)
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, expected, v)
	require.Equal(t, expected, v, msgAndArgs...)
}
