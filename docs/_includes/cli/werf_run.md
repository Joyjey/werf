{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Run container for specified project image

{{ header }} Syntax

```shell
werf run [options] [IMAGE_NAME] [-- COMMAND ARG...]
```

{{ header }} Examples

```shell
  # Run specified image
  $ werf run --stages-storage :local application

  # Run image with predefined docker run options and command for debug
  $ werf run --stages-storage :local --shell

  # Run image with specified docker run options and command
  $ werf run --stages-storage :local --docker-options="-d -p 5000:5000 --restart=always --name registry" -- /app/run.sh

  # Print a resulting docker run command
  $ werf run --stages-storage :local --shell --dry-run
  docker run -ti --rm image-stage-test:1ffe83860127e68e893b6aece5b0b7619f903f8492a285c6410371c87018c6a0 /bin/sh
```

{{ header }} Options

```shell
      --allow-git-shallow-clone=false:
            Sign the intention of using shallow clone despite restrictions (default                 
            $WERF_ALLOW_GIT_SHALLOW_CLONE)
      --bash=false:
            Use predefined docker options and command for debug
      --config='':
            Use custom configuration file (default $WERF_CONFIG or werf.yaml in working directory)
      --config-templates-dir='':
            Change to the custom configuration templates directory (default                         
            $WERF_CONFIG_TEMPLATES_DIR or .werf in working directory)
      --dir='':
            Use custom working directory (default $WERF_DIR or current directory)
      --docker-config='':
            Specify docker config directory path. Default $WERF_DOCKER_CONFIG or $DOCKER_CONFIG or  
            ~/.docker (in the order of priority)
            Command needs granted permissions to read and pull images from the specified stages     
            storage
      --docker-options='':
            Define docker run options (default $WERF_DOCKER_OPTIONS)
      --dry-run=false:
            Indicate what the command would do without actually doing that (default $WERF_DRY_RUN)
      --git-unshallow=false:
            Convert project git clone to full one (default $WERF_GIT_UNSHALLOW)
  -h, --help=false:
            help for run
      --home-dir='':
            Use specified dir to store werf cache files and dirs (default $WERF_HOME or ~/.werf)
      --insecure-registry=false:
            Use plain HTTP requests when accessing a registry (default $WERF_INSECURE_REGISTRY)
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
      --log-project-dir=false:
            Print current project directory path (default $WERF_LOG_PROJECT_DIR)
      --log-quiet=false:
            Disable explanatory output (default $WERF_LOG_QUIET).
      --log-terminal-width=-1:
            Set log terminal width.
            Defaults to:
            * $WERF_LOG_TERMINAL_WIDTH
            * interactive terminal width or 140
      --log-verbose=false:
            Enable verbose output (default $WERF_LOG_VERBOSE).
      --repo-docker-hub-password='':
            Common Docker Hub password for any stages storage or images repo specified for the      
            command (default $WERF_REPO_DOCKER_HUB_PASSWORD)
      --repo-docker-hub-token='':
            Common Docker Hub token for any stages storage or images repo specified for the command 
            (default $WERF_REPO_DOCKER_HUB_TOKEN)
      --repo-docker-hub-username='':
            Common Docker Hub username for any stages storage or images repo specified for the      
            command (default $WERF_REPO_DOCKER_HUB_USERNAME)
      --repo-github-token='':
            Common GitHub token for any stages storage or images repo specified for the command     
            (default $WERF_REPO_GITHUB_TOKEN)
      --repo-implementation='':
            Choose common repo implementation for any stages storage or images repo specified for   
            the command.
            The following docker registry implementations are supported: ecr, acr, default,         
            dockerhub, gcr, github, gitlab, harbor, quay.
            Default $WERF_REPO_IMPLEMENTATION or auto mode (detect implementation by a registry).
      --shell=false:
            Use predefined docker options and command for debug
      --skip-tls-verify-registry=false:
            Skip TLS certificate validation when accessing a registry (default                      
            $WERF_SKIP_TLS_VERIFY_REGISTRY)
      --ssh-key=[]:
            Use only specific ssh key(s).
            Can be specified with $WERF_SSH_KEY* (e.g. $WERF_SSH_KEY_REPO=~/.ssh/repo_rsa",         
            $WERF_SSH_KEY_NODEJS=~/.ssh/nodejs_rsa").
            Defaults to $WERF_SSH_KEY*, system ssh-agent or ~/.ssh/{id_rsa|id_dsa}, see             
            https://werf.io/documentation/reference/toolbox/ssh.html
  -s, --stages-storage='':
            Docker Repo to store stages or :local for non-distributed build (only :local is         
            supported for now; default $WERF_STAGES_STORAGE environment).
            More info about stages: https://werf.io/documentation/reference/stages_and_images.html
      --stages-storage-repo-docker-hub-password='':
            Docker Hub password for stages storage (default                                         
            $WERF_STAGES_STORAGE_REPO_DOCKER_HUB_PASSWORD, $WERF_REPO_DOCKER_HUB_PASSWORD)
      --stages-storage-repo-docker-hub-token='':
            Docker Hub token for stages storage (default                                            
            $WERF_STAGES_STORAGE_REPO_DOCKER_HUB_TOKEN, $WERF_REPO_DOCKER_HUB_TOKEN)
      --stages-storage-repo-docker-hub-username='':
            Docker Hub username for stages storage (default                                         
            $WERF_STAGES_STORAGE_REPO_DOCKER_HUB_USERNAME, $WERF_REPO_DOCKER_HUB_USERNAME)
      --stages-storage-repo-github-token='':
            GitHub token for stages storage (default $WERF_STAGES_STORAGE_REPO_GITHUB_TOKEN,        
            $WERF_REPO_GITHUB_TOKEN)
      --stages-storage-repo-implementation='':
            Choose repo implementation for stages storage.
            The following docker registry implementations are supported: ecr, acr, default,         
            dockerhub, gcr, github, gitlab, harbor, quay.
            Default $WERF_STAGES_STORAGE_REPO_IMPLEMENTATION, $WERF_REPO_IMPLEMENTATION or auto     
            mode (detect implementation by a registry).
  -S, --synchronization='':
            Address of synchronizer for multiple werf processes to work with a single stages        
            storage (default :local if --stages-storage=:local or kubernetes://werf-synchronization 
            if non-local stages-storage specified or $WERF_SYNCHRONIZATION if set). The same        
            address should be specified for all werf processes that work with a single stages       
            storage. :local address allows execution of werf processes from a single host only.
      --tmp-dir='':
            Use specified dir to store tmp files and dirs (default $WERF_TMP_DIR or system tmp dir)
      --virtual-merge=false:
            Enable virtual/ephemeral merge commit mode when building current application state      
            ($WERF_VIRTUAL_MERGE by default)
      --virtual-merge-from-commit='':
            Commit hash for virtual/ephemeral merge commit with new changes introduced in the pull  
            request ($WERF_VIRTUAL_MERGE_FROM_COMMIT by default)
      --virtual-merge-into-commit='':
            Commit hash for virtual/ephemeral merge commit which is base for changes introduced in  
            the pull request ($WERF_VIRTUAL_MERGE_INTO_COMMIT by default)
```

