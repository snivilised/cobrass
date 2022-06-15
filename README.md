# 🐲 ___Cobrass: assistant for cli applications using cobra___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/cobrass.svg)](https://pkg.go.dev/github.com/snivilised/cobrass)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" /></a>
</p>

## 🔰 Introduction

_[Cobra](https://cobra.dev/) is an excellent framework for the development of command line applications, but there are few aspects that could do with being made easier to work with. This package aims to fullfil this purpose, especially in regards to creation of commands, encapulating commands into a container and providing an export mechanism to re-create cli data in a form that is free from cobra (and indeed cobrass) abstractions. The aim of this last aspect to to be able to inject data into the core of an application in a way that removes tight coupling to the cobra framework, which is achieved by representing data only in terms of client defined (native) abstractions._

___Status___: 💤 not yet published
