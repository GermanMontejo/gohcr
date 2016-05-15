# gohcr
CLI-based http request client similar to curl but made from Go.

How to use gohcr (Go HTTP Client Requester):

gohcr -http method- -url- -header- -body-

i.e. gohcr POST https://127.0.0.1:8080/api/users ["Content-type":"application/json"] {"firstname":"german","lastname":"montejo","email":"gemontejo@gmail.com"}

Note, in your url, you can also include your params but you should add an escape character to & so your shell will not interpret it as a & (background) option, i.e.:

gohcr POST https://127.0.0.1:2121/users?id=abc123\&age=25 ["Content-type":"application/json"] {"firstname":"german","lastname":"montejo","email":"gemontejo@gmail.com"}

Headers are working fine, as tested in my machine. Should there be any problems, feel free to open an issue for it.

This project already has a runnable file named: ghcr.
If you want to make changes, then build this project, you can:
Run the script file named: gohcrbuilder.sh
That script builds main.go and renames it to gohcr (for consistency).
