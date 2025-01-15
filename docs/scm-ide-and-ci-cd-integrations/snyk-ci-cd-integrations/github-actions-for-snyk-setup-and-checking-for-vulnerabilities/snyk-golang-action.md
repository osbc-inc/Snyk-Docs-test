# Snyk Golang 액션

이 페이지는 [Golaong](https://github.com/snyk/actions/tree/master/golang)을 위한 Snyk GitHub 액션 사용 예제를 제공합니다. 액션 사용 방법 및 자세한 정보에 대한 지침은 [GitHub 액션 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위해 Snyk Golang 액션 사용하기

다음과 같이 Snyk Golang 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk 사용 예시 Golang 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/golang@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 **높은 심각도의 취약점** 만 확인하도록 Snyk Golang 액션을 사용할 수 있습니다:

```yaml
name: Snyk 사용 예시 Golang 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/golang@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk Golang액션을 사용하여 snyk monitor 실행하기

`snyk monitor` 실행 예시는 [Snyk 모니터 예시](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)에서 GitHub 액션 통합 페이지를 확인하십시오.

## Snyk Golang 액션을 사용하여 Snyk 스캔 결과를 GitHub 코드 스캔에 업로드하기

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub 코드 스캔에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 액션 실행을 방해할 수 있습니다. 따라서 다음 예시에 표시된대로 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Snyk 사용 예시 Golang 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/golang@master
        continue-on-error: true # SARIF 업로드가 호출되도록 함
        env:
          SNYK_TOKEN: $
        with:
          args: --sarif-file-output=snyk.sarif
      - name: GitHub 코드 스캔에 결과 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
개인 저장소에 upload-sarif 옵션을 사용하려면 GitHub 고급 보안이 필요합니다.

`Advanced Security must be enabled for this repository to use code scanning` 오류가 발생하는 경우, GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 정보는 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}
