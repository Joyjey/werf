{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Fetch release history

{{ header }} Syntax

```shell
werf helm history RELEASE_NAME [options]
```

{{ header }} Options

```shell
      --col-width=60:
            Specifies the max column width of output
      --helm-release-storage-namespace='kube-system':
            Helm release storage namespace (same as --tiller-namespace for regular helm, default    
            $WERF_HELM_RELEASE_STORAGE_NAMESPACE, $TILLER_NAMESPACE or 'kube-system')
      --helm-release-storage-type='configmap':
            helm storage driver to use. One of 'configmap' or 'secret' (default                     
            $WERF_HELM_RELEASE_STORAGE_TYPE or 'configmap')
  -h, --help=false:
            help for history
      --home-dir='':
            Use specified dir to store werf cache files and dirs (default $WERF_HOME or ~/.werf)
      --kube-config='':
            Kubernetes config file path (default $WERF_KUBE_CONFIG or $WERF_KUBECONFIG or           
            $KUBECONFIG)
      --kube-config-base64='':
            Kubernetes config data as base64 string (default $WERF_KUBE_CONFIG_BASE64 or            
            $WERF_KUBECONFIG_BASE64 or $KUBECONFIG_BASE64)
      --kube-context='':
            Kubernetes config context (default $WERF_KUBE_CONTEXT)
      --log-color-mode='auto':
            Set log color mode.
            Supported on, off and auto (based on the stdout’s file descriptor referring to a        
            terminal) modes.
            Default $WERF_LOG_COLOR_MODE or auto mode.
      --log-debug=false:
            Enable debug (default $WERF_LOG_DEBUG).
      --log-pretty=true:
            Enable emojis, auto line wrapping and log process border (default $WERF_LOG_PRETTY or   
            true).
      --log-quiet=false:
            Disable explanatory output (default $WERF_LOG_QUIET).
      --log-terminal-width=-1:
            Set log terminal width.
            Defaults to:
            * $WERF_LOG_TERMINAL_WIDTH
            * interactive terminal width or 140
      --log-verbose=false:
            Enable verbose output (default $WERF_LOG_VERBOSE).
  -m, --max=256:
            Maximum number of releases to fetch
      --output='table':
            Output the specified format (json, yaml or table)
      --tmp-dir='':
            Use specified dir to store tmp files and dirs (default $WERF_TMP_DIR or system tmp dir)
```

