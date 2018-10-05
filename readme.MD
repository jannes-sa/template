# TEMPLATE FRAMEWORK
#### dbstructgen to generate structs from postgres DB name tablename
it will created  in
structs/db/{tablename}.go
helper/constant/tablename/{tablename}.go

``` conf/bin/dbstructgen -table="{nametable}"```

#### httproutergen to generate Routing and Controllers
it will created  in
routers/http/{controller_name}.go
routers/componenttest/http/{controller_name}_test.go
controllers/http/{controller_name}.go


```conf/bin/httproutergen -ctrl="{controller_name}"```

#### grpcroutergen to generate Routing and Controllers
it will created  in
routers/grpc/{controller_name}.go
routers/componenttest/grpc/{controller_name}_test.go
controllers/grpc/{controller_name}.go

```conf/bin/grpcroutergen -ctrl="{controller_name}"```