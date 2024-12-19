# Snyk Python-3.8 액션

이 페이지는 [Python (3.8)](https://github.com/snyk/actions/tree/master/python-3.8)을 위한 Snyk GitHub 액션을 사용하는 예제를 제공합니다. 액션 사용 방법 및 자세한 정보는 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점을 확인하기 위한 Snyk Python (3.8) 액션 사용

다음 예제는 Snyk Python GitHub 액션을 어떻게 사용할 수 있는지 보여줍니다.

Snyk는 Python이 Snyk 체크를 실행하기 전에 종속성을 다운로드하도록 요구합니다.

Python 이미지는 현재 경로에 manifest 파일이 있는 경우에만 종속성을 확인하고 설치합니다. 즉, 액션이 트리거되는 경로에서 manifest 파일이 있어야 합니다.

- 현재 경로에 pip가 존재하고 Snyk가 `requirements.txt` 파일을 찾으면, Snyk는 `pip install -r requirements.txt`을 실행합니다.
- 현재 경로에 pipenv가 있고 Snyk가 `Pipfile.lock`이 없는 `Pipfile`을 찾으면, Snyk는 `pipenv update`를 실행합니다.
- 현재 경로에 `pyproject.toml`이 있고 Snyk가 `poetry.lock`을 찾지 못하면, Snyk는 `pip install poetry`를 실행합니다.

루트 외의 위치에 manifest 파일이 있으면 **반드시 설치**해야 합니다.

다음과 같이 Snyk Python 액션 (3.8)을 사용하여 취약점을 확인할 수 있습니다.

```yaml
name: Python-3.8를 위한 Snyk 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점을 확인하기 위해 Snyk 실행
        uses: snyk/actions/python-3.8@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
```

다음과 같이 Snyk Python (3.8) 액션을 사용하여 **높은 심각도의 취약점만** 확인할 수 있습니다.

```yaml
name: Python-3.8를 위한 Snyk 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점을 확인하기 위해 Snyk 실행
        uses: snyk/actions/python-3.8@master
        env:
          SNYK_TOKEN: ${{ secrets.SNYK_TOKEN }}
        with:
          args: --severity-threshold=high
```

## Snyk Python (3.8) 액션을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예제에 대해서는 GitHub Actions 통합 페이지의 [Snyk monitor 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 참조하십시오.

## Snyk Python (3.8) 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드하기

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 다음 예제와 같이 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션은 취약점을 발견하면 실패합니다. 이로 인해 SARIF 업로드 액션이 실행되지 않습니다. 따라서 다음 예제와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob\_idstepscontinue-on-error) 옵션을 사용해야 합니다.

```yaml
name: Python-3.8를 위한 Snyk 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점을 확인하기 위해 Snyk 실행
        uses: snyk/actions/python-3.8@master
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
사설 저장소에 대한 `upload-sarif` 옵션을 사용하려면 GitHub 고급 보안이 필요합니다. &#x20;

`Advanced Security must be enabled for this repository to use code scanning`라는 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[Managing security and analysis settings for your repository](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}