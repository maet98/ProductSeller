# ProductSeller
This is a fullstack webapp in vuejs with a backend in go (chi) and Dgraph for DB.

It shows information about buyers, the purcharses they have done, give some recommendations and the other buyer that have buyed with the same ip.

### DB

For the db, you must install and configure dgraph (https://dgraph.io/docs/graphql/quick-start/). You must change backend/config/config.json

```json
{
  "IP": "IP_HERE",
  "PORT": "PORT_HERE"
}
```

After that in **DB** you will find a file call query.json. This is the dataaset that you could put in your dgraph db. You can use ratelui(it is usually running in port 8000). Then in mutation paste the content of that file and all the data should load.

### FrontEnd
