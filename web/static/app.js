const KRAKEN_API = "http://100.85.100.49:8888"

async function signIn() {
    let email = document.getElementById("email").value
    let password = document.getElementById("password").value
    let res = await fetch(KRAKEN_API + "/user/signin", {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        body: JSON.stringify({
            "email": email,
            "password": password
        }),
        method: "POST",
        // mode: "no-cors",
        credentials: "include"
    }).then(res => {
        if (!res.ok) {
            return res.json().then(err => {
                throw new Error(`${err.message}`)
            })
        }
        window.location.href = "/todo.html";
    }).catch(err => {
        console.log(err)
        alert(err)
    });
}

async function signUp() {
    let username = document.getElementById("username").value;
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;
    let res = await fetch(KRAKEN_API + "/user/signup", {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        body: JSON.stringify({
            "username": username,
            "email": email,
            "password": password
        }),
        method: "POST",
        // mode: "no-cors",
        credentials: "include"
    }).then(res => {
        if (!res.ok) {
            return res.json().then(err => {
                throw new Error(`${err.message}`)
            })
        }
        alert("Account is successfully created. Redirecting to login...")
        window.location.href = "/signin.html"
    }).catch(err => {
        console.log(err)
        alert(err)
    });
}

async function resetPassword() {

}