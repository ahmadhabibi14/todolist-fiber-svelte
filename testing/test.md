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

### Stats
```sh
curl -X GET -i localhost:3000/api/stats
```

### Get Todos
```sh
curl -X GET -i localhost:3000/api/todo/list
```