package identity

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracle/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/oracle/oci-go-sdk/v65/identity"
)

func DynamicGroups() *schema.Table {
	return &schema.Table{
		Name:      "oracle_identity_dynamic_groups",
		Resolver:  fetchDynamicGroups,
		Multiplex: client.TenancyMultiplex,
		Transform: transformers.TransformWithStruct(&identity.DynamicGroup{},
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
