# Snyk Python-3.7 액션

{% hint style="warning" %}
이 이미지는 2024년 8월 12일에 제거되었습니다. 계속된 지원 및 최신 기능을 보장하려면 사용자들이 더 최신의 액션으로 이전하는 것이 권장됩니다. 현재 해당 이미지를 사용 중이라면, 이 날짜 이후의 작업 흐름에서의 장애를 피하기 위해 가능한 빨리 업그레이드 계획을 세워야 합니다.
{% endhint %}

이 페이지에서는 [Python (3.7)](https://github.com/snyk/actions/tree/master/python-3.7)용 Snyk GitHub Action을 사용하는 예제를 제공합니다. 액션 사용 및 추가 정보에 대한 지침은 [GitHub Actions 통합](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration)을 참조하십시오.

## Snyk Python (3.7) 액션

다음 예시에서는 Snyk Python GitHub Action을 사용하는 방법을 보여줍니다.

Snyk는 Python이 Snyk 검사를 실행하기 전에 종속성을 다운로드하는 것을 필요로 합니다.

Python 이미지는 현재 경로에서 매니페스트 파일이 존재하는 경우에만 종속성을 확인하고 설치합니다. 즉, 액션이 트리거되는 경로에서의 경로여야 합니다.

* 만약 pip가 현재 경로에 존재하고 Snyk이 `requirements.txt` 파일을 찾는다면, Snyk은 `pip install -r requirements.txt`를 실행합니다.
* 만약 pipenv가 현재 경로에 존재하고 Snyk이 `Pipfile.lock`가 없는 `Pipfile`을 찾는다면, Snyk은 `pipenv update`를 실행합니다.
* 만약 `pyproject.toml`이 현재 경로에 존재하고 Snyk이 `poetry.lock`을 찾지 못한다면, Snyk은 `pip install poetry`를 실행합니다.

만약 매니페스트 파일이 루트가 아닌 다른 위치에 존재한다면, **설치해야 합니다**.

다음과 같이 Snyk Python (3.7) 액션을 사용하여 취약점을 확인할 수 있습니다:

```yaml
name: Snyk을 사용한 Python-3.7 예시 작업흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약성 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.7@master
        env:
          SNYK_TOKEN: $
```

다음과 같이 Snyk Python (3.7) 액션을 사용하여 **고 심각도 취약점** 만 확인할 수 있습니다:

```yaml
name: Snyk을을 사용한 Python-3.7 예시 작업흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약성 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.7@master
        env:
          SNYK_TOKEN: $
        with:
          args: --severity-threshold=high
```

## Snyk Python 액션을 사용하여 snyk monitor 실행

`snyk monitor`를 실행하는 예시는 [Snyk 모니터 예시](https://docs.snyk.io/integrations/ci-cd-integrations/github-actions-integration#snyk-monitor-example)에서 GitHub Actions 통합 페이지를 참조하십시오.

## Snyk Python (3.7) 액션을 사용하여 Snyk 스캔 결과를 GitHub Code Scanning으로 업로드

`--sarif-file-output` [Snyk CLI 옵션](https://docs.snyk.io/snyk-cli/cli-reference) 및 [GitHub SARIF 업로드 액션](https://docs.github.com/en/code-security/secure-coding/uploading-a-sarif-file-to-github)을 사용하면 아래 예시에 표시된대로 Snyk 스캔 결과를 GitHub Code Scanning에 업로드할 수 있습니다.

Snyk 액션이 취약점을 발견하면 실패합니다. 이는 SARIF 업로드 액션이 실행되지 못하게 합니다. 따라서 다음 예시와 같이 [continue-on-error](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepscontinue-on-error) 옵션을 사용해야 합니다:

```yaml
name: Snyk을 사용한 Python-3.7 예시 작업흐름
on: push
jobs:
  security:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master
      - name: 취약성 확인을 위해 Snyk 실행
        uses: snyk/actions/python-3.7@master
        continue-on-error: true # SARIF 업로드가 호출되도록 하기 위해
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
개인 저장소의 경우 upload-sarif 옵션을 사용하려면 GitHub Advanced Security를 가져야 합니다.

`Advanced Security must be enabled for this repository to use code scanning` 오류가 표시되면 GitHub Advanced Security가 활성화되어 있는지 확인하십시오. 자세한 정보는 "[저장소의 보안 및 분석 설정 관리](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/enabling-features-for-your-repository/managing-security-and-analysis-settings-for-your-repository)"를 참조하십시오.
{% endhint %}
