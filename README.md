# Blog-http

## Available actions

- Register
- Login
- Forgot password
- Admin posts a blog post
- Admin updates a blog post
- Admin deletes a blog post
- Like posts
- Unlike posts
- Comment on blog posts
- View blog posts
- Delete comments
- Update comments
- Like comments
- Unlike comments

## Actions according to modules

- auth
- post
- comment

## To run the server, use any of

- gin -i -a 3000 -p 3001 run main.go (-a is the app port which is configurable from the .env file, -p is the proxy port for the gin package, and -i is to run the server immediately)
- go run main.go
