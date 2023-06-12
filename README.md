# 🐜 Debug Container

Nothing special here at all. It's a small (~10.5MB) container that serves up a Hello World response if your deployment is successful. It's really useful when you're trying to glue together a lot of services and you really just need a stupid container to varifies that things are setup correctly.

The Docker Hub container has been compiled for both `arm64` and `amd64` architectures, so just pull the `latest` tag, and the client will automatically figure out which it needs.

`docker pull kacy/debug-container:latest`

Both `PORT` and `TEST` are environment variables you can use in the 200 response.

A sample response:

```
Success
 port: 8081
 test: default
 random-string: 279b5edf90623797bf03
```
