## kubectl-shovel trace

Get dotnet-trace results

### Synopsis

This subcommand will capture 10 seconds of runtime events with dotnet-trace tool for running in k8s appplication.
Result will be saved locally in nettrace format so you'll be able to convert it and analyze with appropriate tools.
You can find more info about dotnet-trace tool by the following links:

	* https://github.com/dotnet/diagnostics/blob/master/documentation/dotnet-trace-instructions.md
	* https://docs.microsoft.com/en-us/dotnet/core/diagnostics/dotnet-trace

```
kubectl-shovel trace [flags]
```

### Examples

```
The only required flag is `--pod-name`. So you can use it like this:

	kubectl shovel trace --pod-name my-app-65c4fc589c-gznql

Use `-o`/`--output` to define name of dump file:

	kubectl shovel trace --pod-name my-app-65c4fc589c-gznql -o ./myapp.trace

Also use `-n`/`--namespace` if your pod is not in current context's namespace:

	kubectl shovel trace --pod-name my-app-65c4fc589c-gznql -n default

One of the resulting trace usage examples is converting it to speedscope format:

	dotnet trace convert myapp.trace --format Speedscope

And then analyzing it with https://www.speedscope.app/
```

### Options

```
      --as string                      Username to impersonate for the operation
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --cache-dir string               Default cache directory (default "/Users/signal/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
  -h, --help                           help for trace
      --image string                   Image of dumper to use for job (default "dodopizza/kubectl-shovel-dumper:undefined")
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
  -n, --namespace string               If present, the namespace scope for this CLI request
  -o, --output string                  Output file (default "./1609409397.trace")
      --pod-name string                Target pod
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
```

### SEE ALSO

* [kubectl-shovel](kubectl-shovel.md)	 - Get diagnostics from running in k8s dotnet application

