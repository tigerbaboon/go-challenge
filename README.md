# GO CHALLENGE

## 💻 Exercise 1

```bash
go run . ex1
```

## 💻 Exercise 2

```bash
go run . ex2
```

## 💻 Exercise 3

- Http framework with Gin

- Server start on port :8080

```bash
go run . http
```
``` curl --location 'http://localhost:8080/beef/summary' ```

### Response

```bash
{
    "beef": {
        "bresaola": 47,
        "enim": 40,
        "fatback": 38,
        "jowl": 43,
        "meatloaf": 37,
        "pastrami": 48,
        "pork": 177,
        "t-bone": 47
    }
}
```

## 💻 Test

```bash
go test ./tests
```