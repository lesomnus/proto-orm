# proto-orm

ORM support for [*protobuf*](https://protobuf.dev/).

Designed with inspiration from [*Ent*](https://entgo.io/), it supports [*gRPC*](https://grpc.io/) auto-generation based on *Ent*. I look forward to supporting more backends beyond *Ent*.

## Peek

```protobuf
syntax = "proto3";

import "orm.proto";
import "google/protobuf/timestamp.proto";

message User {
	bytes id = 1 [
		(orm.field) = {
			type: TYPE_UUID
			key: true
			default: ""
		}
	];
	repeated Pet pets = 2;

	message Profile {
		message Name {
			string first = 1;
			string last = 2;
		}
		Name name = 1;
		uint32 age = 2;
		string email = 3;
	}
	Profile profile = 3 [
		(orm.field) = {
			type: TYPE_JSON
		}
	];

	google.protobuf.Timestamp date_created = 15 [
		(orm.field) = {
			immutable: true
			default: ""
		}
	];
}

message Pet {
	bytes id = 1 [
		(orm.field) = {
			type: TYPE_UUID
			key: true
			default: ""
		}
	];
	User owner = 2 [
		(orm.edge) = {
			from: 2 // pets
		}
	];
}

```

## Limitations

- Does not support composite key.
