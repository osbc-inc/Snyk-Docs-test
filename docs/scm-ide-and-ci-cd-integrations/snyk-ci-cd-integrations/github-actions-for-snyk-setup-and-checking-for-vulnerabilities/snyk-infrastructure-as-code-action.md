# Snyk Infrastructure as Code 동작

이 페이지는 [Infrastructure as Code](https://github.com/snyk/actions/tree/master/iac)용 Snyk GitHub 동작의 사용 방법 및 예제를 제공합니다. 일반적인 지침과 정보는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

Snyk Infrastructure as Code 테스트 동작을 사용하려면 Snyk API 토큰이 있어야 합니다. [Snyk 토큰 가져오기](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#getting-your-snyk-token)을 참조하거나 [무료로 가입](https://snyk.io/login)할 수 있습니다.

## 취약점 확인을 위한 Snyk Infrastructure as Code 동작 사용

다음과 같이 Snyk Infrastructure as Code 동작을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk Infrastructure as Code 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 Kubernetes manifest 파일을 확인하는 동작 실행
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
```

Snyk Infrastructure as Code 동작 속성

Snyk Infrastructure as Code 동작은 기본 이미지로 전달되는 속성을 가지고 있습니다. 이러한 속성은 `with`를 사용하여 동작으로 전달됩니다:

| 속성     | 기본값   | 설명                                                              |
| --------- | -------- | ----------------------------------------------------------------- |
| `args`    |          | Snyk 이미지에 대한 기본 인수 재정의                             |
| `command` | `"test"` | 실행할 명령을 지정합니다. 현재 `test`만 지원됩니다.                 |
| `file`    |          | 문제가 있는 파일을 스캔할 경로입니다.                                |
| `json`    | `false`  | 표준 출력 외에도 결과를 snyk.json으로 저장                           |
| `sarif`   | `true`   | 표준 출력 외에도 결과를 snyk.sarif으로 저장                         |

## Snyk Infrastructure as Code 동작 예시 

### 경로 지정

테스트 중에 대상으로 지정할 구성 파일 및 디렉토리의 경로를 지정할 수 있습니다.\
경로가 지정되지 않으면 기본적으로 전체 저장소가 스캔됩니다.

```yaml
name: Snyk Infrastructure as Code 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 Kubernetes manifest 파일을 확인하는 동작 실행
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
        with:
          file: your/kubernetes-manifest.yaml your/terraform/directory
```

### 심각도 임계값 지정

높은 심각도 취약점에 대해서만 보고할 수도 있습니다.

```yaml
name: Snyk Infrastructure as Code 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 Kubernetes manifest 파일을 확인하는 동작 실행
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
        with:
          file: your/kubernetes-manifest.yaml
          args: --severity-threshold=high
```

### 테스트 결과 공유

테스트 결과를 [Snyk 플랫폼](https://docs.snyk.io/products/snyk-infrastructure-as-code/share-cli-results-with-the-snyk-web-ui)과 공유할 수 있습니다.

```yaml
name: Snyk Infrastructure as Code 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 Kubernetes manifest 파일을 확인하는 동작 실행
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
        with:
          args: --report
```

### Terraform Plan의 스캔 모드 지정

Terraform Plan 파일을 스캔할 때 [스캔 모드](https://docs.snyk.io/products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code/test-your-terraform-files-with-the-cli-tool#terraform-plan)도 선택할 수 있습니다.

```yaml
name: Snyk Infrastructure as Code 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 Kubernetes manifest 파일을 확인하는 동작 실행
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
        with:
          args: --scan=resource-changes
```

## Snyk Infrastructure as Code 동작을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드하기

Infrastructure as Code 동작은 GitHub Code Scanning과 통합할 수 있으며 GitHub 보안 탭에서 문제를 표시할 수 있습니다. 동작을 실행하면 업로드할 수 있는 `snyk.sarif` 파일이 생성됩니다.

```yaml
name: Snyk Infrastructure as Code
on: push
jobs:
  snyk:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Snyk를 사용하여 구성 파일을 보안 문제 스캔
        # Snyk가 보안 문제를 감지하면 빌드를 중단하는 데 사용할 수 있습니다.
        # 이 경우 문제를 GitHub Code Scanning에 업로드하려고 합니다
        continue-on-error: true
        uses: snyk/actions/iac@master
        env:
          SNYK_TOKEN: $
      - name: GitHub Code Scanning에 결과 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에 대한 업로드-sarif 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[고유 저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}

## 관련 문서

`snyk iac test` 명령어에 대한 자세한 내용은 다음을 참조하십시오:

* [Snyk CLI for Infastructure as Code](https://docs.snyk.io/products/snyk-infrastructure-as-code/snyk-cli-for-infrastructure-as-code)
* [Snyk Infrastructure as Code Test Command](https://docs.snyk.io/snyk-cli/commands/iac-test)