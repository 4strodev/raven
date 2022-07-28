# Raven
Raven is a go project generator that helps you to create your projects easy and fast.
It is build with [owl](https://github.com/4strodev/owl) and [cobra](https://github.com/spf13/cobra).

## Requirements
The only requirement is [git](https://git-scm.com/).

## Installation
For the moment the best way to install raven is use `go install`.

    go install github.com/4strodev/raven

## Usage
If you want to create a simple project first you need to find or create a template.

Here a give you a couple of templates created by me.

- [go-cli](https://github.com/4strodev/go-cli)
- [raven-fiber-template](https://github.com/4strodev/raven-fiber-template)

For the example I will use fiber tempate. When you create a project you have to pass the module name mandatorily.

    raven create hello_fiber -t https://github.com/4strodev/raven-fiber-template -m github.com/4strodev/hello_fiber -v

I you will use this template frequently it is a good idea download it.
    
    raven download https://github.com/4strodev/raven-fiber-template

Now your template will be in `~/.raven/templates/github.com/4strodev/raven-fiber-template`. And you only have to
pass the path relative to `~/.raven/templates` every time you want to use it.

    raven create hello_fiber -t github.com/4strodev/raven-fiber-template -m github.com/4strodev/hello_fiber -v

## Roadmap
- [ ] Add task commands like npm or makefile to your projects.
