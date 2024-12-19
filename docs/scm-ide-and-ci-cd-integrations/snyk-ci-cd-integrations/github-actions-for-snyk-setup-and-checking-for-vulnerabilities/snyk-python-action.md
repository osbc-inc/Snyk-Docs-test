# Snyk Python 액션

이 페이지에서는 [Python](https://github.com/snyk/actions/tree/master/python)을 위한 Snyk GitHub 액션을 사용하는 예제를 제공합니다. 액션 사용 방법 및 추가 정보에 대한 지침은 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점 확인을 위한 Snyk Python 액션 사용하기

다음 예시에서는 Snyk Python GitHub 액션을 사용하는 방법을 보여줍니다.

Snyk는 Python이 의존성을 다운로드한 후에 Snyk 체크를 실행하거나 트리거하기를 요구합니다.

Python 이미지는 현재 경로에 매니페스트 파일이 있는 경우에만 종속성을 확인하고 설치합니다. 즉, 액션이 트리거되는 경로에서의 경로입니다.

* 만약 현재 경로에 pip가 있고, Snyk가 `requirements.txt` 파일을 찾는다면, Snyk는 `pip install -r requirements.txt`를 실행합니다.
* 만약 현재 경로에 pipenv가 있고, Snyk가 `Pipfile.lock` 없이 `Pipfile`을 찾는다면, Snyk는 `pipenv update`를 실행합니다.
* 현재 경로에 `pyproject.toml`이 있고, Snyk가 `poetry.lock`을 찾지 못한다면, Snyk는 `pip install poetry`를 실행합니다.

만약 매니페스트 파일이 루트 외의 다른 위치에 있다면, **반드시 설치**되어야 합니다.

다음과 같이 Snyk Python 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Example workflow for Python using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

다음과 같이 Snyk CocoaPods 액션을 사용하여 **고 심각도 취약점만** 확인할 수 있습니다:

```yaml
name: Example workflow for Python using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## `Snyk monitor` 실행을 위한 Snyk Python 액션 사용

`snyk monitor`를 실행하는 예시에 대해서는 GitHub Actions 통합 페이지의 [Snyk monitor 예시](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 참조하십시오.

## Snyk Python 액션을 사용하여 Snyk 검사 결과를 GitHub Code Scanning에 업로드하기

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 액션이 실행되는 것을 막을 수 있습니다. 따라서 다음 예시처럼 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Example workflow for Python using Snyk
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: Run Snyk to check for vulnerabilities
        uses: snyk/actions/python@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위함
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --sarif-file-output=snyk.sarif
      - name: Upload result to GitHub Code Scanning
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에서 upload-sarif 옵션을 사용하려면 GitHub Advanced Security가 있어야합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning` 오류가 나타나면 GitHub Advanced Security가 활성화되어 있는지 확인하십시오. 자세한 정보는 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}