# 5 Mistakes C/CPP Devs make writing Go
My Gophercon 2018 Talk

## Topics

This talk brings my experiences coming from a predominantly C/C++ background and moving to Go. Talking about certain biases I had earlier and general newbie mistakes moving to Go. I'll talk about the following concepts:

### Escape Analysis
- M[0] = New doesnt' imply heap, var doesn't imply stack

### Memory Leaks

- M[1] = Do not defer in an infinite Loop

- M[2] = Pointer in non visible(but accessible) slice portion

- Note: memory safety  and bytes.Split bytes.Fields


### Goroutine Leaks
- M[4] = Error handling with Channels where cap(errc) < # goroutines


### Error Handling
- M[5] = Errors are not strings, but much more

- Handling the same error multiple times(logging and returning)

- Note: Why Wrap is cool(Preserves context of an error)



## Installation

1. Clone this repo in your `$GOPATH`
```
$ git clone https://github.com/Ch3ck/5-mistakes-c-cpp-devs-make-writing-go
```

2. Install the Go present tool
 ```
 go get -u -v golang.org/x/tools/cmd/present
 ```

3. Run the present tool
 ```
 cd 5-mistakes-c-cpp-devs-make-writing-go && present
 ```

The slides will be available at [http://127.0.0.1:3999/](http://127.0.0.1:3999/5-mistakes-c-cpp-devs-make-writing-go.slide#1)



## Issues

- New ideas are always welcomed, so you can submit an [issue](https://github.com/Ch3ck/5-mistakes-c-cpp-devs-make-writing-go/issues) with a question comment or make a [pull request](https://github.com/Ch3ck/5-mistakes-c-cpp-devs-make-writing-go/pulls) if you see something better that can be done. I'll cultimate all these ideas on to a blog post on [medium](https://medium.com/@checko) and you should find the slides on my [SlideDeck](https://speakerdeck.com/ch3ck)

## Author
- Ping me on Twitter [@nyah_check](https://twitter.com/nyah_check)
