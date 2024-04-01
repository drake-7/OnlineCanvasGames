const socket = new WebSocket(`ws://${document.location.host}/ws/gameplay`)

socket.onopen = (event) => {
    socket.send("jajaja")
}

socket.onmessage = (event) => {
    console.log(event)
}