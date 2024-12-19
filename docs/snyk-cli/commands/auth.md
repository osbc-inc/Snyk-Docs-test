# 인증

## 사용법

`snyk auth [<API_TOKEN>] [<OPTIONS>]`

## 설명

`snyk auth` 명령은 Snyk CLI를 Snyk 계정과 연결하기 위해 귀하의 기기를 인증합니다.

`$ snyk auth`를 실행하면 브라우저 창이 열리며 Snyk 계정에 로그인하고 인증할 수 있는 프롬프트가 표시됩니다. 이 단계에서는 리포지토리 권한이 필요하지 않고 이메일 주소만 필요합니다.

인증이 완료되면 CLI를 사용할 수 있으며 [CLI로 시작하기](https://docs.snyk.io/snyk-cli/getting-started-with-the-cli)를 참조하십시오.

**참고:** 1.1293 버전부터 Snyk CLI는 브라우저를 통해 인증할 때 OAuth를 사용합니다.

OAuth는 더 짧은 수명을 가진 만료되는 권한을 발급하여 자동 새로 고침의 편의성을 제공하면서 향상된 보안을 제공합니다.

이전 버전의 Snyk CLI(1.1293 미만)는 전통적인 브라우저 상호작용을 통해 만료되지 않는 API 토큰을 획들했습니다.

Snyk API 토큰은 여전히 대체 옵션으로 사용할 수 있습니다. 다음과 같이 `snyk auth --auth-type=token` 옵션을 명시적으로 추가해야 합니다.

## 옵션

### `--auth-type=<TYPE>`

사용할 인증의 \<TYPE>을 지정합니다. 지원되는 유형은 `oauth` (1.1293.0 버전부터 기본값) 및 `token`입니다.

### `--client-secret=<SECRET>` 및 `--client-id=<ID>`

[OAuth2 클라이언트 자격 증명 권한 부여](https://docs.snyk.io/enterprise-configuration/service-accounts/service-accounts-using-oauth-2.0#oauth-2.0-with-client-secret)을 사용하기 위해 클라이언트 비밀번호와 ID를 설정할 수 있습니다.

두 값은 함께 제공되어야 합니다. `--auth-type=oauth`와 함께만 유효하며 그렇지 않으면 무시됩니다.

`<SECRET>` 및 `<ID>`를 얻는 방법에 대한 정보는 [OAuth 2.0를 사용하는 서비스 계정](https://docs.snyk.io/enterprise-setup/service-accounts/service-accounts-using-oauth-2.0#oauth-2.0-with-client-secret)을 참조하십시오.

## 토큰 값

일부 환경과 구성에서는 `<API_TOKEN>`을 사용해야 합니다. [CLI를 계정으로 인증](https://docs.snyk.io/snyk-cli/authenticate-the-cli-with-your-account)을 참조하십시오.

값은 사용자 토큰 또는 서비스 계정 토큰일 수 있습니다. [서비스 계정](https://docs.snyk.io/enterprise-setup/service-accounts)을 참조하십시오.

CI/CD 환경에서 `SNYK_TOKEN` 환경 변수를 사용하십시오. Snyk CLI 구성을 참조하십시오. [Snyk CLI 구성](https://docs.snyk.io/snyk-cli/configure-the-snyk-cli)

이 환경 변수를 설정한 후 CLI 명령을 사용할 수 있습니다.

## 디버그

디버그 로그를 출력하려면 `-d` 옵션을 사용하십시오.