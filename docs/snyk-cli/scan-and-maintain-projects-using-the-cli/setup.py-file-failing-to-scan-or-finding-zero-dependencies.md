# 왜 내 setup.py 파일이 스캔되지 않거나 0개의 종속성을 찾지 못합니까?

`setup.py` 파일을 스캔하기 위해 `snyk test --file=setup.py` 명령을 실행할 때, 일반적으로 일부 Python `setup.py` 프로젝트가 실패하거나 0개의 종속성이 출력될 수 있습니다. 이는 Snyk가 모든 패키지가 나열된 `install_requires` 키만을 읽기 때문입니다.

다른 `requires` 키에 있는 종속성은 스캔되지 않습니다.

향후 지원은 변경될 수 있습니다.