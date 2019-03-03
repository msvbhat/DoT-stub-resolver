# dot-stub-resolver
A DNS proxy which supports DNS over TLS

This is a DNS stub resolver which reads the input and resolves
the address by sending it over to the Cloudflare's DNS over TLS
DNS resolver.

Note: This is not 100% complete yet, but will be done soonish.

As of now, it has been only tested with `nc`. Also it is written just
for testing purpose. And hence it can not be used with real world
application yet.


### Building and testing

On a Linux machine build the static binary by running the `make build`

And then run the binary `./dot-resolver`

You can test it by using below command.

```bash
echo "example.com" | nc localhost 53
```

### Application in Kubernetes

Ideally in a Kubernetes environment, this should run as a side car which
takes the DNS name and resolves it using DNS over TLS. And the application
container should be configured to use localhost as the DNS resolver.

This can also work on host by running as a daemon set.


### Security implications

1. This doesn't handle the certificates properly yet.


### Areas of improvement

1. Handle the DNS queries properly and also listen on UDP.
1. Return all addresses returned by the upstream DNS resolver.
1. Use caching to improve performace.
1. Handle SRV requetes as well.
1. Use multiple upstream resolver (like Google DNS etc)
1. Make the listen port and address configurable.
1. Improve perf by using long lived tcp connection to upstream
