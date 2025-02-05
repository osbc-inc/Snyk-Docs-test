# 오픈 소스용 Python

## 오픈 소스 지원을 위한 파이썬

지원되는 패키지 관리자 및 기능에 대해서는 [Python 상세](./)를 참조하십시오.

도움이 필요한 경우, [Snyk 지원팀에 문의](https://support.snyk.io).

{% hint style="info" %}
귀하의 요금제에 따라 일부 기능을 사용할 수 없는 경우가 있습니다. 자세한 내용은 [요금제 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

다음은 Snyk가 오픈 소스용 파이썬에서 지원하는 패키지 관리자를 요약한 것입니다.

| 패키지 관리자 기능                                                                                                                       | CLI 지원                                                                                                                           | SCM 지원 | 라이선스 스캐닝 | PR 수정 |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | ------ | -------- | ----- |
| [Pip](https://pypi.org/project/pip/)                                                                                             | ✔︎                                                                                                                               | ✔︎     | ✔︎       | ✔︎    |
| [Poetry](https://python-poetry.org)                                                                                              | ✔︎                                                                                                                               | ✔︎     | ✔︎       |       |
| [Pipenv](https://pipenv.pypa.io/en/latest/)                                                                                      | ✔︎                                                                                                                               | ✔︎     | ✔︎       |       |
| [setup.py](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/snyk-cli-for-python#setup.py-and-cli) | ✔︎[ 지침 참조](https://docs.snyk.io/supported-languages-package-managers-and-frameworks/python/snyk-cli-for-python#setup.py-and-cli) |        | ✔︎       |       |

프로젝트를 스캔하려면 먼저 해당 패키지 관리자를 설치하고 프로젝트에 지원되는 매니페스트 파일이 포함되어 있는지 확인해야 합니다.
