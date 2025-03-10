package bigquery

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/cloudquery/plugins/source/gcp/client"
	pb "google.golang.org/api/bigquery/v2"
)

func Datasets() *schema.Table {
	return &schema.Table{
		Name:                "gcp_bigquery_datasets",
		Description:         `https://cloud.google.com/bigquery/docs/reference/rest/v2/datasets#Dataset`,
		PreResourceResolver: datasetGet,
		Resolver:            fetchDatasets,
		Multiplex:           client.ProjectMultiplexEnabledServices("bigquery.googleapis.com"),
		Transform:           transformers.TransformWithStruct(&pb.Dataset{}, append(client.Options(), transformers.WithPrimaryKeys("Id"))...),
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			Tables(),
		},
	}
}
