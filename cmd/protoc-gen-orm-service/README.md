# protoc-gen-orm-service

## Generated Codes

### Select

```proto
message Entity {
	bytes id = 1;
	string name = 2;
	Other other = 3 [
		(orm.edge) = {}
	];
}

message EntitySelect {
	optional bool all = 1;
	optional bool name = 2;
	optional OtherSelect other = 3;
}

message EntityGetRequest {
	EntitySelect select = 1;
	oneof key {
		bytes id = 2;
	}
}
```

Any entities that is selected contains `id`.

#### Get without Select

If `EntityGetRequest.select` is not set, all fields and edges are selected, but for edges, only their IDs are selected.
This is effectively equivalent to:
```json
{
	"select": {
		"all": true,
		"other": {}
	}
}
```


#### `.all`

- not set

	It is treated as `false`.

- `false`

	Only `id` is selected.

- `true`

	All fields are selected.
	The following two selects are identical:
	```json
	{
		"all": true
	}
	```
	```json
	{
		"name": true
	}
	```
