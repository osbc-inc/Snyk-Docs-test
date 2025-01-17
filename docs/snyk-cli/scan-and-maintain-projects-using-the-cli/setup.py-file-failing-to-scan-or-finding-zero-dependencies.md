# Setup.py 파일 검사에 실패하거나 종속성이 0인 파일 발견

`setup.py` 파일을 스캔하기 위해 `snyk test --file=setup.py` 명령을 실행할 때, 일반적으로 일부 Python `setup.py` 프로젝트가 실패하거나 0개의 종속성이 출력될 수 있습니다. 이는 Snyk이 모든 패키지가 나열된 `install_requires` 키만을 읽기 때문입니다.

다른 `requires` 키에 있는 종속성은 스캔되지 않습니다.

향후 지원은 변경될 수 있습니다.
