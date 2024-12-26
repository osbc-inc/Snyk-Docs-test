# Snyk 도커 액션

이 페이지는 [Docker](https://github.com/snyk/actions/tree/master/docker)에 대한 Snyk GitHub 액션을 사용하는 방법과 예제에 대한 지침을 제공합니다. 일반적인 지침 및 정보는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

도커 액션을 사용하려면 Snyk API 토큰이 필요합니다. [Snyk 토큰 가져오기](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#getting-your-snyk-token)을 참조하거나 무료로 [가입할 수 있습니다](https://snyk.io/login).

## 취약점 확인을 위해 Snyk 도커 액션 사용

다음과 같이 Snyk 도커 액션을 사용하여 Docker 이미지에서 취약점을 확인할 수 있습니다:

```yaml
name: Snyk를 사용한 Docker 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - name: Snyk를 실행하여 Docker 이미지의 취약점 확인
      uses: snyk/actions/docker@master
      env:
        SNYK_TOKEN: $
      with:
        image: your/image-to-test
```

## Snyk 도커 액션 속성

Snyk 도커 액션에는 기본 이미지로 전달되는 속성이 있습니다. 이러한 것들은 `with`를 사용하여 액션에 전달됩니다.

| 속성     | 기본값 | 설명                                            |
| -------- | ------- | --------------------------------------------- |
| args     |         | Snyk 이미지에 대한 기본 인수 재정의              |
| command  | test    | 예를 들어 test 또는 monitor와 같은 명령 지정     |
| image    |         | 테스트할 이미지의 이름                            |
| json     | false   | stdout 외에도 결과를 snyk.json으로 저장           |
| sarif    | true    | stdout 외에도 결과를 snyk.sarif로 저장           |

예를 들어 다음과 같이 Snyk 도커 액션을 사용하여 **고심각도의 취약점만 확인**할 수 있습니다:

```yaml
name: Snyk를 사용한 Docker 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
    - name: Snyk를 실행하여 Docker 이미지의 취약점 확인
      uses: snyk/actions/docker@master
      env:
        SNYK_TOKEN: $
      with:
        image: your/image-to-test
        args: --severity-threshold=high
```

## Snyk 도커 액션을 사용하여 GitHub 코드 스캔에 스캔 결과 업로드

도커 액션은 GitHub 코드 스캔과 통합을 지원하며 GitHub 보안 탭에서 문제를 표시할 수 있습니다. `args`에 `--file=Dockerfile`을 참조하면 `snyk.sarif` 파일이 생성되어 GitHub 코드 스캔에 업로드할 수 있습니다.

```yaml
name: 
on: push
jobs:
  snyk:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Docker 이미지 빌드
      run: docker build -t your/image-to-test .
    - name: Snyk를 실행하여 Docker 이미지의 취약점 확인
      continue-on-error: true
      uses: snyk/actions/docker@master
      env:
        SNYK_TOKEN: $
      with:
        image: your/image-to-test
        args: --file=Dockerfile
    - name: GitHub 코드 스캔에 결과 업로드
      uses: github/codeql-action/upload-sarif@v2
      with:
        sarif_file: snyk.sarif
```

{% hint style="info" %}
개인 저장소에 대한 업로드-sarif 옵션 사용을 위해 GitHub Advanced Security가 필요합니다. &#x20;

"고급 보안을 사용하도록 설정해야 함"이라는 오류가 표시되면 GitHub Advanced Security가 활성화되어 있는지 확인하십시오. 자세한 정보는 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}