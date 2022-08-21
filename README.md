# IF-ELSE-Backend-2022
Repository backend IF ELSE 2022

1. [About](#about)
2. [How to deploy](#how-to-deploy)
3. [Links and other documentation](#links-and-other-documentation)

### About
This is the backend repository for the IF-ELSE's web application for 2022.

It is written with Go version 1.19 and uses MariaDB / MySQL as its database.

### How to deploy

1. Copy [.env.example](.env.example), rename the copied file to .env

2. fill in the [.env](.env) file's variables approriately 

3. Open terminal on the project's [directory](.)

4. Type in `go get`

5. To run the project, type in `go run main.go`

6. If you want to deploy this with systemd daemon for auto enable and restart whenever the system is down or rebooted:
    1. go to [Other](./Other) and run `cp -v if-else.service /etc/systemd/system`
    2. copy the entire project to the `/opt` directory.
    3. run `systemctl daemon-reload && systemctl enable --now if-else.service && sleep 2 && systemctl status if-else.service`

### Links and other documentation

1. [Why use the .env file?](https://www.freecodecamp.org/learn/back-end-development-and-apis/basic-node-and-express/use-the--env-file)

2. [GIN's documentation](https://gin-gonic.com/docs/)

3. [GORM's documentation](https://gorm.io/docs/)

4. [GNU GPL-v3 license](./LICENSE)
