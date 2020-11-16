import React from "react"

export default function TodoList({todos}) {
    return (
        <ul className="list-group">
            {todos.map((todo, index) =>
                <li key={index} className={`list-group-item ${!todo.completed || 'active'}`}>{todo.title}</li>
            )}
        </ul>
    )
}