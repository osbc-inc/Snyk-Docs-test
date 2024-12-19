# 번들 푸시하기

사용자 정의 규칙 번들을 생성한 후, `push` 명령어를 사용하여 지원되는 OCI 레지스트리 중 하나로 자동으로 배포할 수 있습니다.

```
snyk-iac-rules push -r docker.io/example/test bundle.tar.gz
```

{% hint style="info" %}
`snyk-iac-rules push` 명령을 실행하기 전에 Docker로 `docker login` 명령을 사용하여 컨테이너 레지스트리에 먼저 로그인해야 합니다.
{% endhint %}

Snyk은 [OCI 아티팩트 사양](https://github.com/opencontainers/artifacts)을 지원하는 OCI 레지스트리를 사용하며 이를 수행하기 위해 [ORAS](https://github.com/deislabs/oras)를 활용합니다. 현재 지원되는 레지스트리는 다음과 같습니다.

* [Google Container Registry](https://cloud.google.com/container-registry)
* [DockerHub](https://hub.docker.com)
* [Elastic Container Registry](https://aws.amazon.com/ecr/)
* [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/)
* [JFrog Artifactory ](https://www.jfrog.com/confluence/display/JFROG/Docker+Registry)참고: OCI 아티팩트는 Artifactory v7.11.1 이상에서 지원됩니다.
* [Harbor](https://goharbor.io)
* [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

{% hint style="info" %}
Snyk은 보안되지 않은 레지스트리를 지원하지 않습니다. Snyk가 지원하는 유일한 프로토콜은 HTTPS입니다.
{% endhint %}

명령어를 실행한 후, 사용자 정의 규칙 번들이 `latest` 태그를 사용하여 OCI 레지스트리로 푸시될 것입니다.

번들에 버전을 부여하려면 자체 태그를 제공할 수도 있습니다:

```
snyk-iac-rules push -r docker.io/example/test:v0.0.1 bundle.tar.gz
```

이제 [새로 빌드된 사용자 정의 번들로 snyk iac test를 실행할 수 있습니다.](../use-iac-custom-rules-with-cli/)