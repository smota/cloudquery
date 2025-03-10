# Table: aws_resiliencehub_recommendation_templates

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html

The composite primary key for this table is (**account_id**, **region**, **name**).

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|assessment_arn|String|
|format|String|
|name (PK)|String|
|recommendation_template_arn|String|
|recommendation_types|StringArray|
|status|String|
|app_arn|String|
|end_time|Timestamp|
|message|String|
|needs_replacements|Bool|
|recommendation_ids|StringArray|
|start_time|Timestamp|
|tags|JSON|
|templates_location|JSON|