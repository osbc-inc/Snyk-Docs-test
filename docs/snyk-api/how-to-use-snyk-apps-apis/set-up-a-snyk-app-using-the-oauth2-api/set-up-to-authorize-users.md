# 사용자 인증 설정

사용자가 Snyk 계정을 귀하의 App에 연결할 때, 선택한 조직 또는 그룹에 대한 액세스를 승인하고 요청된 스코프를 승인해야 합니다. 이 프로세스는 사용자를 Snyk Apps 인증 웹 페이지로 이동시키고 다음과 같은 매개변수를 전달하여 시작됩니다: `https://app.snyk.io/oauth2/authorize?response_type=code&client_id={clientId}&redirect_uri={redirectURI}&state={state}&code_challenge={codeChallenge}&code_challenge_method=S256`

귀하의 Snyk App이 여러 리디렉션 URI로 구성된 경우 `redirect_uri`를 전달하여 사용할 리디렉션 URI를 선택합니다. 주어진 값은 귀하의 Snyk App에 정의된 값 중 하나와 정확히 일치해야 합니다. 지정되지 않으면 Snyk App에 정의된 첫 번째 값이 사용됩니다.

`state` 값은 이 `/authorize` 호출에서 콜백으로 `redirect_uri`로 전달되는 어떤 App 특정 상태를 전달하는 데 사용됩니다(예: 사용자 ID). [CSRF 공격을 방지](https://datatracker.ietf.org/doc/html/rfc6749#section-10.12)하기 위해 귀하의 콜백에서 확인해야 합니다.

`code_challenge` 값은 생성된 코드 확인자의 SHA256 해시의 URL 안전한 base64-인코딩된 문자열입니다. 이 코드 확인자는 `/authorize`를 호출하기 전에 앱 측에서 생성된 매우 무작위의 문자열이어야 합니다. 인증 코드를 액세스 토큰으로 교환할 때 이 코드 확인자가 전송됩니다. 자세한 내용은 [OAuth Public Clients를 위한 Proof Key for Code Exchange](https://datatracker.ietf.org/doc/html/rfc7636)를 참조하십시오. Snyk Apps는 코드 확인 방법으로 `S256`만 지원하는 점에 유의하십시오.

<figure><img src="../../../.gitbook/assets/image (118) (1) (1).png" alt="조직에 대한 액세스 승인"><figcaption><p>조직에 대한 액세스 승인</p></figcaption></figure>

연결이 완료되면 사용자는 다음 단계에 필요한 쿼리 문자열 매개변수 `code`와 `state`가 추가된 지정된 리디렉션 URI로 리디렉션됩니다.
