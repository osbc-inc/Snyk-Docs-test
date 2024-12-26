# Snyk PHP 액션

이 페이지는 [PHP](https://github.com/snyk/actions/tree/master/php)용 Snyk GitHub 액션을 사용하는 예제를 제공합니다. 액션 사용 방법 및 자세한 정보는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위해 Snyk PHP 액션 사용

다음과 같이 Snyk PHP 액션을 사용하여 취약점을 확인할 수 있습니다.

```yaml
name: Snyk 사용 예시 PHP 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/php@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 **높은 심각도 취약점**만 확인하려면 Snyk PHP 액션을 사용할 수 있습니다.

```yaml
name: Snyk 사용 예시 PHP 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/php@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk PHP 액션을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예시는 GitHub Actions 통합 페이지의 [Snyk monitor 예시](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 참조하십시오.

## Snyk PHP 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 작업이 실행되지 못하게 됩니다. 따라서 다음 예시에서와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) 옵션을 사용해야 합니다.

```yaml
name: Snyk 사용 예시 PHP 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/php@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위함
        env:
          SNYK_TOKEN: $
        with:
          args: --sarif-file-output=snyk.sarif
      - name: GitHub Code Scanning으로 결과 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에 대해 `upload-sarif` 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}