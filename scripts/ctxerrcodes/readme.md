# ctxerrcodes

This is a script that will tell you if your `ctxerr` error codes don't match a format that you want. It can also update the codes automatically in some cases.

Install and run

```bash
go install
ctxerrcodes -format uuid -location . -fix true
```

There is also the format of `uppercase` if you want you don't have uuids.
