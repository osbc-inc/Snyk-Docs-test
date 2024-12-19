# 사용자 정의 CLI의 커스텀 이미지

## CLI를 위한 사용자 정의 커스텀 이미지의 컨텍스트

Snyk이 [Snyk CLI 이미지](https://headwayapp.co/snyk-io-updates/deprecation-notice-for-snyk-cli-images-292562) 및 [Snyk 이미지](https://updates.snyk.io/deprecation-notice-for-obsolete-snyk-images-292563)에 관한 공지 후, Snyk은 고객이 자체 커스텀 이미지를 구축하는 데 유용한 지침을 제공합니다. 현재 지원되는 이미지 목록에 대해 [여기](https://github.com/snyk/snyk-images?tab=readme-ov-file#current-images)에서 Snyk 이미지 빌드 도구 체인을 방문할 수 있습니다.

{% hint style="info" %}
Snyk은 업스트림 공급 업체에 의해 더 이상 지원되지 않는 소프트웨어가 포함된 이미지를 작성하거나 유지하지 않습니다. 지원되지 않는 이미지 목록은 [GitHub repo](https://github.com/snyk/snyk-images?tab=readme-ov-file#vendor-unsupported-base-images)에서 확인할 수 있습니다.
{% endhint %}

커스텀 이미지 사용은 Snyk CLI가 지원하는 [환경](../../snyk-cli/install-or-update-the-snyk-cli/)를 확장할 수 있습니다.

## CLI를 위한 사용자 정의 커스텀 이미지 요구사항

사용자 정의 이미지가 지원되려면 다음이 필요합니다:

- Snyk CLI가 지원하는 [환경](../../snyk-cli/install-or-update-the-snyk-cli/) 사용
- Snyk에서 지원하는 언어 및 프레임워크 [사용](../../supported-languages-package-managers-and-frameworks/)
- 이미지에 Snyk CLI가 설치되어 있어야 함; CLI를 설치하는 방법은 [Snyk CLI 설치 또는 업데이트](../../snyk-cli/install-or-update-the-snyk-cli/)를 참조
- 공개적으로 접근 가능해야 함; 통합이 이미지를 가져올 것임

## CLI를 위한 사용자 정의 커스텀 이미지 사용

커스텀 이미지 제공은 환경에 대한 더 많은 통제를 제공합니다. 예를 들어, 커스텀 이미지를 사용하지 않으면 Node LTS를 사용하는 환경을 사용할 수 없습니다.

### 예시: Dockerfile을 사용하여 Node LTS 지원 커스텀 이미지 만들기

기본 요구 사항을 갖춘 상태에서 다음 Dockerfile을 사용하여 Node LTS를 지원하는 커스텀 이미지를 만들 수 있습니다:

{% code title="Dockerfile" %}
```docker
FROM alpine:3.18

# curl 설치
RUN apk add --no-cache curl

# Node LTS 설치
RUN apk add --no-cache nodejs

# Snyk CLI 설치 & 설정
RUN curl -o ./snyk-alpine https://downloads.snyk.io/cli/stable/snyk-alpine && \
    curl -o ./snyk-alpine.sha256 https://downloads.snyk.io/cli/stable/snyk-alpine.sha256 && \
    sha256sum -c snyk-alpine.sha256 && \
    mv snyk-alpine /usr/local/bin/snyk && \
    chmod +x /usr/local/bin/snyk
```
{% endcode %}

기본 이미지는 가벼운 Alpine을 사용합니다. 이미 Node 및 Snyk CLI를 설치했으므로 요구 사항의 75%를 충족합니다.

Dockerfile이 정의된 후 [docker build](https://docs.docker.com/engine/reference/commandline/build/)을 사용하여 이미지를 빌드하고 태그를 붙인 다음 [docker push](https://docs.docker.com/engine/reference/commandline/push/)를 사용하여 이미지를 푸시할 수 있습니다:

```sh
# 이미지 빌드
docker build <PATH-TO-DOCKERFILE> --tag foobar/snyk:node-lts

# 이미지 푸시
docker push foobar/snyk:node-lts
```

### 예시: BitBucket 파이프라인에서 커스텀 이미지 사용

BitBucket 파이프라인 통합의 호환성은 통합이 실행되는 Docker 컨테이너가 지원하는 환경으로 제한됩니다. Snyk 배포(Decoupling)에 대한 공지 [Snyk Scan을 Snyk CLI Docker 이미지에서 분리](https://updates.snyk.io/decoupling-snyk-scan-from-snyk-cli-docker-images-277502)를 따르면 v1.0.0 이전에는 환경이 Snyk CLI Docker 이미지에서 지원되는 것으로 제한되었습니다.

v1.0.0이 출시되면 사용자가 커스텀 이미지를 정의할 수 있습니다. [LANGUAGE](bitbucket-pipelines-integration-using-a-snyk-pipe/snyk-pipe-parameters-and-values-bitbucket-cloud.md#snyk-pipe-variables) 변수가 제공하는 환경 목록이 특정 빌드 환경을 지원하지 않는 경우, 커스텀 Docker 이미지 형태로 자체 빌드 환경을 정의할 수 있습니다.

[Bitbucket 파이프라인 통합 사전 준비 사항](bitbucket-pipelines-integration-using-a-snyk-pipe/prerequisites-for-bitbucket-pipelines-integration.md)이 충족되어 있는지 확인하십시오.

푸시된 이미지가 공개적으로 접근 가능하다면 Bitbucket 파이프라인에서 `SNYK_BASE_IMAGE` 및 `LANGUAGE` 변수를 사용하여 사용자 정의 이미지 및 태그를 참조할 수 있습니다:

{% code title="bitbucket-pipelines.yml" %}
```yaml
script:
  - npm install
  - npm test

  - pipe: snyk/snyk-scan:1.0.0
    variables:
      SNYK_TOKEN: $SNYK_TOKEN
      LANGUAGE: "node-lts"
      SNYK_BASE_IMAGE: "foobar/snyk"

# 나머지 스크립트
```
{% endcode %}