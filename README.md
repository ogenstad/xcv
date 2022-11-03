# xcv - x509 Certificate Viewer

 `xvc` is a tool to display the contents of x509 certificates and certificate chains. It's useful
 if you want to display the contents of a single certificate or a complete chain from Root CA to
 end certificate.

 ### Usage

 ```bash
 cat fullchain.pem | xcv

 cat server.crt | xcv
 ```

 Alternatively it's possible to run `xcv` without piping in any input and instead pasting the
 certificate info and type Ctrl+d (or Ctrl+z on Windows).

### Missing info

Currently the tool only gives the most basic information about the certificates and additional
extensions will be added as needed.

### Try it out

Grab the latest release from https://github.com/ogenstad/xcv/releases or install directly from Go:

```bash
go install github.com/ogenstad/xcv@latest
```
