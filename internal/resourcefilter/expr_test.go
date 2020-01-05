package resourcefilter

import (
	"testing"

	"github.com/google/cel-go/parser"
	"github.com/stretchr/testify/require"
)

func TestExpr(t *testing.T) {
	e := Or(
		In(Select(Ident("shipment"), "origin_site_id"), Strings("1", "2", "3")),
		In(Select(Ident("shipment"), "destination_site_id"), Strings("1", "2", "3")),
	)
	str, err := parser.Unparse(e, nil)
	require.NoError(t, err)
	require.Equal(t, `shipment.origin_site_id in ["1","2","3"] || shipment.destination_site_id in ["1","2","3"]`, str)
}
