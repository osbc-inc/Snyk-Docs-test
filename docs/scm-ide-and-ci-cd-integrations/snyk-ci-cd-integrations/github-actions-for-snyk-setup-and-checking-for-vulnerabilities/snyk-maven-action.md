# Snyk Maven 액션

이 페이지는 [Maven](https://github.com/snyk/actions/tree/master/maven)을 위한 Snyk GitHub 액션 사용 예제를 제공합니다. 액션 사용 방법 및 자세한 정보에 대한 지침은 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점을 확인하기 위한 Snyk Maven 액션 사용

다음과 같이 Snyk Maven 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk 사용 예제 Maven 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/maven@master
        env:
          SNYK_TOKEN: $
```

**고위험도 취약점만** 확인하기 위해 Snyk Maven 액션을 사용할 수 있습니다.

```yaml
name: Snyk 사용 예제 Maven 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/maven@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk Maven 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션은 취약점이 발견되면 실패합니다. 이로 인해 SARIF 업로드 액션이 실행되지 않도록 됩니다. 따라서 다음 예제처럼 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Snyk 사용 예제 Maven 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/maven@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위함
        env:
          SNYK_TOKEN: $
        with:
          args: --sarif-file-output=snyk.sarif
      - name: GitHub Code Scanning에 결과 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에 대해 upload-sarif 옵션을 사용하려면 GitHub 고급 보안이 필요합니다.

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}
