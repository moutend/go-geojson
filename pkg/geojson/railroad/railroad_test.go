package railroad

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalFeatureCollection(t *testing.T) {
	t.Parallel()

	data, err := ioutil.ReadFile(filepath.Join("testdata", "input.geojson"))

	if err != nil {
		return
	}

	var fc FeatureCollection

	require.NoError(t, json.Unmarshal(data, &fc))
	require.NotEqual(t, 0, len(fc.Features))
}
