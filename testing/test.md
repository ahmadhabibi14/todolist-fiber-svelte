### Login
```sh
curl -X POST -d 'username=habi&password=habi124' -i localhost:3000/api/login 
```

### Logout
```sh
curl -X DELETE
   >  -H 'X-Session-ID: 7ac8b666-915a-4143-846d-bfaaa5226376'
   >  -i
   >  --url localhost:3000/api/logout
```

### Add Todo
```sh
curl -X POST
   >  -d 'text=Haloohabi'
   >  -H 'X-Session-ID: 7ac8b666-915a-4143-846d-bfaaa5226376'
   >  -i
   >  --url localhost:3000/api/todo/add
```

### Edit Todo
```sh
curl -X PUT
   >  -d 'text=habibidoifiefi'
   >  -H 'X-Session-ID: 2c1579eb-7078-423f-a825-732738908d9c'
   >  -i localhost:3000/api/todo/overwrite
```

### Stats
```sh
curl -X GET -i localhost:3000/api/stats
```

### Get Todos
```sh
curl -X GET -i localhost:3000/api/todo/list
```