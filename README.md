# Books API (Golang)

A simple RESTful API for managing books, built with Go.  
This application uses **in-memory storage** and clean modular architecture.

## How to Run

### Requirements
- Go 1.22 or newer

### Start the App
This application runs a web server accessible at [http://localhost:8080](http://localhost:8080).

## How to Run

Open your terminal and execute the following command:

```
go run main.go
```

## API Endpoint Documentation

### Create Book
- Method: POST
- URL: /books
- Body (JSON):
```
{
    "title": "How To Create Book",
    "author": "Roberto Mancini",
    "isbn": "999222333111",
    "release_date": "2010-10-01",
    "status": true
}
```
- Response (JSON):
```
{
    "data": {
        "title": "How To Create Book",
        "author": "Roberto Mancini",
        "isbn": "999222333111",
        "release_date": "2010-10-01",
        "status": true
    },
    "message": "Book successfully created",
    "status": 201
}
```

### Get All Book
- Method: GET
- URL: /books?page=1&limit=10
- Response (JSON):
```
{
    "data": [
        {
            "title": "The Art Of Stoicism",
            "author": "Adora Kinara",
            "isbn": "999222333444",
            "release_date": "2023-01-01",
            "status": true
        },
        {
            "title": "The Simplicity Pricinple",
            "author": "Julia Hobsbawm",
            "isbn": "999222333555",
            "release_date": "2020-02-01",
            "status": true
        },
        {
            "title": "The Intelligent Investor",
            "author": "Benjamin Graham",
            "isbn": "999222333666",
            "release_date": "1949-10-31",
            "status": true
        },
        {
            "title": "Rich Dad Poor Dad",
            "author": "Robert T. Kiyosaki & Sharon Lechter",
            "isbn": "999222333777",
            "release_date": "1997-07-08",
            "status": true
        },
        {
            "title": "How To Create Book",
            "author": "Roberto Mancini",
            "isbn": "999222333111",
            "release_date": "2010-10-01",
            "status": true
        },
        {
            "title": "The Psychology of Money",
            "author": "Morgan Housel",
            "isbn": "999111222333",
            "release_date": "2012-11-22",
            "status": true
        }
    ],
    "message": "Get All Book successfully",
    "status": 200
}
```

### Get Book By ISBN
- Method: GET
- URL: /books/{isbn}
- Response (JSON):
```
{
    "data": {
        "title": "The Art Of Stoicism",
        "author": "Adora Kinara",
        "isbn": "999222333444",
        "release_date": "2023-01-01",
        "status": true
    },
    "message": "Get Book successfully",
    "status": 200
}
```

### Update Book By ISBN
- Method: PUT
- URL: /books
- Body (JSON):
```
{
    "title": "The Simplicity Pricinple Power",
    "author": "Julia Hobsbawm P.",
    "isbn": "999222333555",
    "release_date": "2025-05-01",
    "status": true
}
```
- Response (JSON):
```
{
    "data": {
        "title": "The Simplicity Pricinple Power",
        "author": "Julia Hobsbawm P.",
        "isbn": "999222333555",
        "release_date": "2025-05-01",
        "status": true
    },
    "message": "Book successfully updated",
    "status": 200
}
```

### Delete Book By ISBN
- Method: DELETE
- URL: /books/{isbn}
- Response (JSON):
```
{
    "data": null,
    "message": "Book successfully deleted",
    "status": 200
}
```

### Restore Book By ISBN
- Method: PUT
- URL: /books/restore/{isbn}
- Response (JSON):
```
{
    "data": {
        "title": "The Art Of Stoicism",
        "author": "Adora Kinara",
        "isbn": "999222333444",
        "release_date": "2023-01-01",
        "status": true
    },
    "message": "Book successfully restore",
    "status": 200
}
```

