# singro

**S**imple W**in**dows and **G**olang Mac**r**o s**o**ftware

A simple project, just for fun.

## 🌈 How it works

The aim of singro is to be highly customizable and work on a maximum range of software programs across the system.

The software uses native Windows API, like the `SendInput` method to execute the macros. Each of these is defined in a YAML configuration file defined by the user.

The main code is written in Go, and some parts in C for a better compatibility with standard libraries that interact with the system.

## 💻 Development

First, check the following requirements:

- Git, for version control
- Golang 1.16 or higher with go-modules for dependencies
- A C compiler and cgo setup

Clone the project on your local machine:

```bash
git clone https://github.com/theovidal/singro  # HTTP
git clone git@github.com:theovidal/singro      # SSH
```

To run and test the software, simply use `go run .`. To build an executable, use `go build .`.

## 🔐 Credits

- Maintainer: [Théo Vidal](https://github.com/theovidal)

## 📜 License

[GNU GPL v3](./LICENSE)
