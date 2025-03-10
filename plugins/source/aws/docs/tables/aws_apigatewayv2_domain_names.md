# Table: aws_apigatewayv2_domain_names

https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_apigatewayv2_domain_names:
  - [aws_apigatewayv2_domain_name_rest_api_mappings](aws_apigatewayv2_domain_name_rest_api_mappings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|arn (PK)|String|
|domain_name|String|
|api_mapping_selection_expression|String|
|domain_name_configurations|JSON|
|mutual_tls_authentication|JSON|
|tags|JSON|