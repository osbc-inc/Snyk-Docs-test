# Snyk dotNET 작업

이 페이지는 [dotNET](https://github.com/snyk/actions/tree/master/dotnet)에 대한 Snyk GitHub 작업을 사용하는 예제를 제공합니다. 작업을 사용하는 방법 및 자세한 정보에 대한 지침은 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위한 Snyk dotNET 작업 사용

다음과 같이 Snyk dotNET 작업을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk 사용하는 dotNET 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: .NET 설정
        uses: actions/setup-dotnet@3.0.3
      - name: 종속성 복원
        run: dotnet restore ./path/to/your.sln
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/dotnet@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 **높은 심각도의 취약점**만 확인하려면 Snyk dotNET 작업을 사용할 수 있습니다:

```yaml
name: Snyk 사용하는 dotNET 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
      - name: .NET 설정
        uses: actions/setup-dotnet@3.0.3
      - name: 종속성 복원
        run: dotnet restore ./path/to/your.sln
        uses: snyk/actions/dotnet@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

{% hint style="info" %}
Snyk 작업을 실행하기 전에 `dotnet restore` 또는 `nuget restore`를 사용하여 종속성을 복원해야 합니다.
{% endhint %}

## Snyk dotNET 작업을 사용하여 snyk monitor 실행

`snyk monitor`를 실행한 예제에 대해서는 [Snyk 모니터 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 GitHub Actions 통합 페이지에서 확인하십시오.

## Snyk dotNET 작업을 사용하여 Snyk 검사 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference)과 [GitHub SARIF 업로드 작업](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 검사 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 작업은 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 작업이 실행되지 못하게 합니다. 따라서 이 예제에서 보이는 것과 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다.

```yaml
name: Snyk 사용하는 dotNET 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: .NET 설정
        uses: actions/setup-dotnet@3.0.3
      - name: 종속성 복원
        run: dotnet restore ./path/to/your.sln
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/dotnet@master
        continue-on-error: true # SARIF 업로드 호출을 위해
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
사용자가 비공개 저장소의 경우 sarif 업로드 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 정보는 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}