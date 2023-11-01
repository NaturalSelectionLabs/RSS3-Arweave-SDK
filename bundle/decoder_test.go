package bundle_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/naturalselectionlabs/arweave-go/bundle"
	"github.com/stretchr/testify/require"
)

func TestDecoder(t *testing.T) {
	t.Parallel()

	type arguments struct {
		transactionID string
	}

	testcases := []struct {
		name      string
		arguments arguments
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "rmJZGbi9_UY2yvQPnnwaEGMoLqIVmD5-XJi_aH6FvR0",
			arguments: arguments{
				transactionID: "rmJZGbi9_UY2yvQPnnwaEGMoLqIVmD5-XJi_aH6FvR0",
			},
			wantError: require.NoError,
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			response, err := http.Get(fmt.Sprintf("https://arweave.net/%s", testcase.arguments.transactionID))
			require.NoError(t, err)

			defer func() {
				_ = response.Body.Close()
			}()

			decoder := bundle.NewDecoder(response.Body)

			dataInfos, err := decoder.DecodeDataInfos()
			testcase.wantError(t, err)

			t.Logf("Data Infos: %d", len(dataInfos))

			for _, dataInfo := range dataInfos {
				t.Logf("%+v", dataInfo)
			}

			t.Logf("Data Items")

			for decoder.Next() {
				dataItem, err := decoder.DecodeDataItem()
				testcase.wantError(t, err)

				require.NoError(t, dataItem.Reader.Close())

				t.Logf("%+v", dataItem)
			}
		})
	}
}
