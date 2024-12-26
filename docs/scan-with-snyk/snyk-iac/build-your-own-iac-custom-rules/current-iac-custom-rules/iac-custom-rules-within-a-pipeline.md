# 파이프라인 내 IaC 사용자 정의 규칙

[GitHub Actions](https://github.com/features/actions)와 같은 CI/CD를 사용하면 사용자 정의 규칙을 관리하고 배포하며 강제하는 데 이상적입니다.

## 파이프라인에서 IaC 사용자 정의 규칙 개요

이 예시에서 보안 팀이 다음을 수행할 수 있습니다:

- GitHub 저장소에 규칙 저장
- 개발 시간 단계를 파이프라인에 추가하기 위해 GitHub Actions 사용
- 다른 GitHub 저장소를 구성하여 변경 사항을 검사하는 커스텀 규칙을 사용하는 GitHub Actions 파이프라인 실행

이 예시에서 Snyk은 [snyk/custom-rules-example](https://github.com/snyk/custom-rules-example) 저장소를 사용합니다. 이 저장소는 [SDK를 이용하여 규칙 작성 시작하기](writing-rules-using-the-sdk/) 중에 작성된 모든 사용자 정의 규칙을 포함합니다.

**목표:** 파이프라인을 다음과 같이 구성합니다:

- 새 규칙이나 기존 규칙에 대한 변경이 기존 기능을 망가뜨리지 않는지 확인
- `main`에 있는 규칙을 OCI 레지스트리에 게시
- 다른 파이프라인에서 사용자 정의 규칙 사용 강제
- 선택 사항: 환경 변수를 사용하여 사용자 정의 규칙 구성

## GitHub Action을 사용한 PR 체크 추가

PR 체크의 예시는 [https://github.com/snyk/custom-rules-example/pull/5](https://github.com/snyk/custom-rules-example/pull/5)에서 볼 수 있습니다. 여기서 `my_rule`이라는 새 규칙을 추가하려는 시도가 있습니다. 이는 [규칙 작성 방법을 학습하는 방법](writing-rules-using-the-sdk/writing-a-rule.md)에 나오는 동일한 규칙입니다.

이 규칙이 예상대로 작동하는지 확인하기 위해 단위 테스트를 구현했습니다. PR 체크의 일부로 단위 테스트를 실행하려면 이전에 `.github/workflows/test.yml`에 구성된 GitHub Action을 사용하십시오:

{% code title=".github/workflows/test.yml" %}
```
name: Test Custom Rules

on:
  push:
    branches:
      - '**'        # 모든 브랜치와 일치
      - '!main'     # main 제외

jobs:
  unit_test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules

      - name: Run unit tests
        run: snyk-iac-rules test
```
{% endcode %}

이 workflow에 대한 몇 가지 주목해야 할 사항:

- PR이 열릴 때 작동하도록 모든 비-`main` 브랜치에 대해 설정되었습니다.
- Node.js 환경을 설정하고 `npm`을 통해 `snyk-iac-rules` SDK를 설치할 수 있도록 단계가 추가되었습니다.
- `snyk-iac-rules test`를 실행하는 단계가 추가되었으며, 이는 테스트 중에 하나라도 실패하면 PR 체크가 실패합니다.

{% hint style="info" %}
사용자가 `main`에 직접 푸시할 수 없도록 먼저 `Settings` -> `Branches`에서 `main` 브랜치를 구성해야 합니다.
{% endhint %}

##  GitHub Action

규칙을 테스트하는 또 다른 방법은 [Snyk CLI](../../../../snyk-cli/)를 사용하여 컨트랙트를 테스트하는 것입니다. 이를 위해 [ GitHub Action](https://github.com/snyk/actions/tree/master/iac)을 사용하여 생성된 번들이 CLI에서 읽을 수 있는지 확인합니다.

이를 위해 Snyk CLI를 설치하는 단계와 `SNYK_TOKEN`이 필요합니다. 이는 Snyk 계정 설정에서 찾을 수 있습니다.

{% code title=".github/workflows/test.yml" %}
```
jobs:
  contract_test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules

      - name: Build bundle
        run: snyk-iac-rules build .

      - name: Run contract with Snyk to check Infrastructure as Code files for issues
        continue-on-error: true
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
        with:
          args: --rules=bundle.tar.gz
```
{% endcode %}

이러한 테스트를 [Shellspec](https://github.com/shellspec/shellspec)을 사용하여 확장하여 원하는 취약점이 트리거되는지 확인할 수 있지만, Snyk는 이를 위해 단위 테스트를 사용하는 것을 권장합니다.

## 사용자 정의 규칙 발행

이전 섹션의 모든 체크에서 PR이 통과하고 `main` 브랜치에 병합되면 Snyk 규칙을 OCI 레지스트리에 발행할 수 있습니다. 이를 통해 별도의 파이프라인을 구성하고 이 위치에서 사용자 정의 규칙 번들을 다운로드하여 사용자 오류 구성 사항을 찾을 수 있습니다.

이를 위해 `.github/workflows`에 `publish.yml`이라는 다른 workflow를 추가하십시오:

{% code title=".github/workflows/publish.yml" %}
```
name: Publish Custom Rules

on:
  push:
    branches:
      - 'main'

jobs:
  publish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - uses: actions/setup-node@v1
        with:
          node-version: 15

      - name: Install snyk-iac-rules
        run: npm i -g snyk-iac-rules
        
      - name: Build bundle
        run: snyk-iac-rules build .
        
      - name: Login to Docker Hub
        uses: docker/login-action@v1
        with:
          username: $
          password: $

      - name: Publish rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: "$:v1"
```
{% endcode %}

이전 workflow와 유사하지만 이 workflow에 대해 주목해야 할 몇 가지 사항이 있습니다:

- PR이 병합될 때만 `main` 브랜치에서만 실행되도록 구성되었습니다.
- Docker Hub로 인증하는 단계가 추가되었습니다. 지원되는 레지스트리 목록은 [번들을 푸시하기](writing-rules-using-the-sdk/pushing-a-bundle.md)에서 읽을 수 있습니다. 이를 위해 [docker/login-action](https://github.com/docker/login-action) GitHub Action을 사용하고 GitHub Secrets을 설정하십니다.
- 생성된 사용자 정의 규칙 번들을 OCI 레지스트리에 발행하는 `snyk-iac-rules build` 다음에 `snyk-iac-rules push`를 실행하는 단계가 추가되었습니다.

## 규칙 버전 관리

CI/CD 파이프라인 전체에 영향을 주지 않고 사용자 정의 규칙의 실험 버전을 출시하고 싶다면 번들에 버전을 지정하기 위해 태깅을 사용하세요.

`v1` 대부분의 서비스에서 계속 사용하는 동안 `v2-beta` 번들을 시도하려면 다음을 시작해보세요:

{% code title=".github/workflows/publish.yml" %}
```
      - name: Publish experimental rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: "$:v1"
      - name: Publish rules
        run: snyk-iac-rules push --registry $OCI_REGISTRY_URL bundle.tar.gz
        env:
          OCI_REGISTRY_URL: "$:v2-beta"
```
{% endcode %}

{% hint style="info" %}
이 workflow를 사용하려면 GitHub Secrets에 이미 OCI\_REGISTRY\_NAME에 태그 또는 프로토콜이 포함되어 있지 않아야 합니다.
{% endhint %}

## 사용자 정의 규칙 강제 실행

사용자 정의 규칙을 OCI 레지스트리에 발행한 후, 이러한 규칙을 사용하도록 별도의 파이프라인을 구성할 수 있습니다.

이를 위한 한 가지 방법은 API 엔드포인트 [그룹을 위한 Infrastructure as Code 설정 업데이트](https://apidocs.snyk.io/patch-/groups/-group\_id-/settings/#patch-/groups/-group\_id-/settings/iac)를 사용하는 것입니다.

이는 위의 GitHub Action을 Snyk가 구성된 사용자 정의 규칙 번들을 사용하도록 업데이트하기 위해 다른 작업으로 구성하는 것을 의미합니다:

```
      - name: Update Snyk
        run: |
          curl --location --request PATCH 'https://api.snyk.io/rest/groups/<group id>/settings/iac/?version=2021-11-03~beta' \
          --header 'Content-Type: application/vnd.api+json' \
          --header 'Authorization: token $' \
          --data-raw '{
            "data": {
                  "type": "iac_settings",
                  "attributes": {
                    "custom_rules": {
                      "oci_registry_url": "registry-1.$",
                      "oci_registry_tag": "v1",
                      "is_enabled": true
                    }
                }
            }
          }'
```

이 API 호출은 선택한 Snyk 그룹 및 해당 하위 조직이 구성된 사용자 정의 규칙 번들을 사용하도록 업데이트합니다.

{% hint style="info" %}
`v2-beta`와 같이 다른 번들을 사용하는 조직을 구성하려면 Snyk 설정 페이지를 사용하십시오. 새 번들을 구성하거나 사용자 정의 규칙을 사용하기 위해 CI/CD 파이프라인에서 환경 변수 사용을 허용하려면 사용자 정의 규칙을 비활성화하세요.
{% endhint %}

이 그룹 하위의 조직 중 하나와 함께 인증하고,  GitHub Action을 workflow에 추가합니다:

```
name: Snyk Infrastructure as Code Custom Rules

on:
  push:

jobs:
  snyk-iac-security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Run Snyk to check Infrastructure as Code files for issues
        continue-on-error: false
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
```

결과적으로 GitHub 액션이 생성된 구성 오류가 해결될 때까지 실패합니다:

```
Testing example.tf...


Infrastructure as code issues:
  ✗ IAM Role missing one of the required tags: owner, description or type [Medium Severity] [CUSTOM-RULE-8]
    introduced by input > resource > aws_iam_role[new_role] > tags

  ✗ Vendor or Service must have either owneralternate or ticket group or both tags. [Medium Severity] [CUSTOM-RULE-9]
    introduced by input > resource > aws_iam_role[new_role] > tags
```

## 사용자 정의 규칙 구성

API 또는 Snyk 설정 페이지를 사용하는 것이 제한적으로 보인다면, Snyk는 환경 변수를 사용하여 사용자 정의 규칙을 구성하는 방법도 제공합니다.

환경 변수인 `SNYK_CFG_OCI_REGISTRY_URL`, `SNYK_CFG_OCI_REGISTRY_USERNAME`, `SNYK_CFG_OCI_REGISTRY_PASSWORD`를 사용하여 설정 파일을 스캔하여 위반된 사용자 정의 규칙을 확인할 수 있습니다.

GitHub Action은 이러한 환경 변수를 읽고 이전 단계에서 발행된 번들을 구성된 OCI 레지스트리로 가져오게 됩니다. GitHub 액션이 다음과 같습니다:

```
name: Snyk Infrastructure as Code Custom Rules

on:
  push:

jobs:
  snyk-iac-security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Run Snyk to check Infrastructure as Code files for issues
        continue-on-error: false
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
          SNYK_CFG_OCI_REGISTRY_URL: $
          SNYK_CFG_OCI_REGISTRY_USERNAME: $
          SNYK_CFG_OCI_REGISTRY_PASSWORD: $
```