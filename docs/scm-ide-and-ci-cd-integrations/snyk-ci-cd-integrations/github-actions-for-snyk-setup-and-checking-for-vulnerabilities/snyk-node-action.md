# Snyk 노드 액션

이 페이지는 [Node](https://github.com/snyk/actions/tree/master/node)를 위한 Snyk GitHub 액션 사용 예시를 제공합니다. 액션 사용 방법 및 자세한 정보는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위한 Snyk 노드 액션 사용

다음처럼 Snyk 노드 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk를 사용하는 Node 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

다음과 같이 **높은 심각도의 취약점만 확인**하려면 Snyk 노드 액션을 사용할 수 있습니다:

```yaml
name: Snyk를 사용하는 Node 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/node@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## snyk monitor 실행을 위한 Snyk 노드 액션 사용

`snyk monitor`를 실행하는 예시는 [Snyk monitor 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 GitHub Actions 통합 페이지에서 참조하십시오.

## Snyk 노드 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 다음 예제에 표시된대로 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션은 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 액션이 실행되지 못하게 합니다. 따라서 다음 예제에 표시된대로 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Snyk를 사용하는 Node 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/node@master
        continue-on-error: true # SARIF 업로드가 호출되도록 함
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
사적 리포지토리에 대해 sarif 업로드 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}