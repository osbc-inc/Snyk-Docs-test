# Snyk Gradle-jdk11 작업

이 페이지는 [Gradle (jdk11)](https://github.com/snyk/actions/tree/master/gradle-jdk11)에 대한 Snyk GitHub 작업의 예제를 제공합니다. 작업을 사용하고 추가 정보를 보려면 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위한 Snyk Gradle (jdk11) 작업 사용

다음과 같이 Snyk Gradle(jdk11) 작업을 사용하여 취약점을 확인할 수 있습니다.

```yaml
name: Snyk를 사용한 Gradle (jdk11) 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/gradle-jdk11@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

다음과 같이 **높은 심각도 취약점에 대해서만** Snyk Gradle (jdk11) 작업을 사용할 수 있습니다.

```yaml
name: Snyk를 사용한 Gradle (jdk11) 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/gradle-jdk11@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## Snyk Gradle (jdk11) 작업을 사용하여 snyk 모니터 실행하기

`snyk monitor`를 실행하는 예제는 GitHub Actions 통합 페이지의 [Snyk 모니터 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 참조하십시오.

## Snyk Gradle (jdk11) 작업을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference)과 [GitHub SARIF 업로드 작업](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 작업은 취약점을 찾으면 실패합니다. 이로 인해 SARIF 업로드 작업이 실행되지 못할 수 있습니다. 따라서 이 예제에서보다시피 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) 옵션을 사용해야 합니다.

```yaml
name: Snyk를 사용한 Gradle (jdk11) 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/gradle-jdk11@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위해
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --sarif-file-output=snyk.sarif
      - name: GitHub Code Scanning에 결과 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
개인 저장소에서 upload-sarif 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 발생하는 경우 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 정보는 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}