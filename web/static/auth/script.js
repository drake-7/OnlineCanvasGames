const InputType = {
    Login: "login",
    Register: "register"
}

function postInput(type){
    const inputData = {
        username: $("#username").val(),
        password: $("#password").val(),
        type: type
    }

    fetch ("/api/userauth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(inputData)
    })
    .then (response => {
        response.text().then(data => $("#info").html(data))
        if(response.status == 200){
            window.location.replace("/")
        }
    })
    .catch (error => {
        console.error(error)
    })
}

$(document).ready(() => {
    $("#login").click(() => {
        postInput(InputType.Login)
    })

    $("#register").click(() => {
        postInput(InputType.Register)
    })
})