# Python

## 적용 가능성

Snyk는 [코드 분석용 Python](python-for-code-analysis.md) 및 [오픈 소스용 Python](python-for-open-source.md)를 지원합니다.

버전 및 패키지 관리자 사용에 대한 구체적인 정보는 [Python을 위한 Snyk CLI](snyk-cli-for-python.md) 및 [Git 저장소 및 Python](git-repositories-and-python.md)를 참조하십시오.

언어 사용 가능성을 확인하여 Snyk 제품을 사용하여 애플리케이션으로 가져오거나 테스트하거나 모니터링합니다.

사용 가능한 기능:

* SCM 가져오기, Snyk Open Source 및 Snyk Code용 사용 가능. Snyk Open Source와 함께 사용되는 Python의 경우, Pip, pipenv, Poetry를 지원합니다.
* CLI 및 IDE를 통한 앱 테스트 또는 모니터링, Snyk Open Source 및 Snyk Code용 사용 가능.
* pkg:pypi를 사용하여 앱의 SBOM 테스트
* pkg:pypi를 사용하여 앱의 패키지 테스트

## 패키지 관리자 및 지원하는 파일 확장자

Snyk은 Python에 대해 Pip, Poetry, pipenv 및 setup.py를 지원합니다. 지원되는 Python 버전 목록은 [Git 저장소 및 Python](git-repositories-and-python.md) 페이지를 확인하십시오.

패키지 레지스트리로 [pypi.org](https://pypi.org/)가 지원됩니다.

Python에 대한 Snyk은 다음 파일 형식을 지원합니다:

* Snyk Open Source:
  * Poetry를 위한: `pyproject.toml`, `poetry.lock`
  * Pip를 위한: `requirements.txt`
  * Pipenv를 위한: `pipfile`, `pipfile.lock`
  * setup.py를 위한: `setup.py`
* Snyk Code: `.py`

## 프레임워크 및 라이브러리

다음 프레임워크 및 라이브러리가 Snyk에서 Python을 지원합니다:

* AioHTTP - 포괄적
* iopg - 포괄적
* argparse - 포괄적
* anthropic - 포괄적
* bottle - 포괄적
* CherryPy - 포괄적
* Django - 포괄적
* defusedxml - 포괄적
* fastapi - 부분적
* flask - 포괄적
* flask\_pymongo - 포괄적
* google.cloud.bigquery - 포괄적
* google\_generativeai - 포괄적
* huggingface\_hub - 포괄적
* httpx - 포괄적
* ldap3 - 포괄적
* libxml - 포괄적
* lxml - 포괄적
* mistralai - 포괄적
* mongoengine - 포괄적
* openai - 포괄적
* pandas - 부분적
* paramiko - 포괄적
* peewee - 포괄적
* pickle - 포괄적
* pilyaml - 포괄적
* pyca/cryptography - 포괄적
* pymongo - 포괄적
* pymssql - 포괄적
* pyramid - 포괄적
* psycopg - 포괄적
* python-ldap - 포괄적
* Python Standard Library - 포괄적
* requests - 포괄적
* sqlite3 (또는 pysqlite2) - 포괄적
* sqlalchemy - 포괄적
* turboGears - 포괄적
* urllib - 포괄적
* werkzeug - 포괄적

## 기능

다음 기능이 Snyk에서 Python을 지원합니다:

| Snyk Open Source                                     | Snyk Code                                                    |
| ---------------------------------------------------- | ------------------------------------------------------------ |
| <ul><li>PR 수정</li><li>라이선스 스캐닝</li><li>보고서</li></ul> | <ul><li>보고서</li><li>사용자 정의 규칙</li><li>Interfile 분석</li></ul> |

## Python 버전 지원

일부 Python 프로젝트에는 특정 Python 버전이 필요한 종속성이 포함될 수 있습니다. 따라서 Snyk가 생성하는 종속성 트리에 영향을 미치는 스캔 시 사용하는 Python 버전을 명시해야 합니다.

CLI와 Git 통합에서 Snyk이 사용하는 Python 버전을 지정할 수 있습니다.

Python 버전 및 Pip, Poetry, Pipevn, setup.py의 설치 및 사용 정보에 대한 자세한 내용은 [Python을 위한 Snyk CLI](snyk-cli-for-python.md)를 참조하십시오.

Python 버전 및 설치 및 Python 및 pip의 사용, Poetry 및 pipenv의 사용에 대한 자세한 정보는 [Git 저장소 및 Python](git-repositories-and-python.md)를 참조하십시오.

### Pipenv 및 지원되는 Python 버전

지원되는 Python 버전은 `3.8`, `3.9`, `3.10`, `3.11`, `3.12`입니다.

Snyk는 각 `Pipfile`에 지정된 Python 버전 정보를 사용하여 스캔에 사용할 주요 및 보조 버전을 선택합니다. 예를 들어:

```python
[requires]
python_version = "3.6"
```

특정 패치 버전은 무시되며, Snyk는 각 시리즈에서 최신 패치 버전을 사용합니다.

`Pipfile`에 다음이 포함된 경우 Snyk는 Python `3.10`을 사용합니다:

* Python 버전 정보가 없는 경우
* 주요 버전만 있는 경우
* 지원되지 않는 버전인 경우

### Poetry 및 지원되는 Python 버전

Poetry 프로젝트의 Python 버전을 Snyk에 알려줄 필요가 없습니다.

Poetry 파일에는 네이티브 도구를 실행하지 않고도 전체 종속성 트리를 작성하는 데 필요한 충분한 정보가 포함되어 있습니다.

## IDE 및 Snyk를 사용한 CI/CD

지원되는 IDE를 사용하여 Python을 작성하는 경우 Python 매니페스트 파일을 올바르게 스캔하려면 일부 구성을 추가해야 합니다.

가상 환경을 사용하는 경우 Snyk 통합 설정의 **추가 옵션** 텍스트 입력란에 `PYTHON_PATH`를 추가해야 합니다. 예를 들어 `--command=.venv/bin/python`입니다. Snyk는 IDE에서 볼 수 있듯이 프로젝트의 루트에 `*req*.txt` 파일을 찾으려고 합니다.

그러나 프로젝트의 루트 내 다른 디렉토리에 매니페스트 파일이 있는 경우 Snyk는 이를 식별할 수 없습니다. Snyk가가 이를 찾으려면 `--all-projects` 옵션을 사용해야 합니다. Snyk는 그런 다음 각 프로젝트 디렉토리를 재귀적으로 검색하여 모든 매니페스트 파일을 찾습니다.

각 디렉토리가 실행에 다른 가상 환경을 필요로 하는 경우 Snyk 스캔은 실행에 필요한 종속성을 찾기 위해 하나의 가상 환경을 사용하므로 성공하지 못할 수 있습니다. 이 경우 각 프로젝트 디렉토리에 나열된 모든 종속성에 대한 취약성 정보를 얻기 위해 IDE보다는 CLI 또는 Git 통합을 사용하는 것이 좋습니다.

## Python을 위한 Snyk 문제해결

도움이 필요한 경우, [Snyk 지원팀에 문의](https://support.snyk.io)하십시오.
