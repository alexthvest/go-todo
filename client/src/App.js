import React, {useEffect, useState} from "react"

import TodoList from "./components/todos/TodoList"
import TodoForm from "./components/todos/TodoForm"

export default function App() {
    const [todos, setTodos] = useState([])

    useEffect(() => {
        fetchTodos()
    }, [])

    async function fetchTodos() {
        const response = await fetch("/todos")
        const todos = await response.json()

        setTodos(todos)
    }

    function onTodoAdded(todo) {
        setTodos([...todos, todo])
    }

    return (
        <div>
            <h1>Go Todo</h1>
            <p>Welcome to Todo Application written on Go</p>

            <TodoForm onAdded={onTodoAdded} />
            <TodoList todos={todos} />
        </div>
    );
}
