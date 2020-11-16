import React, {useState} from "react"

export default function TodoForm({todos, onAdded}) {
    const [title, setTitle] = useState("")

    function addTodo(e) {
        e.preventDefault()
        if (!title) return

        fetch("/todos", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                title,
                completed: false
            })
        })
            .then(r => r.json())
            .then(todo => {
                setTitle("")
                onAdded && onAdded(todo)
            })
    }

    function handleTitleChange(e) {
        setTitle(e.target.value)
    }

    return (
        <form action="POST" className="form-inline mb-4" onSubmit={addTodo}>
            <div className="form-group">
                <input className="form-control mr-2" placeholder="Title..." value={title} onChange={handleTitleChange}/>
                <button type="submit" className="btn btn-primary">Add</button>
            </div>
        </form>
    )
}