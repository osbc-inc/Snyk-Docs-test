# BitBucket Pipelines 이전

## `snyk/snyk-scan` < v1.0.0 <a href="#users-using-snyk-snyk-scan-less-than-v1.0.0" id="users-using-snyk-snyk-scan-less-than-v1.0.0"></a>를 사용하는 사용자들을 위해

`snyk/snyk-scan` \<v1.0.0는 Snyk CLI 이미지를 사용합니다. 모든 Snyk CLI 이미지가 제거될 예정이므로, `snyk/snyk-scan` < v1.0.0도 작동을 멈출 것입니다.

`[Snyk 문서의 업그레이드 가이드](../../bitbucket-pipelines-integration-using-a-snyk-pipe/migrating-to-bitbucket-pipelines-v1.0.0.md)`에서 `snyk/snyk-scan` >= v1.0.0로 업그레이드하는 방법을 참조하세요.

## `snyk/snyk-scan` >= v1.0.0 <a href="#users-using-snyk-snyk-scan-greater-than-v1.0.0" id="users-using-snyk-snyk-scan-greater-than-v1.0.0"></a>를 사용하는 사용자들을 위해

### 사용자 정의 이미지 만들기 <a href="#create-your-own-custom-image" id="create-your-own-custom-image"></a>

사용자들은 자체 사용자 지정 이미지를 만들어 사용할 수 있습니다. 이 옵션은 `snyk/snyk-scan` >= v1.0.0에서만 사용할 수 있습니다. 자세한 내용은 [CLI를 위한 사용자 정의 이미지](../../user-defined-custom-images-for-cli.md)를 참조하세요.

사용자 정의 이미지를 만드는 것은 시스템과의 호환성을 보장해야 합니다. 그러나 사용자 정의 이미지를 만들 수 없는 경우 업그레이드할 수 있는 대체 이미지가 있습니다.

### 지원되는 Snyk 이미지로 업그레이드 <a href="#upgrade-to-a-supported-snyk-image" id="upgrade-to-a-supported-snyk-image"></a>

사용 중인 Snyk 이미지가 삭제될 것임을 확인한 후, [여기](bitbucket-pipelines-migration.md#users-using-snyk-snyk-scan-less-than-v1.0.0)에 설명된 것처럼, [Snyk 이미지 마이그레이션](snyk-images-migration.md) 가이드를 참조하여 구성에 대한 업그레이드 경로를 확인하세요.

{% hint style="info" %}
더 나은 안정성을 위해 가능한 경우 고정 버전을 사용하세요. 예를 들어 `snyk/snyk:dotnet-8.0`는 `snyk/snyk:dotnet` 보다 우선합니다.
{% endhint %}

**지원되는 Snyk 이미지로 업그레이드하는 예시**가 있습니다.

다음 `bitbucket-pipeline.yml` 구성 예시에서 2024년 8월 12일에 삭제될 예정인 Snyk 이미지가 설정되어 있습니다:

```yaml
#  `snyk/snyk:node-16` Snyk 이미지를 사용하는 예시 bitbucket-pipelines.yml
#  NodeJS 빌드용 템플릿

#  이 템플릿을 사용하면 NodeJS 코드를 검증할 수 있습니다.
#  Workflow를 사용하여 테스트 및 코드 린팅을 기본 브랜치에서 실행할 수 있습니다.

image: atlassian/default-image:latest

pipelines:
  default:
    - parallel:
        - step:
            name: 빌드
            caches:
              - node
            script:
              - npm install
        - step:
            name: Snyk 스캔
            script:
              - pipe: snyk/snyk-scan:1.0.1
                variables:
                  SNYK_TOKEN: $SNYK_TOKEN
                  LANGUAGE: "node-16" # <------ `snyk/snyk:node-16` Snyk 이미지 사용
                  EXTRA_ARGS: "--all-projects" # 선택 사항
                  DEBUG: "true" # 선택 사항
```

[Snyk 이미지 마이그레이션](snyk-images-migration.md) 가이드를 따라, 다음과 같이 지원되는 Snyk 이미지로 업그레이드할 수 있습니다:

```yaml
#  지원되는 Snyk 이미지 `snyk/snyk:node-22`로 업그레이드
#  NodeJS 빌드용 템플릿

#  이 템플릿을 사용하면 NodeJS 코드를 검증할 수 있습니다.
#  Workflow를 사용하여 테스트 및 코드 린팅을 기본 브랜치에서 실행할 수 있습니다.

image: atlassian/default-image:latest

pipelines:
  default:
    - parallel:
        - step:
            name: 빌드
            caches:
              - node
            script:
              - npm install
        - step:
            name: Snyk 스캔
            script:
              - pipe: snyk/snyk-scan:1.0.1
                variables:
                  SNYK_TOKEN: $SNYK_TOKEN
                  LANGUAGE: "node-22" # <------ `snyk/snyk:node-22` Snyk 이미지로 업그레이드
                  EXTRA_ARGS: "--all-projects" # 선택 사항
                  DEBUG: "true" # 선택 사항
```

## Snyk CLI 직접 다운로드 및 설치 <a href="#download-and-install-snyk-cli-directly" id="download-and-install-snyk-cli-directly"></a>

Bitbucket의 `snyk/snyk-scan` 통합을 사용하고 싶지 않은 경우, 직접 Snyk CLI를 설치하여 사용할 수 있습니다.

{% hint style="info" %}
이 옵션을 사용하면 코드 인사이트 결과와 같은 통합 기능을 사용할 수 없습니다.
{% endhint %}

다음은 **CLI 직접 사용 예시**입니다.

다음 `bitbucket-pipeline.yml` 구성 예시에서 CLI를 다운로드하고:

- CLI를 SHASUM으로 확인합니다.
- 코드를 테스트하기 위해 CLI를 실행합니다.

```yaml
image: node:18

pipelines:
  default:
    - step:
        name: 빌드
        caches:
          - node
        script:
          - npm install
    - step:
        name: Snyk 스캔
        script:
          # Snyk Linux CLI 다운로드
          - curl https://downloads.snyk.io/cli/latest/snyk-linux -o snyk-linux
          # Snyk Linux CLI SHASUM 다운로드
          - curl https://downloads.snyk.io/cli/latest/snyk-linux.sha256 -o snyk.sha256
          # SHASUM을 사용하여 바이너리 유효성 검사
          - sha256sum -c snyk.sha256
          # 실행을 위해 CLI 구성
          - chmod +x snyk-linux
          - mv snyk-linux /usr/local/bin/snyk
          # Snyk CLI 실행
          - snyk test --all-projects -d
```

&#x20;

\
