# MongoDB

Mongo shell tutorial: <https://www.youtube.com/watch?v=-56x56UppqQ>
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

# Notice: dont mind the '_id' value; these examples come from different mongo sessions

> show roles
{
	"role" : "dbAdmin",
	"db" : "test",
	"isBuiltin" : true,
	"roles" : [ ],
	"inheritedRoles" : [ ]
}
(... 5 more)

# change current DB to 'grocery'
> use grocery 
switched to db grocery

# list collections in DB 'grocery'
> show collections
items

# find all documents in collection 'items'
> db.items.find()
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }

# create new document
> db.items.insert ({ name: "Hummus", price: NumberLong(8) })
WriteResult({ "nInserted" : 1 })

> db.items.find()
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }
{ "_id" : ObjectId("5fb1428b477ca6c011f102d1"), "name" : "Hummus", "price" : NumberLong(8) }

# find by exact name
> db.items.find({name:"Hummus"})
{ "_id" : ObjectId("5fb1428b477ca6c011f102d1"), "name" : "Hummus", "price" : NumberLong(8) }

# find by minimu price (gte = greater or equal)
> db.items.find({price: {$gte:4}})
{ "_id" : ObjectId("5fb1513feaf0c581fc652c6f"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb1513feaf0c581fc652c70"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb1428b477ca6c011f102d1"), "name" : "Hummus", "price" : NumberLong(8) }


# find the first 2 documents
> db.items.find().limit(2)
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }

# get document count
> db.items.find().count()
4

# find and sort descending by price
> db.items.find().sort({price:-1})
{ "_id" : ObjectId("5fb1428b477ca6c011f102d1"), "name" : "Hummus", "price" : NumberLong(8) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }

# change price of Hummus to 11
> db.items.update({name: "Hummus"}, {$set: {price: NumberLong(11)}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })

> db.items.find().sort({price:-1})
{ "_id" : ObjectId("5fb1425e477ca6c011f102d0"), "name" : "Hummus", "price" : NumberLong(11) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }

# increment price of Hummus by 2
> db.items.update({name: "Hummus"}, {$inc: {price: 2}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })

> db.items.find().sort({price:-1})
{ "_id" : ObjectId("5fb1425e477ca6c011f102d0"), "name" : "Hummus", "price" : NumberLong(13) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }

# rename document field price -> cost
> db.items.update({name: "Hummus"}, {$rename: {price: "cost"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })

> db.items.find().sort({price:-1})
{ "_id" : ObjectId("5fb14201bc3a7070f22e229f"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e229e"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14201bc3a7070f22e22a0"), "name" : "Pasztet", "price" : NumberLong(1) }
{ "_id" : ObjectId("5fb1425e477ca6c011f102d0"), "name" : "Hummus", "cost" : NumberLong(13) }

# remove document
> db.items.remove({ name: "Hummus"})
WriteResult({ "nRemoved" : 1 })

> db.items.find().sort({price:-1})
{ "_id" : ObjectId("5fb14df947904917007cf0d8"), "name" : "Milk", "price" : NumberLong(7) }
{ "_id" : ObjectId("5fb14df947904917007cf0d7"), "name" : "Bread", "price" : NumberLong(4) }
{ "_id" : ObjectId("5fb14eb327263736c7e3fad9"), "name" : "Pasztet", "price" : NumberLong(1) }


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