# Snyk Python-3.6 Action

{% hint style="warning" %}
본 이미지는 2024년 8월 12일에 제거되었습니다. 권장사항으로 사용자들은 지속적인 지원과 최신 기능을 보장하기 위해 최신 액션으로의 이전을 고려하는 것이 매우 권장됩니다. 현재 본 이미지를 사용 중이라면 이 날짜 이후의 업무 흐름 중단을 피하기 위해 가능한 빨리 업그레이드 계획을 세우십시오.
{% endhint %}

이 페이지는 [Python (3.6)](https://github.com/snyk/actions/tree/master/python-3.6)을 위한 Snyk GitHub 액션 사용 예제를 제공합니다. 작업 방법 및 자세한 정보에 대한 지침은 [GitHub 액션 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## 취약점을 확인하기 위한 Snyk Python (3.6) 액션 사용

다음 예제에서는 Snyk Python GitHub 액션을 사용하는 방법을 보여줍니다.

Snyk는 Python이 종속성을 다운로드한 후에 Snyk 체크를 실행하거나 트리거해야 한다고 요구합니다.

Python 이미지는 현재 경로에 매니페스트 파일이 있는 경우에만 종속성을 확인하고 설치합니다. 즉, 작업이 트리거되는 경로에서의 경로입니다.

* 현재 경로에 pip가 존재하고 Snyk가 `requirements.txt` 파일을 찾으면 Snyk가 `pip install -r requirements.txt`을 실행합니다.
* 현재 경로에 pipenv가 존재하고 Snyk가 `Pipfile.lock`이 없는 `Pipfile`을 찾으면 Snyk가 `pipenv update`를 실행합니다.
* 현재 경로에 `pyproject.toml`이 존재하고 Snyk가 `poetry.lock`을 찾지 못하면 Snyk가 `pip install poetry`를 실행합니다.

루트 이외의 위치에 매니페스트 파일이 있는 경우 **반드시 설치**되어야 합니다.

다음과 같이 Snyk Python (3.6) 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk를 사용한 Python-3.6 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.6@master
        env:
          SNYK_TOKEN: $
```

**높은 심각도의 취약점만** 확인하도록 Snyk Python (3.6) 액션을 사용할 수 있습니다:

```yaml
name: Snyk를 사용한 Python-3.6 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.6@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk Python (3.6) 액션을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예제는 [Snyk 모니터 예제](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)를 GitHub 액션 통합 페이지에서 확인하십시오.

## Snyk Python (3.6) 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이로 인해 SARIF 업로드 액션이 실행되지 않습니다. 따라서 다음 예제와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Snyk를 사용한 Python-3.6 예제 워크플로우
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약점 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.6@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위해
        env:
          SNYK_TOKEN: $
        with:
          args: --sarif-file-output=snyk.sarif
      - name: 결과를 GitHub Code Scanning에 업로드
        uses: github/codeql-action/upload-sarif@v2
        with:
          sarif_file: snyk.sarif
```

{% hint style="info" %}
비공개 저장소에 sarif 업로드 옵션을 사용하려면 GitHub 고급 보안이 있어야 합니다.

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub 고급 보안이 활성화되어 있는지 확인하십시오. 자세한 내용은 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"을 참조하십시오.
{% endhint %}
