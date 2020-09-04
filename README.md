## REST USING GOLANG

Just a learning project for golang, Implemented CRUD operations with golang as its server and MySQL as database.

### API

- `/home` - Home page
- `/song/all` - Get all entry
- `/song/create` - Create new entry
- `/song/{song}` - Get a particular entry
- `/song/update/{song}` - Update a particular entry
- `/song/delete/{song}` - Delete a particular entry

### Difficulties / Improvements :

- Setting up the file structure was pain in the ass. [ Better approach is to use go modules ]
- Need to learn about struct method and interfaces
- Better Approach for handling database error
- Learn about migration using golang
