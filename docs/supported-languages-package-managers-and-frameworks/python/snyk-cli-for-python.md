# Snyk Python CLI

## CLI에서 Python 버전 설정

CLI에서 Python 버전을 설정하려면 `snyk test` 또는 `snyk monitor`에 다음 옵션을 추가하고 Python 이진 파일의 이름을 지정하십시오:

```sh
--command=python3
```

자세한 내용은 [`snyk test`](../../snyk-cli/commands/test.md)와 [`snyk monitor`](../../snyk-cli/commands/monitor.md) 도움말에서 Python 프로젝트 옵션을 확인하세요.

## Pip 및 CLI

{% hint style="info" %}
CLI를 사용하기 전에 `pip install`을 실행하고 스캔을 진행하세요. 아래는 예시입니다:

```
pip install -r requirements.txt
```
{% endhint %}

Pip `requirements.txt` 파일은 최상위 종속성만을 지정하며 중첩된 또는 전이적인 종속성을 지정하지 않습니다. 따라서 CLI가 완전한 종속성 트리를 구축할 수 있도록 전체 Pip 프로젝트가 설치되어야 합니다.

## Poetry 및 CLI

Poetry 애플리케이션의 종속성 트리를 빌드하기 위해 Snyk은 `pyproject.toml` 및 `poetry.lock` 파일을 사용합니다. Snyk가 Poetry 종속성을 스캔하고 문제를 식별하려면 두 파일이 모두 있어야 합니다.

`poetry.lock` 파일이 없는 경우 스캔하기 전에 `poetry lock`을 실행하여 생성해야 합니다.

{% hint style="info" %}
[PEP 621](https://peps.python.org/pep-0621/)은 `pyproject.toml` 파일에서 직접 종속성을 정의하는 표준이며, 이는 Poetry가 하는 방식과 다릅니다.&#x20;

Snyk는 PEP 621을 지원하지 않습니다.
{% endhint %}

## Pipenv 및 CLI

Pipenv 애플리케이션의 종속성 트리를 빌드하기 위해 Snyk은 `Pipfile` 및 `Pipfile.lock` 파일을 사용합니다. Snyk가 Pipenv 종속성을 스캔하고 문제를 식별하려면 두 파일이 모두 있어야 합니다.

CLI를 사용하기 전에 `pip install`을 실행합니다.

`pipenv graph`를 사용하여 최신 및 정확한 종속성 트리를 빌드할 수 있도록 `pipenv install`을 실행합니다.

## setup.py 및 CLI

종속성 트리를 빌드하기 위해 Snyk는 `setup.py` 파일을 분석하고 `install_requires` 키에 나열된 패키지를 감지합니다.

이 파일은 CLI에 자동으로 발견되지 않습니다. 수동으로 `--file` 옵션을 사용하여 명시해야 하며 다음과 같이 지정할 수 있습니다:

```python
snyk test --file=setup.py
```

또한 `setup.py`를 가상 환경에 패키지를 설치한 후 `pip freeze`를 실행하여 `requirements.txt`로 변환할 수 있습니다.