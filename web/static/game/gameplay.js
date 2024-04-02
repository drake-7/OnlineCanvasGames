const socket = new WebSocket(`ws://${document.location.host}/ws/gameplay`)
const canvas = $("#game-canvas")
let gameID = 0

socket.onopen = (event) => {
    canvas.ready(() => {
        gameID = canvas.data("game-id")
        socket.send("i am connected to " + gameID)
    })
}

socket.onmessage = (event) => {
    console.log(event)
}