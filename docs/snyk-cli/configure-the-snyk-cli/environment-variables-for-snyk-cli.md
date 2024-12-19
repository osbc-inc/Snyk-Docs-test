# Snyk CLI를 위한 환경 변수

CLI 설정을 변경하기 위해 다음 환경 변수를 설정할 수 있습니다.

`SNYK_TOKEN`

`SNYK_TOKEN`을 설정하여 Snyk 구성 설정 (`~/.config/configstore/snyk.json`)에 있는 토큰을 재정의할 수 있습니다. CI/CD 환경에서 `SNYK_TOKEN`을 사용하세요. `SNYK_TOKEN`을 설정한 후에는 CLI를 사용하여 [시작](../getting-started-with-the-snyk-cli.md)할 수 있습니다.

계정 토큰을 얻는 방법에 대한 정보는 [계정을 사용하여 CLI 인증](../authenticate-to-use-the-cli.md)을 참조하십시오. 또한 서비스 계정을 사용하여 인증할 수도 있습니다. 자세한 내용은 [서비스 계정](../../enterprise-setup/service-accounts/)을 참조하십시오. 추가 정보는 [제3자 도구용 인증](../../enterprise-setup/authentication-for-third-party-tools.md)을 참조하십시오.

`SNYK_CFG_<KEY>`

`snyk config` 옵션으로 사용할 수 있는 모든 키를 재정의할 수 있습니다.

예를 들어, `SNYK_CFG_ORG=myorg`는 `config`에 있는 기본 org 옵션을 "myorg"로 재정의합니다.

`SNYK_REGISTRY_USERNAME`

`snyk container` 명령에 대해 컨테이너 레지스트리에 연결할 때 사용할 사용자 이름을 지정합니다. `--username` 플래그를 사용하면 이 값이 재정의됩니다. Docker가 설치된 상태에서는 로컬 Docker 이진 자격 증명을 우선합니다.

`SNYK_REGISTRY_PASSWORD`

`snyk container` 명령에 대해 컨테이너 레지스트리에 연결할 때 사용할 암호를 지정합니다. `--password` 플래그를 사용하면 이 값이 재정의됩니다. Docker가 설치된 상태에서는 로컬 Docker 이진 자격 증명을 우선합니다.