# Snyk 스칼라 동작

{% hint style="warning" %}
해당 이미지는 2024년 8월 12일에 제거되었습니다. 계속된 지원과 최신 기능을 보장하기 위해 사용자들이 새로운 동작으로 이전하는 것이 강력히 권장됩니다. 현재 해당 이미지를 사용 중이라면 이 날짜 이후의 작업 흐름에서 중단 사항이 발생하지 않도록 가능한 빨리 업그레이드를 계획하세요.
{% endhint %}

이 페이지는 [Scala](https://github.com/snyk/actions/tree/master/scala)를 위한 Snyk GitHub 동작을 사용하는 예제를 제공합니다. 동작 사용 및 추가 정보에 대한 지침은 [GitHub 동작 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점을 확인하기 위해 Snyk 스칼라 동작 사용

다음과 같이 Snyk 스칼라 동작을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk를 사용하는 스칼라 예제 작업 흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/scala@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 Snyk 스칼라 동작을 사용하여 **높은 심각도의 취약점만을 확인**할 수 있습니다:

```yaml
name: Snyk를 사용하는 스칼라 예제 작업 흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/scala@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk 스칼라 동작을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예제는 [Snyk monitor 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)에서 GitHub 동작 통합 페이지를 참조하십시오.

## Snyk 스칼라 동작을 사용하여 Snyk 스캔 결과를 GitHub 코드 스캔에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 동작](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub 코드 스캔에 업로드할 수 있습니다. 다음 예제에서 보여지는 대로입니다.

Snyk 동작은 취약점을 발견했을 때 실패합니다. 이는 SARIF 업로드 동작이 실행되지 못하도록합니다. 따라서 다음 예제에서와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다.

```yaml
name: Snyk를 사용하는 스칼라 예제 작업 흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/scala@master
        continue-on-error: true # SARIF 업로드가 호출되도록 합니다
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
개인 저장소에서 upload-sarif 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning`라는 오류가 표시되면 GitHub 고급 보안이 활성화되었는지 확인하십시오. 자세한 내용은 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}