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
}

function signUp() {

}

// to do list
// const form = document.getElementById('todo_form')
// const input = document.getElementById('todo_input')
// const todosUL = document.getElementById('todos')
// const todos = JSON.parse(localStorage.getItem('todos'))


// if(todos) {
//     todos.forEach(todo => addTodo(todo))
// }
//
// form.addEventListener('submit', (e) => {
//     e.preventDefault()
//
//     addTodo()
// })

// function addTodo(todo) {
//     let todoText = input.value
//
//     if(todo) {
//         todoText = todo.text
//     }
//
//     if(todoText) {
//         const todoEl = document.createElement('li')
//         if(todo && todo.completed) {
//             todoEl.classList.add('completed')
//         }
//
//         todoEl.innerText = todoText
//
//         todoEl.addEventListener('click', () => {
//             todoEl.classList.toggle('completed')
//             updateLS()
//         })
//
//         todoEl.addEventListener('contextmenu', (e) => {
//             e.preventDefault()
//
//             todoEl.remove()
//             updateLS()
//         })
//
//         todosUL.appendChild(todoEl)
//
//         input.value = ''
//
//         updateLS()
//     }
// }

// function updateLS() {
//     todosEl = document.querySelectorAll('li')
//
//     const todos = []
//
//     todosEl.forEach(todoEl => {
//         todos.push({
//             text: todoEl.innerText,
//             completed: todoEl.classList.contains('completed')
//         })
//     })
//
//     localStorage.setItem('todos', JSON.stringify(todos))
// }


// // quote
// const quotesContainer = document.getElementById('quotes-container');
// const quoteElement = document.getElementById('quote');

// function getQuote() {
//     fetch('https://api.quotable.io/random')
//         .then(response => response.json())
//         .then(data => {
//             quoteElement.textContent = `"${data.content}" - ${data.author}`;
//         })
//         .catch(error => {
//             console.log(error);
//         });
// }

// getQuote();
//
// setInterval(getQuote, 100000);

