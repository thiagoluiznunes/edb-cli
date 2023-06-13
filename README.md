# EDB Kubernetes CLI

The objetctive is to create a CLI application that implements ‘kubectl get pods’ using the respective API.
## Requirements ###

It is necessary to install previously Go language and Make.

* **[Golang 1.9.x](https://go.dev/doc/install)** :white_check_mark:
* **[Make](https://www.gnu.org/software/make/#download)** :white_check_mark:

## Release EDB CLI
- After installing the requirements run the followed command to build a new release of the RPN cli.
```console
$ make release
```
## Local path must be reachable
- After that, you need to add the /usr/local/bin path in your $PATH environment to expose the path of the new software for the whole thesystem.
```console
$ export PATH=$PATH:/usr/local/bin
```
**Info**: To apply this new PATH release in every new terminal session, you will need to update it in your **.bash_profile** or if your are using zsh you will need add in **.zshrc**. Example below:
```console
$ export PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
```

## Check installation
- You can check the installation running the command in your terminal:
```console
$ edbctl -h
```
