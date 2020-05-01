package http

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testProviders = map[string]*schema.Provider{
	"http": Provider(),
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestMain(m *testing.M) {
	acctest.UseBinaryDriver("http", Provider)
	resource.TestMain(m)
}
