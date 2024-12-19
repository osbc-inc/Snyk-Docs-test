# Nexus Repository - Snyk Broker을 위한 환경 변수

## Nexus 3 구성을 위한 환경 변수

다음 환경 변수는 Nexus 3에 대한 Broker 클라이언트를 사용자 정의하는 데 필요합니다.

`BROKER_TOKEN`\
Snyk Broker 토큰은 Nexus 통합 설정에서 얻을 수 있습니다 (**통합 > Nexus**).

`BASE_NEXUS_URL`\
Nexus 3 배포의 URL입니다.\
예:\
`BASE_NEXUS_URL=https://[<username_or_token><password_or_token>]@<your.nexus.hostname>`\
슬래시(/)로 끝나면 안 됩니다.\
다음 필드는 선택 사항입니다:\
`Auth`: 인증이 필요 없으면 생략합니다.\
평문 또는 두 부분 토큰(`Nexus Pro`)이 될 수 있습니다.\
인증 오류를 피하기 위해 사용자 이름, 비밀번호 및 토큰을 URL 인코딩합니다.\
간단한 예: `acme.com`\
복잡한 예: `https://alice:mypassword@acme.com`

`BROKER_CLIENT_VALIDATION_URL`\
Broker Client의 `systemcheck` 엔드포인트에서 확인하는 Nexus 유효성 URL입니다.\
Nexus 사용자가 `auth`를 필요로 하는 경우 `$BASE_NEXUS_URL/service/rest/v1/status/check`를 사용합니다.\
예:\
`https://<user>:<pass>@<your.nexus.hostname>/service/rest/v1/status/check`)\
그렇지 않은 경우 `$BASE_NEXUS_URL/service/rest/v1/status`를 사용합니다.\
예:\
`https://<your.nexus.hostname>/service/rest/v1/status`).

선택 사항. `RES_BODY_URL_SUB`\
이 URL 대체는 **npm/Yarn 통합에 필요**하며 자격 증명이 추가되지 않은 Nexus URL이어야 합니다.\
예:\
`https://<your.nexus.hostname>/repository`. 슬래시(/)로 끝나면 안 됩니다.

## Nexus 2 구성을 위한 환경 변수

다음 환경 변수는 Nexus 2에 대한 Broker 클라이언트를 사용자 정의하는 데 필요합니다:

`BROKER_TOKEN` - Snyk Broker 토큰은 Nexus 통합 설정에서 얻을 수 있습니다 (**통합 > Nexus**).

`BASE_NEXUS_URL`- Nexus 2 배포의 URL입니다.\
예:\
`BASE_NEXUS_URL=https://[username_or_token:password_or_token]@acme.com`\
슬래시(/)로 끝나면 안 됩니다.\
다음 필드는 선택 사항입니다:\
`Auth`: 인증이 필요 없으면 생략합니다.\
평문 또는 두 부분 토큰(`Nexus Pro`)이 될 수 있습니다.\
인증 오류를 피하기 위해 사용자 이름, 비밀번호 및 토큰을 URL 인코딩합니다.\
간단한 예: `https://acme.com`\
복잡한 예: `https://alice:mypassword@acme.com: 8000`

`NEXUS_URL`: 만약 귀하의 저장소가 /nexus/content 아래에 있지 않은 경우 재정의로 사용됩니다\
Nexus 내의 저장소 기본 위치를 가리키는 URL입니다. 기본적으로 브로커는 `BASE_NEXUS_URL`/nexus/content/ 를 가치로 가정합니다.

`RES_BODY_URL_SUB`\
기본 인증 자격 증명이 없는 Nexus 인스턴스의 URL입니다. **npm/Yarn 통합에만 필요**합니다. 슬래시(/)로 끝나면 안 됩니다.

예시:\
`https://acme.com/nexus/content/groups`\
`https://acme.com/nexus/content/repositories`