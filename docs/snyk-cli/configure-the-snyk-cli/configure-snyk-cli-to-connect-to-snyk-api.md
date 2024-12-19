# Snyk API와의 연결을 설정하는 방법

기본적으로 Snyk CLI는 `https://api.snyk.io/`에 연결합니다. 아래 변수들을 사용하여 연결을 구성할 수 있습니다.

`SNYK_API`

Snyk 요청에 사용할 API 호스트를 설정합니다. [지역 호스팅](../../working-with-snyk/regional-hosting-and-data-residency.md#cli-and-ci-pipelines-urls), 온프레미스 인스턴스 또는 프록시 서버를 사용할 때 유용합니다. `http` 프로토콜로 설정되면, CLI는 요청을 `https`로 업그레이드하며, `SNYK_HTTP_PROTOCOL_UPGRADE`가 `0`으로 설정되지 않은 경우에만 해당합니다.

`SNYK_HTTP_PROTOCOL_UPGRADE=0`

값을 `0`으로 설정하면, `http` URL을 대상으로 하는 API(및 CLI) 요청은 `https`로 업그레이드되지 않습니다. 프로토콜이 설정되지 않았을 경우, 기본 동작은 이러한 요청을 `http`에서 `https`로 업그레이드하는 것입니다. 예를 들어, 리버스 프록시에 유용합니다.

`SNYK_DISABLE_ANALYTICS=1`

모든 Snyk CLI 분석 기능을 비활성화합니다.

`SNYK_OAUTH_TOKEN=<OAuth 토큰>`

검증에 필요한 OAuth 토큰을 지정합니다.