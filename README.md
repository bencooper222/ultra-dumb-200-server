# ultra-dumb-200-server
It is a server that starts up and responds with HTTP 200 on the port and (optionally) path you designate. It does nothing else. It will never do anything else.


# Usage
```bash
export PORT=8080 # must be valid number and in port range
export URL_PATH=/hello # must start with /, defaults to / if unset
`./<binary>`
```