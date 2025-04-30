# hajimeteno-golang-specification

```shell
❯ go run main.go
*T.Bar()
T.Foo()
*T.Bar()
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4809b1]

goroutine 1 [running]:
main.main()
        ./hajimeteno-golang-specification/main.go:28 +0xd1
exit status 2

```
