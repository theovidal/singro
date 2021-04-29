# singro

A simple Windows macro software written in Golang, with an Electron.js configuration UI

## 🌈 How it works

The aim of singro is to be highly customizable and work on a maximum range of software programs across the system.

The software uses native Windows API, like the `SendInput` method to execute the macros. Each of these is defined in a YAML configuration file defined by the user.

The main code is written in Go, and some parts in C for a better compatibility with standard libraries that interact with the system.

The interface to configure macros uses JavaScript with Vue.js and Vuetify, and Electron to build a native app. The core itself can be run without it.

## 💻 Development

First, check the following requirements:

- Git, for version control
- Golang 1.16 or higher with go-modules for dependencies
- A C compiler and cgo setup
- Node.js 14 or higher with either npm or yarn

Clone the project on your local machine:

```bash
git clone https://github.com/theovidal/singro  # HTTP
git clone git@github.com:theovidal/singro      # SSH
```

Install the node.js dependencies:

```bash
npm i         # NPM
yarn install  # Yarn
```

To run and test the software, simply use `go run .`. To build an executable, use `go build .`.

To run and test the interface, run the `electron:serve` script, and `electron:build` to build the native app.

## 🔐 Credits

- Maintainer: [Théo Vidal](https://github.com/theovidal)
- Cloned from [Highest template](https://github.com/highest-app/template)
- Libraries: Vue.js, Vuetify, Electron

## 📜 License

[GNU GPL v3](./LICENSE)
