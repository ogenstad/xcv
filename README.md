# xcv - x509 Certificate Viewer

 `xvc` is a tool to display the contents of x509 certificates and certificate chains. It's useful
 if you want to display the contents of a single certificate or a complete chain from Root CA to
 end certificate.

 ### Usage

 ```bash
 cat fullchain.pem | xcv
 ```

 ```bash
 cat server.crt | xcv
 ```

### Missing info

Currently the tool only gives the most basic information about the certificates and additional
extensions will be added as needed.

