# fake_proc

For some scripts I am working on i need to fake proccesses in a consistant way so that tools like `pidof` show them correctly

## Usage

Compile this code

```shell
myprog="httpd"
CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o "${myprog}" main.go
mv "${myprog}" /path/to/
```

Run it with options

```shell
$ /path/to/httpd -D &
```

Check Proccess List

```shell
$ ps aux | grep httpd
vscode    6020  0.0  0.0 1225792 2940 pts/0    Sl   17:15   0:00 /path/to/httpd -D
$ pidof httpd
6020
```
