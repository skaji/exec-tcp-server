# exec tcp server

If you execute a tcp server, then it is common that

* listen a low-numbered (<1024) port as `root` user, and
* switch to non-root user such as `nobody`

In some programming languages such as golang, this might be difficult.

`exec-tcp-server` helps this.

# Usage

```
$ exec-tcp-server --addr 0.0.0.0:80 --user nobody your-server arg1 arg2 arg3
```

# Example

Listen 80 port, and switch from `root` to `nobody`:

```
❯ make
cd example && go build -o example
sudo ./exec-tcp-server --addr 0.0.0.0:80 --user nobody example/example
```

# See also

* https://github.com/kazuho/p5-Server-Starter
* https://github.com/lestrrat-go/server-starter

# Author

Shoichi Kaji

# License

MIT

For listener/listener.go, see https://github.com/lestrrat-go/server-starter
