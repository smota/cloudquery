package quicksight

import (
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Folders() *schema.Table {
	return &schema.Table{
		Name:                "aws_quicksight_folders",
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Folder.html",
		Resolver:            fetchQuicksightFolders,
		PreResourceResolver: getFolder,
		Transform:           transformers.TransformWithStruct(&types.Folder{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveTags(),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
