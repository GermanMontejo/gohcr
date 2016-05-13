# gohcr
CLI-based http request client similar to curl but made from Go.

How to use gohcr (Go HTTP Client Requester):

gohcr <http method> <url> <header> <body>
i.e. gohcr ghcr POST https://127.0.0.1:8080/api/users ["Content-type":"application/json"] {"firstname":"german","lastname":"montejo","email":"gemontejo@gmail.com"}

Note, in your url, you can also include your params but you should add an escape character to & so your shell will not interpret it as a & (background) option, i.e.:

ghcr POST https://127.0.0.1:2121/users?id=abc123\&age=25 ["Content-type":"application/json"] {"firstname":"german","lastname":"montejo","email":"gemontejo@gmail.com"}

Note this project is still in WIP. Some headers might not work, which are yet to be tested. Project will be updated soon.

This project already has a runnable file named: ghcr.
If you want to make changes, then build this project, you can:
go build main.go
Once built, the executable file will be named as main, just change it to gohcr to make it consistent. I will still have to find a way to address this. To make this runnable where ever you want, you can add the project directory to your path variable.
