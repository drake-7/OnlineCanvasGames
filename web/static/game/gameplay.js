const socket = new WebSocket(`ws://${document.location.host}/ws/gameplay`)

socket.onopen = (event) => {
    console.log("connected")
    socket.send("jajaja")
}

socket.onmessage = (event) => {
    console.log(event)
}