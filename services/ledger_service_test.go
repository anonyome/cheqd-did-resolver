package services

import (
	"testing"
	"time"

	didTypes "github.com/cheqd/cheqd-node/x/did/types"
	resource "github.com/cheqd/cheqd-node/x/resource/types"
	"github.com/cheqd/did-resolver/types"
	"github.com/stretchr/testify/require"
)

func TestQueryDIDDoc(t *testing.T) {
	subtests := []struct {
		name                     string
		did                      string
		expectedDidDocWithMetada *didTypes.DidDocWithMetadata
		expectedIsFound          bool
		expectedError            error
	}{
		{
			name:                     "DeadlineExceeded",
			did:                      "fake did",
			expectedDidDocWithMetada: nil,
			expectedIsFound:          false,
			expectedError:            types.NewInvalidDIDError("fake did", types.JSON, nil, false),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			timeout, err := time.ParseDuration("5s")
			require.NoError(t, err)

			ledgerService := NewLedgerService(timeout, false)
			didDocWithMetadata, err := ledgerService.QueryDIDDoc("fake did")
			require.EqualValues(t, subtest.expectedDidDocWithMetada, didDocWithMetadata)
			require.EqualValues(t, subtest.expectedError.Error(), err.Error())
		})
	}
}

func TestQueryResource(t *testing.T) {
	subtests := []struct {
		name             string
		collectionDid    string
		resourceId       string
		expectedResource *resource.Resource
		expectedError    error
	}{
		{
			name:             "DeadlineExceeded",
			collectionDid:    "321",
			resourceId:       "123",
			expectedResource: nil,
			expectedError:    types.NewInvalidDIDError("321", types.JSON, nil, true),
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.name, func(t *testing.T) {
			timeout, err := time.ParseDuration("5s")
			require.NoError(t, err)

			ledgerService := NewLedgerService(timeout, false)
			resource, err := ledgerService.QueryResource(subtest.collectionDid, subtest.resourceId)
			require.EqualValues(t, subtest.expectedResource, resource)
			require.EqualValues(t, subtest.expectedError.Error(), err.Error())
		})
	}
}
