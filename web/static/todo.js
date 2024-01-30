// to do list
const form = document.getElementById('todo_form')
const input = document.getElementById('todo_input')
const todosUL = document.getElementById('todos')
let todos = []

getAllTodos()

form.addEventListener('submit', (e) => {
    e.preventDefault()
    addTodo()
})

function addTodo() {
    fetch(KRAKEN_API + '/todo', {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
            "Title": input.value,
        })
    }).then(async res => {
        if (!res.ok) {
            const err = await res.json()
            throw new Error(`${err.message}`)
        }
        
        return res.json()
    }).then(_ => {
        input.value = ""
        getAllTodos()
        renderTodos()
    }
    ).catch(err => {
        console.log(err)
        alert(err)
    });
}

function renderTodos() {
    while(todosUL.hasChildNodes()) {
        todosUL.removeChild(todosUL.childNodes[0])
    }

    if (todos == null) return

    todos.forEach(todo => {
        const todoLi = document.createElement("li")

        todoLi.innerText = todo.Title
        todoLi.classList.add('completed')

        if (!todo.Completed) {
            todoLi.classList.toggle("completed")
        }

        todoLi.addEventListener("click", () => changeCompletion(todo, todoLi))

        todoLi.addEventListener("contextmenu", (e) => deleteTodo(e, todo))

        todosUL.appendChild(todoLi)
    })
}

function deleteTodo(e, todo) {
    e.preventDefault()

    fetch(KRAKEN_API + '/delete/' + todo.Id, {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        method: "GET",
        credentials: "include",
    }).then(async res => {
        if (!res.ok) {
            const err = await res.json()
            throw new Error(`${err.message}`)
        }
        
        return res.json()
    }).then(_ => {
        getAllTodos()
        renderTodos()
    }
    ).catch(err => {
        console.log(err)
        alert(err)
    });
}

function changeCompletion(todo, htmlElem) {
    htmlElem.classList.toggle("completed")
    
    fetch(KRAKEN_API + '/update/' + todo.Id, {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        method: "GET",
        credentials: "include",
    }).then(async res => {
        if (!res.ok) {
            const err = await res.json()
            throw new Error(`${err.message}`)
        }
        
        return res.json()
    }).then(_ => {
        getAllTodos()
        renderTodos()
    }
    ).catch(err => {
        console.log(err)
        alert(err)
    });
}

// quote
const quotesContainer = document.getElementById('quotes-container');
const quoteElement = document.getElementById('quote');

function getQuote() {
    fetch('https://api.quotable.io/random')
        .then(response => response.json())
        .then(data => {
            quoteElement.textContent = `"${data.content}" - ${data.author}`;
        })
        .catch(error => {
            console.log(error);
        });
}

getQuote();

setInterval(getQuote, 100000);

function getAllTodos() {
    fetch(KRAKEN_API + '/todos', {
        headers: {
            "accept": "application/json",
            "accept-language": "en-US,en;q=0.9",
        },
        method: "GET",
        credentials: "include",
    }).then(async res => {
        if (!res.ok) {
            const err = await res.json()
            throw new Error(`${err.message}`)
        }
        
        return res.json()
    }).then(data => {
        todos = data.Data.Todos
        renderTodos()
    }
    ).catch(err => {
        console.log(err)
        alert(err)
    });
}