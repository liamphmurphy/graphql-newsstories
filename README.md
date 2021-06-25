# News Stories
This is just a small project to test implementing a GraphQL server in Golang. Currently the DB is just in-memory.

## Running
``
go run server.go
``

A GraphQL repl should now be running at `http://localhost:8080/`. If the `PORT` env variable is set, the port will be different.

## Queries

### Return all stories

```
{
  stories {
    id
    title
    body
  }
}
```

### Return a specific story by title
```
{
  story(title: "Testing"){
    id
    title
    body
  }
}
```

## Mutations

### Add a new story

```
mutation createStory {
  createStory(input: {storyTitle: "Testing a new story again for a third time", storyBody:"Test Body Content"}) {
    title
  }
}
```