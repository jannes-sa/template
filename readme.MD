# TEMPLATE FRAMEWORK
#### dbstructgen to generate structs from postgres DB name tablename
it will created  in
structs/db/{table_name}.go
helper/constant/tablename/{table_name}.go

``` conf/bin/dbstructgen -table={table_name}```

#### httproutergen to generate Routing and Controllers
it will created  in
routers/http/{controller_name}.go
routers/componenttest/http/{controller_name}_test.go
controllers/http/{controller_name}.go


```conf/bin/httproutergen -ctrl={controller_name}```

#### grpcroutergen to generate Routing and Controllers
it will created  in
routers/grpc/{controller_name}.go
routers/componenttest/grpc/{controller_name}_test.go
controllers/grpc/{controller_name}.go

```conf/bin/grpcroutergen -ctrl={controller_name}```

#### logicgen to generate logic, dblayer and stub
it will created  in
models/db/interfaces/{logic_name}/{logic_name}.go
models/db/pgsql/{logic_name}/{logic_name}.go
models/db/logic/{logic_name}/{logic_name}.go
models/db/stub/{logic_name}/{logic_name}.go

```conf/bin/logicgen -logic={logic_name} -table={table_name}```