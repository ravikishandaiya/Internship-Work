## Performing CRUD on ZOOM with GO using API
Code is self explaintary, the following are main key point.
### Prerequisite
GO environment (to run go source file)
OS independent.

### Understanding Program
1. ```ListUsers``` and ```User``` are struct type data structure to map data from site.
2. ```handleRequest()``` function connects to the site and does autherization using token.
3. ```ListUser``` function is main function, it calls ```handleRequest()``` function and maps returned data to the above mentioned struct. 

And ready to use data.

Happy Learning.
