# League Backend Challenge

This is a simple API for the League Backend Challenge, written in GoLang. The webservice can perform the following operations:

|Endpoint |Data    |Response type|
| ------- |------- |------------ |
|/echo    |csv matrix file as `multipart/form-data`|square matrix representation|
|/invert  |csv matrix file as `multipart/form-data`|square matrix representation |
|/flatten |csv matrix file as `multipart/form-data`|single line matrix representation|
|/sum     |csv matrix file as `multipart/form-data`|single number|
|/multiply|csv matrix file as `multipart/form-data`|single number|

## Data
The input file to these functions is a matrix represented as a csv file, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

## Simple usage
send a simple request to the web server
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## Building and Running the web server from Docker

### building the web server
```
docker build -t league-challenge .
```

### running the web server
```
docker run -it -p 8080:8080 league-challenge:latest
```

### running the test
```
docker run -it league-challenge:latest go test ./...
```

### Developing form inside the container

to be able to develop the web server without always rebuilding or restarting the docker container, run the container with this command:
```
docker run -it -p 8080:8080 -v /your/path/to/the/repo/backend:/go/src/league/challenge/backend league-challenge:latest ./dev.sh
```

This is accomplished by using the https://github.com/pilu/fresh project and docker volumes mapping. Fresh will watch for file events, and every time you create/modify/delete a file, it will build and restart the application automatically. Any compilation error will be printed in the terminal windows running the docker container, watch out if using the -d flag.

#### TIP
To run the unit tests from inside the container without stopping the web server, use this command in a separate terminal to start a shell in the dev container:
```
docker exec -it container_name sh
```
and then run the tests as needed with
```
go test ./...
```

## Run web server localy
```
go run .
```

## Send request


## Operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given) 
    - Return the matrix as a string in matrix format.

    ```
    curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
    ```

    Result:
    
    ```
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert 
    - Return the matrix as a string in matrix format where the columns and rows are inverted

    ```
    curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"
    ```

    Result:
    
    ```
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.

    ```
    curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"
    ```

    Result:
    
    ```
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix

    ```
    curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
    ```

    Result:
    
    ```
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix

    ```
    curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"
    ```

    Result:
    
    ```
    362880
    ``` 
