---
title: Работа с файлами
sidebar: applications-guide
permalink: documentation/guides/applications-guide/template/050-files.html
layout: guide
toc: false
---

{% filesused title="Файлы, упомянутые в главе" %}
- .helm/templates/deployment.yaml
- .helm/secret-values.yaml
- config/initializers/aws.rb
- Gemfile
{% endfilesused %}

В этой главе мы настроим в нашем базовом приложении работу с пользовательскими файлами. Для этого нам нужно персистентное хранилище.

В идеале — нужно добиться, чтобы приложение было stateless, а данные хранились в S3-совместимом хранилище, например minio или aws s3. Это обеспечивает простое масштабирование, работу в HA режиме и высокую доступность.

{% offtopic title="А есть какие-то способы кроме S3?" %}
Первый и более общий способ — это использовать как [volume](https://kubernetes.io/docs/concepts/storage/volumes/) хранилище [NFS](https://kubernetes.io/docs/concepts/storage/volumes/#nfs), [CephFS](https://kubernetes.io/docs/concepts/storage/volumes/#cephfs) или [hostPath](https://kubernetes.io/docs/concepts/storage/volumes/#hostpath).

Мы не рекомендуем этот способ, потому что при возникновении неполадок с такими типами volume’ов мы будем влиять на работоспособность всего докера, контейнера и демона docker в целом, тем самым могут пострадать приложения которые даже не имеют отношения к вашему приложению.

Более надёжный путь — пользоваться S3. Таким образом мы используем отдельный сервис, который имеет возможность масштабироваться, работать в HA режиме, и иметь высокую доступность. Можно воспользоваться cloud решением, таким, как AWS S3, Google Cloud Storage, Microsoft Blobs Storage и т.д.

Если мы будем сохранять файлы какой - либо директории у приложения запущенного в kubernetes - то после перезапуска контейнера все изменения пропадут.
{% endofftopic %}

Данная настройка производится полностью в рамках приложения, а нам остается только передать необходимые переменные окружения при запуске приложения.

Добавим в `requirements.txt` зависимости для работы с s3 aws:

{% snippetcut name="requirements.txt" url="#" %}
{% raw %}
```
django-storages==1.9.1
```
{% endraw %}
{% endsnippetcut %}

Далее необходимо задать параметры для работы с s3 в файле настроек, подробнее про это можно прочитать в [документации пакета](https://django-storages.readthedocs.io/en/latest/backends/amazon-S3.html).

Для работы с S3 необходимо пробросить в ключи доступа в приложение. Для этого стоит использовать [механизм секретных переменных](https://ru.werf.io/documentation/reference/deploy_process/working_with_secrets.html). *Вопрос работы с секретными переменными рассматривался подробнее, [когда мы делали базовое приложение](020-basic.html#secret-values-yaml)*

{% snippetcut name="secret-values.yaml (расшифрованный)" url="#" %}
{% raw %}
```yaml
app:
  aws_access_key_id:
    _default: EXAMPLEKVFOOOWWPYA
  aws_secret_access_key:
    _default: exampleBARZHS3sRew8xw5hiGLfroD/b21p2l
  s3_bucket:
    _default: my-s3-development
```
{% endraw %}
{% endsnippetcut %}

После того, как значения корректно прописаны и зашифрованы — мы можем пробросить соответствующие значения в Deployment.

{% snippetcut name="deployment.yaml" url="#" %}
{% raw %}
```yaml
        - name: S3_BUCKET
          value: {{ pluck .Values.global.env .Values.app.s3_bucket | first | default .Values.app.s3_bucket._default }}
        - name: AWS_ACCESS_KEY_ID
          value: {{ pluck .Values.global.env .Values.app.aws_access_key_id | first | default .Values.app.aws_access_key_id._default }}
        - name: AWS_SECRET_ACCESS_KEY
          value: {{ pluck .Values.global.env .Values.app.aws_secret_access_key | first | default .Values.app.aws_secret_access_key._default }}
```
{% endraw %}
{% endsnippetcut %}

Примеры использования пакеты `django-storages` можно посмотреть в [документации](https://django-storages.readthedocs.io/en/latest/backends/amazon-S3.html), либо в [отдельной статье](https://simpleisbetterthancomplex.com/tutorial/2017/08/01/how-to-setup-amazon-s3-in-a-django-project.html).
TODO: надо дать отсылку на какой-то гайд, где описано, как конкретно использовать гем aws-sdk. Мало же просто его установить — надо ещё как-то юзать в коде.

<div>
    <a href="060-email.html" class="nav-btn">Далее: Работа с электронной почтой</a>
</div>