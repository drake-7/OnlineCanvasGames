const canvas = $("#game-canvas")
let gameID = 0
let socket

function connectToSocket(){
    gameID = canvas.data("game-id")
    socket = new WebSocket(`ws://${document.location.host}/ws/gameplay/${gameID}`)

    socket.onopen = (event) => {
        console.log("opened socket connection")
        socket.send("hello websocket!")
    }

    socket.onmessage = (event) => {
        console.log(event)
    }

    socket.onerror = (event) => {
        console.error(event.code)
    }

    socket.onclose = (event) => {
        console.log("closed socket connection")
    }
}

canvas.ready(() => {
    connectToSocket()
})