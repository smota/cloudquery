package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func NetworkSources() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_network_sources",
		Resolver:  fetchNetworkSources,
		Multiplex: client.TenancyMultiplex,
		Transform: transformers.TransformWithStruct(&identity.NetworkSourcesSummary{},
			transformers.WithTypeTransformer(client.OracleTypeTransformer)),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
