# Table: heroku_releases

https://devcenter.heroku.com/articles/platform-api-reference#release

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|addon_plan_names|StringArray|
|app|JSON|
|created_at|Timestamp|
|current|Bool|
|description|String|
|output_stream_url|String|
|slug|JSON|
|status|String|
|updated_at|Timestamp|
|user|JSON|
|version|Int|