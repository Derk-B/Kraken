const KRAKEN_API = "http://localhost:8080"

function signIn() {
    let username = document.getElementById("username").value
    let password = document.getElementById("password").value
    let res = fetch(KRAKEN_API, {
        "headers": {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        "body": {
            "username": username,
            "password": password
        },
        "method": "POST",
        "mode": "cors",
        "credentials": "include"
    }).then(res => res.json())
    // Process the response
}

function signUp() {
    let username = document.getElementById("username").value;
    let email = document.getElementById("email").value;
    let password = document.getElementById("password").value;
    let res = fetch(KRAKEN_API, {
        "headers": {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        "body": {
            "username": username,
            "email": email,
            "password": password
        },
        "method": "POST",
        "mode": "cors",
        "credentials": "include"
    }).then(res => res.json())
    // Process the response
}
