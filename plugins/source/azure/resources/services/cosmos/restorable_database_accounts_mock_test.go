package cosmos

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createRestorableDatabaseAccounts(router *mux.Router) error {
	var item armcosmos.RestorableDatabaseAccountsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/restorableDatabaseAccounts", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestRestorableDatabaseAccounts(t *testing.T) {
	client.MockTestHelper(t, RestorableDatabaseAccounts(), createRestorableDatabaseAccounts)
}
