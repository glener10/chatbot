# **Chatbot**

<div>
  <div>
    <table>
      <thead>
        <tr>
          <th colspan="4">Repository Informations</th>
          <th colspan="2">Open Tasks</th>
          <th colspan="2">Copyright</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td><img src="https://img.shields.io/github/repo-size/glener10/authentication" alt="GitHub Repo size"></td>
          <td><img src="https://img.shields.io/github/stars/glener10/authentication" alt="GitHub Stars"></td>
          <td><img src="https://img.shields.io/github/forks/glener10/authentication" alt="Forks"></td>
          <td><img src="https://github.com/glener10/authentication/workflows/go/badge.svg" alt="Build Status"></td>
          <td><img src="https://img.shields.io/bitbucket/issues/glener10/authentication" alt="Open Issues"></td>
          <td><img src="https://img.shields.io/bitbucket/pr-raw/glener10/authentication" alt="Open Pull Requests"></td>
          <td><img src="https://img.shields.io/github/license/glener10/authentication" alt="License"></td>
          <td><img src="https://img.shields.io/github/contributors/glener10/authentication.svg" alt="Contributors"></td>
        </tr>
      </tbody>
    </table>
  </div>

<p align="center"> 🚀 Project created to be my chatbot </p>

<h3>🏁 Table of Contents</h3>

<br>

===================

<!--ts-->

💻 [Dependencies and Environment](#dependenciesandenvironment)

🚀 [Installing](#installing)

🧹 [Formatting the Code](#formatting)

🧪 [Testing](#testing)

☕ [Using](#using)

🔒 [License](#license)

👷 [Author](#author)

<!--te-->

===================

<div id="dependenciesandenvironment"></div>

## 💻 **Dependencies and Environment**

My dependencies and versions

[**Go**](https://golang.org/): go version go1.24.5

<div id="installing"></div>

## 🚀 **Installing**

- To install the dependencies you can run the following command in the root folder:

```
$ go mod tidy
$ go mod download
```

**OBS**: We have the development [.env](.env) file committed to the project, but you can change it as you see fit

<div id="formatting"></div>

## 🧹 **Formatting the Code**

To check the code format you will need [instal golangci-lint](https://golangci-lint.run/welcome/install/) and run the following command in the root folder:

```
$ golangci-lint run
```

To format all files, use:

```
go fmt ./...
```

<div id="testing"></div>

## 🧪 **Testing**

To exec all the tests run the following command in the root folder:

```
$ go test -p 1 ./src/...
```

You can add the "**-v**" flag to see detailed output

```
$ go test -v -p 1 ./src/...
```

<div id="using"></div>

## ☕ **Using**

First, check the [dependencies](#dependenciesandenvironment) and the [installation](#installing) process:

Going to _root_ folder and exec:

```
$ go run .\server.go
```

Now you can open [http://localhost:1323](http://localhost:1323) with your browser to see the result.

<div id="license"></div>

## 🔒 **License**

Projeto contêm [GNU GENERAL PUBLIC LICENSE](LICENSE).

<div id="author"></div>

#### **👷 Author**

Made by Glener Pizzolato! 🙋

[![Linkedin Badge](https://img.shields.io/badge/-Glener-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/glener-pizzolato/)](https://www.linkedin.com/in/glener-pizzolato-6319821b0/)
[![Gmail Badge](https://img.shields.io/badge/-glenerpizzolato@gmail.com-c14438?style=flat-square&logo=Gmail&logoColor=white&link=mailto:glenerpizzolato@gmail.com)](mailto:glenerpizzolato@gmail.com)
