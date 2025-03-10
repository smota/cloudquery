# Table: aws_apigateway_rest_api_gateway_responses

https://docs.aws.amazon.com/apigateway/latest/api/API_GatewayResponse.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region|String|
|rest_api_arn|String|
|arn (PK)|String|
|default_response|Bool|
|response_parameters|JSON|
|response_templates|JSON|
|response_type|String|
|status_code|String|