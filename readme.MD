# TEMPLATE FRAMEWORK
#### dbstructgen to generate structs from postgres DB name tablename
it will created  in
structs/db/{tablename}.go
helper/constant/tablename/{tablename}.go

``` conf/bin/dbstructgen -table="{nametable}"```

#### routergen to generate structs from postgres DB name tablename
it will created  in
routers/http/{controller_name}.go
controllers/http/{controller_name}.go

```conf/bin/routergen -ctrl="{controller_name}"```