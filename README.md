# recalbox-manager-go
A web manager for Recalbox, written in Go.

## The project is a complete rewriting of the current web manager. It's WIP and it's not working for the moment.

### If you want the actual manager, [it's here](https://github.com/DjLeChuck/recalbox-manager)

## HOWTO build project
* Use `dep` to install dependencies:
```sh
go get -u github.com/golang/dep/cmd/dep
dep ensure
```
* Compile templates with `go-bindata`:
```sh
go get -u github.com/shuLhan/go-bindata/...
go-bindata templates/...
```
* Compile assets with `bindata`:
```sh
go get -u github.com/kataras/bindata/cmd/bindata
bindata assets/...
```
* Build the project:
```sh
go build
```

## Note about `go-bindata` and `bindata`
The content of `templates` and `assets` directories would be compiled into Go files.
This mean once the project is build, this two directories are not necessary for the project to work.
All in the `templates` directory can be accessible without prefix like this for example: `| {{ render "views/forms/audio.pug" }}`
All in the `assets` directory can be accessible with a `static` prefix like this for example: `/static/css/bootstrap.min.css`
