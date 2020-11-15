# MongoDB

## Highlights

- the hierarchy is: database.collection.documents
- database and collection is automatically created when it is first inserted a document
- documents are BSON documents: 
  - Document: bson.D{ {"name", "Andrzej"}, {"age", 33} }
  - Map: bson.M{"poland":"warsaw", "italy":"rome"}
  - Array bson.A{1, 2, 3, 4}

## Play around with mongo shell (first: ./run_all.sh)

```bash
docker exec -it mymongo bash
mongo --username myuser --password mypass
> show users
(empty)

> show databases
admin        0.000GB
coffee-shop  0.000GB # created by app
config       0.000GB 
grocery      0.000GB # created by app
local        0.000GB

> show roles
{
	"role" : "dbAdmin",
	"db" : "test",
	"isBuiltin" : true,
	"roles" : [ ],
	"inheritedRoles" : [ ]
}
(... 5 more)

> use grocery # change current DB to 'grocery'
switched to db grocery

> show collections
items

> db.items.find()
{ "_id" : ObjectId("5fb0105cca62310485375a01"), "id" : ObjectId("5fb0105cca62310485375a00"), "title" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb0105cca62310485375a03"), "id" : ObjectId("5fb0105cca62310485375a02"), "title" : "Milk", "price" : NumberLong(3) }
{ "_id" : ObjectId("5fb0105cca62310485375a05"), "id" : ObjectId("5fb0105cca62310485375a04"), "title" : "Guacamole", "price" : NumberLong(23) }

> db.stats()
{
	"db" : "grocery",
	"collections" : 1,
	"views" : 0,
	"objects" : 3,
	"avgObjSize" : 71,
	"dataSize" : 213,
	"storageSize" : 20480,
	"indexes" : 1,
	"indexSize" : 20480,
	"totalSize" : 40960,
	"scaleFactor" : 1,
	"fsUsedSize" : 118452633600,
	"fsTotalSize" : 754873942016,
	"ok" : 1
}

```