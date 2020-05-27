>    
> <div class="details">
> <a href="javascript:void(0)" class="details__summary">build-and-publish & deploy jobs</a>
> <div class="details__content" markdown="1">
> 
> {% raw %}
> ```yaml
> build-and-publish:
>   name: Build and Publish
>   runs-on: ubuntu-latest
>   steps:
> 
>     - name: Checkout code
>       uses: actions/checkout@v2
>       with:
>         fetch-depth: 0
> 
>     - name: Build and Publish
>       uses: flant/werf-actions/build-and-publish@v1
>       with:
>         kube-config-base64-data: ${{ secrets.KUBE_CONFIG_BASE64_DATA }}
> 
> deploy:
>   name: Deploy
>   needs: build-and-publish
>   runs-on: ubuntu-latest
>   steps:
> 
>     - name: Checkout code
>       uses: actions/checkout@v2
>       with:
>         fetch-depth: 0
> 
>     - name: Deploy
>       uses: flant/werf-actions/deploy@v1
>       with:
>         env: production
>         kube-config-base64-data: ${{ secrets.KUBE_CONFIG_BASE64_DATA }}
> ```
> {% endraw %}
> 
> </div>
> </div>