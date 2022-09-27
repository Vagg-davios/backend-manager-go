# REST API written in <img src="https://user-images.githubusercontent.com/101106849/192489171-0445c9b5-8aa4-4b72-828f-793232d02cfd.png" width="85" height="35" /> <img src="https://user-images.githubusercontent.com/101106849/192488357-bb6df28e-5e63-426a-9ea6-54a03a28fc6a.png" width="30" height="35" />

## Installation

1. Install [GO](https://go.dev/dl/)
2. Git clone https://github.com/Vagg-davios/rest-api-in-go.git
3. Open your favorite IDE / text-editor inside that folder
4. *cd ./backend*
5. Go run main.go
6. Run your frontend
7. Boom. Enjoy! 🎉

<hr>

## Back-end
This is a backend manager REST API written in GO. It includes the functionality of adding, viewing and deleting items.

- It's a REST API using [Gin](https://pkg.go.dev/github.com/gin-gonic/gin@v1.8.1) for the server & the routing and [net/http](https://pkg.go.dev/net/http) for validating. 
- Communication is done through port **:8080** with Go handling the requests sent from Javascript.
- JSON is used for the request bodies. <br>

Requests available: <br>

> "*GET http://127.0.0.1:8080/items*" - **View all items** <br>
> "*GET http://127.0.0.1:8080/items/3*" - **Get the item with the id of 3** <br>
> "*POST http://127.0.0.1:8080/items + JSON Body*" - **Add that item to the inventory** <br>
> "*PATCH http://127.0.0.1:8080/items/3*" - **Delete the item with the id of 3**

## Front-end
Plain vanilla HTML, JS and CSS were used without any styling etc. purely because this project was a way for me to understand and practice writing in GoLang, but also learn how to write APIs.
UI includes:

- A table with the available inventory, having the item ids, names, prices and quantity of each
- [XMLHttpRequest](https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest) was used to send the requests.

![image](https://user-images.githubusercontent.com/101106849/192492484-187f769c-d11c-4bec-ba47-5ac27667f6c6.png)


