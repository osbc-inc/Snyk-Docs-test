# Snyk Gradle-jdk12 액션

이 페이지는 [Gradle (jdk12)](https://github.com/snyk/actions/tree/master/gradle-jdk12)을 위한 Snyk GitHub Action 사용 예시를 제공합니다. 실행 방법과 추가 정보에 대한 안내는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하세요.

## 취약점을 확인하기 위한 Snyk Gradle (jdk12) 액션 사용

다음과 같이 Snyk Gradle (jdk12) 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Example workflow for Gradle (jdk12) using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/gradle-jdk12@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 **높은 심각도의 취약점만** 확인하려면 Snyk Gradle (jdk12) 액션을 사용할 수 있습니다:

```yaml
name: Example workflow for Gradle (jdk12) using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/gradle-jdk12@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk Gradle (jdk12) 액션을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예제는 GitHub Actions 통합 페이지의 [Snyk monitor 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 참조하세요.

## Snyk Gradle (jdk12) 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 액션이 실행되지 못하게 합니다. 따라서 아래 예시와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Example workflow for Gradle (jdk12) using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/gradle-jdk12@master
        continue-on-error: true # SARIF 업로드가 호출되도록 함
        env:
          SNYK_TOKEN: $
        with:
          args: --sarif-file-output=snyk.sarif
      - name: Upload result to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에 업로드하기 위해서는 GitHub Advanced Security가 필요합니다.

`Advanced Security must be enabled for this repository to use code scanning` 오류가 발생하면 GitHub Advanced Security가 활성화되어 있는지 확인하세요. 자세한 정보는 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하세요.
{% endhint %}
