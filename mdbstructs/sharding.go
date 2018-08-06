package mdbstructs

import "github.com/globalsign/mgo/bson"

// Shard reflects a document in the config server 'config.shards'
// collection (or the 'shards' array of the 'listShards' server
// command).
//
// https://docs.mongodb.com/manual/reference/config-database/#config.shards
//
type Shard struct {
	Id    string `bson:"_id"`
	Host  string `bson:"host"`
	State int    `bson:"state"`
}

// ListShards reflects the output of the MongoDB 'listShards' command.
//
// https://docs.mongodb.com/manual/reference/command/listShards/
//
type ListShards struct {
	Shards        []*Shard            `bson:"shards"`
	Ok            int                 `bson:"ok"`
	OperationTime bson.MongoTimestamp `bson:"operationTime"`
	ClusterTime   ClusterTime         `bson:"$clusterTime"`
}
