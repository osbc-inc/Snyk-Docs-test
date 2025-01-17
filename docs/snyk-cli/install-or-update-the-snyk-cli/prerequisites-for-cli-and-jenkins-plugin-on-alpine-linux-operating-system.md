# Alpine Linux운영 체제에서 CLI 및 Jenkins 플러그인을 위한 전제 조건

Alpine Linux 운영 체제에서 Snyk CLI 및 Snyk Jenkins 플러그인을 실행하기 전에, 다음 단계에 따라 선행 조건을 설정하십시오.

1. `libstdc++`가 설치되었는지 확인합니다.\
   `libstdc++`를 설치하려면, 관련 컨테이너 내에서 터미널을 열고 `apk add libstdc++`를 실행하십시오.\
   명령어는 공유 라이브러리를 컨테이너 환경에 추가하거나 `/usr/lib`에 이미 설치된 경우 `OK`를 반환합니다.
2. 시스템에 Snyk CLI 버전 1.185.5 이상이 설치되어 있는지 확인합니다.\
   Snyk은 **최신** 버전을 설치하는 것을 권장합니다.
3. Snyk Jenkins 플러그인을 사용하려면, Snyk Jenkins 플러그인 버전 2.10.0 이상이 설치되어 있는지도 확인하십시오.\
   Snyk는 **최신** 버전을 설치하는 것을 권장합니다.
4. Snyk CLI를 설치한 후에는 반드시 [인증](https://docs.snyk.io/snyk-cli/authenticate-the-cli-with-your-account)해야 합니다.

추가 정보는 CI/CD 문서의 [Snyk CLI 설치 또는 업데이트](./) 및 [Jenkins 플러그인](../../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/jenkins-plugin-integration-with-snyk.md)를 참조하십시오.
