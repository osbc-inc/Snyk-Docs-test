# Artifactory Repository - Snyk Broker을 위한 환경 변수

다음 환경 변수들은 Artifactory Repository를 위해 Broker Client를 사용자 정의하는 데 필요합니다:

`BROKER_TOKEN` - Artifactory 통합 설정에서 가져온 Snyk Broker 토큰입니다 (**Integrations > Artifactory**).

`ARTIFACTORY_URL` - Artifactory 배포의 URL입니다. 예를 들어, `<yourdomain>.artifactory.com/artifactory`.

다음 필드들은 선택 사항입니다:

* _포트_: 필요한 포트 번호가 없는 경우 생략합니다.
* _기본 인증_: 기본 인증이 필요하지 않은 경우 생략합니다.\
  인증 오류를 방지하기 위해 사용자명과 암호 정보를 URL 인코드합니다.
* _프로토콜_: 기본값은 `https://`입니다.\
  인증서가 없는 경우 `http://`이 필요한 경우에만 지정해야 합니다.

옵션 필드가 있는 `ARTIFACTORY_URL` 형식:\
`[http://][username:password@]hostname[:port]/artifactory`\
예시:\
`http://alice:mypassword@acme.com:8080/artifactory`

옵션. `RES_BODY_URL_SUB` - http://를 포함하고 기본 인증 자격 증명이 없는 Artifactory 인스턴스의 URL입니다. **npm/Yarn 통합에만 필요합니다**.\
예시:\
`http://acme.com/artifactory`