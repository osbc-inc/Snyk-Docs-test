# CLI에서 사용할 조직 선택하는 방법

CLI로 `snyk monitor` 및 `snyk test`와 같은 명령을 실행할 때, Snyk은 계정을 처음 등록할 때 만들어진 기본 조직을 사용합니다. 다른 조직에 대한 명령을 실행하려면 해당 조직을 전역적으로 또는 개별적으로 지정할 수 있습니다.

**참고**: Snyk CLI 도움말에서는 `--org=<ORG_ID>`를 사용하도록 명시하고 있으며, 이 문서에서는 `orgslugname`을 사용하도록 명시하고 있습니다. `ORG_ID`는 CLI와 API에서 모두 작동합니다. 조직 슬러그 이름은 CLI에서는 작동하지만 API에서는 작동하지 않습니다.

## 전역적으로 조직 지정하기

`snyk config set org=orgslugname`을 실행합니다. 참고: `orgslugname`은 Snyk UI에서 표시되는 조직 URL의 슬러그 이름과 일치해야 합니다: `https://app.snyk.io/org/[orgslugname]`.

## 개별적으로 조직 지정하기

개별 실행을 위해 전역 설정을 재정의하려면 다음과 같이 조직을 지정합니다:

`snyk monitor --org=orgslugname` 또는 `snyk test --org=orgslugname`.

예: `$ snyk test --org=my-team`