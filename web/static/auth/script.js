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

    fetch ("/api/user", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(inputData)
    })
    .then (response => {
        response.json().then(data => $("#info").html(data.message))
        if(response.status == 200){
            window.location.reload()
        }
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