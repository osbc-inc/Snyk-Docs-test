# CLI 사용을 인증합니다.

프로젝트를 스캔하려면 Snyk로 인증해야 합니다.&#x20;

{% hint style="info" %}
시스템 기본 환경 `SNYK-US-01`에 없는 경우 [`snyk config environment`](commands/config-environment.md) 명령을 사용하여 `snyk auth`를 실행하기 전에 환경을 설정하십시오.
{% endhint %}

## CLI 로컬 사용을 위한 인증 방법

### OAuth 2.0 프로토콜을 사용하여 인증하는 단계

CLI를 로컬로 사용할 때 **Snyk는 OAuth 2.0 프로토콜을 사용할 것을 권장합니다.** 다음 단계를 따라 주세요:

1. `snyk auth` CLI 명령을 실행합니다.
2. 로그인이 요청된 경우 로그인합니다.
3. 다음 페이지에서 CLI가 귀하를 대신하여 작동할 수 있는 권한을 요청합니다. **앱 액세스 부여**를 클릭합니다.
4. 성공적으로 인증한 후 확인 메시지를 확인한 다음 브라우저 창을 닫고 터미널의 CLI로 돌아갑니다.&#x20;

인증이 완료되면 액세스 및 리프레시 토큰 쌍이 로컬에 저장되어 향후 사용에 준비됩니다.&#x20;

문제가 발생하는 경우 [OAuth 2.0 인증이 작동하지 않음](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/troubleshooting-ides/how-to-set-environment-variables-by-operating-system-os-for-ides-and-cli-1.md)을 참조하십시오.

{% hint style="info" %}
OAuth 2.0 토큰은 정적이 아닙니다. 이러한 토큰은 Snyk 계정 페이지에서 복사할 수 없습니다.
{% endhint %}

### Snyk API 토큰을 가져와 인증하는 단계

{% hint style="warning" %}
이 방법은 OAuth 2.0 방법에 비해 뒷받침이 떨어집니다.
{% endhint %}

Snyk API 토큰을 사용하여 인증하려면 다음 단계를 따르십시오:

1. `snyk auth --auth-type=token` CLI 명령을 실행합니다.
2. 로그인이 요청된 경우 로그인합니다.
3. 다음 페이지에서 CLI 또는 IDE 플러그인을 계정과 연결하기 위해 기계를 인증하라는 메시지가 나타납니다. **인증**을 클릭합니다.
4. 성공적으로 인증한 후 확인 메시지가 표시됩니다. 브라우저 창을 닫고 터미널의 CLI로 돌아갑니다.&#x20;

대화형 과정을 완료하면 API 토큰이 로컬에 저장되어 향후 사용에 준비됩니다.&#x20;

모든 이후 `test` 명령은 자동으로 인증됩니다.&#x20;

### 알려진 Snyk API 토큰을 사용하여 인증하는 단계

개인 API 토큰을 **일반 계정 설정**(사용자 이름 아래)에서 복사한 다음 CLI에서 로컬로 사용하도록 구성할 수 있습니다.

모든 CLI `test` 명령은 환경 변수 `SNYK_TOKEN`을 자동으로 인식하고 이를 인증에 사용할 수 있습니다. 자세한 내용은 [Snyk CLI를 위한 환경 변수](configure-the-snyk-cli/environment-variables-for-snyk-cli.md)를 참조하세요.

API 토큰 기반의 인증을 사용하려면 `SNYK_TOKEN` 환경 변수를 설정하고 `test` 명령을 실행하십시오. 예를 들어:\
`SNYK_TOKEN=<SNYK_API_TOKEN> snyk test`

대신, 환경 변수를 내보내어 향후 `test` 명령에서 사용할 수 있도록 만들 수도 있습니다:\
`export SNYK_TOKEN=<SNYK_API_TOKEN>`\
`snyk test`

이러한 형태의 인증은 특히 CI/CD 파이프라인에 유용합니다. [CI/CD 파이프라인에서 CLI 사용을 인증하는 방법](authenticate-to-use-the-cli.md#how-to-authenticate-to-use-the-cli-in-ci-cd-pipelines)을 참조하십시오.

나중에 사용할 Snyk API 토큰을 로컬로 저장하려면 다음 CLI 명령을 실행하십시오:\
`snyk auth <SNYK_API_TOKEN>`

이어지는 모든 테스트 호출은 자동으로 인증됩니다. 자세한 내용은 [Auth 명령 도움말](commands/auth.md)을 참조하십시오.

## CI/CD 파이프라인에서 CLI 사용을 인증하는 방법

**무료 및 팀 플랜 사용자**는 OAuth 2.0을 사용하는 대신 CI/CD 파이프라인에서 이 방법을 더 자주 사용할 것입니다. **기업 플랜 고객**은 CI/CD 파이프라인에서 [**서비스 계정**](../enterprise-setup/service-accounts/)을 사용할 것을 권장합니다. API 토큰 또는 서비스 계정 토큰을 언제 사용해야 하는지에 대한 자세한 내용은 [API용 인증](../snyk-api/rest-api/authentication-for-api/)을 참조하십시오.

모든 CLI `test` 명령은 환경 변수 `SNYK_TOKEN`을 자동으로 인식하고 이를 인증에 사용할 수 있습니다. 자세한 내용은 [Snyk CLI를 위한 환경 변수](configure-the-snyk-cli/environment-variables-for-snyk-cli.md)를 참조하세요.

API 토큰 기반의 인증을 사용하려면 `SNYK_TOKEN` 환경 변수를 설정하고 `test` 명령을 실행하십시오. 예를 들어:\
`SNYK_TOKEN=<SNYK_API_TOKEN> snyk test`

대신, 환경 변수를 내보내어 향후 `test` 명령에서 사용할 수도 있습니다:\
`export SNYK_TOKEN=<SNYK_API_TOKEN>`\
`snyk test`

이러한 형태의 인증은 특히 CI/CD 파이프라인에 유용합니다. [CI/CD 파이프라인에서 CLI 사용을 인증하는 방법](authenticate-to-use-the-cli.md#how-to-authenticate-to-use-the-cli-in-ci-cd-pipelines)을 참조하십시오.

나중에 사용할 Snyk API 토큰을 로컬로 저장하려면 다음 CLI 명령을 실행하십시오:\
`snyk auth <SNYK_API_TOKEN>`

이어지는 모든 테스트 호출은 자동으로 인증됩니다. 자세한 내용은 [Auth 명령 도움말](commands/auth.md)을 참조하십시오.