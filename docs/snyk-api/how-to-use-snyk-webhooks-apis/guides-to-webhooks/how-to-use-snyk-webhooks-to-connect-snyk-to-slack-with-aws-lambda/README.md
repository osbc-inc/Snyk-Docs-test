# Snyk 웹훅을 사용하여 AWS Lambda를 통해 Snyk를 Slack에 연결하는 방법

Snyk Webhooks를 이용하여 Lambda 함수를 사용하여 Snyk에서 발견된 새로운 취약점을 Slack에서 받아들이고 필터링할 수 있습니다.

Snyk Webhooks를 Lambda 함수에 노출시키는 것은 API Gateway 및 Lambda 함수 URL을 통해 처리됩니다. 귀하의 요구 사항과 환경에 가장 적합한 옵션을 선택하십시오. API Gateway를 사용하지 않는 경우, 이 안내서의 지침을 무시할 수 있습니다.

**전제 조건**은 다음과 같습니다:

- 다음에 액세스할 수 있는 AWS 계정:
  - 새 역할 생성 제어 (또는 기존 역할 사용)
  - Lambda 함수 수정
  - API Gateway 수정 (Lambda 함수를 API Gateway를 통해 발행하는 경우)
- 조직 관리자 액세스 권한이 있는 Snyk 계정
- 기존 채널에 웹훅이 추가된 새 Slack 앱을 만들 수 있는 Slack 계정
- 터미널에서 npm 명령을 실행할 수 있는 능력

Snyk Webhook, AWS Lambda 함수 및 Slack 알림은 다음처럼 **작동**합니다:

- Snyk가 새 취약점을 발견할 때마다, Snyk API Webhook을 트리거합니다.
- 이는 Lambda 함수를 트리거하여 Slack로 알림을 보냅니다. Lambda 함수의 목표는 이러한 알림을 필터링하고 Slack 페이로드를 사용자에게 흥미로운 정보를 보여주도록 사용자화하는 것입니다.
- Slack 알림을 받으면 새로 발견된 취약점에 대응할 수 있습니다.

이 안내서는 AWS Lambda 함수를 사용하여 Snyk 웹훅에서 Slack으로 페이로드를 필터링하는 방법을 **설명**합니다.

다음은 **데이터 및 트래픽 흐름**을 설명합니다:

Snyk 프로젝트 스냅샷 웹훅은 API Gateway를 통해 헤더와 POST 본문을 전달하여 AWS Lambda 함수를 트리거합니다. 그런 다음 Lambda 함수는 해시 헤더 서명 확인이 성공하고 페이로드에 유효한 데이터가 포함되어 있는 경우 Slack 웹훅에 필터링된 페이로드(사용자 지정 메시지)를 보냅니다. Lambda 함수는 그런 다음 POST 본문을 필터링하여 사용자 지정 메시지를 구성합니다.

<figure><img src="https://lh6.googleusercontent.com/VROtTsX240dfLMERpOkm-5epOnvZxQUxjM-qKJYNEOtD_1flwBrpBTiJedo2Uy0RZz6kKplKNQQcINzOW3H30Lf7R9U0teZ4WvivBt1u7TdN_4J3ha_ZmY9wdn3xvXCNxl9036JdYeEzaBMtU53lo6e-do3Bhbmi4Y9tcWDO5y00NT_XRvmt5Z9ipg" alt="Snyk 웹훅을 사용하여 AWS Lambda를 통해 Snyk를 Slack에 연결하는 데이터 및 트래픽 흐름"><figcaption><p>Snyk 웹훅을 사용하여 AWS Lambda를 통해 Snyk를 Slack에 연결하는 데이터 및 트래픽 흐름</p></figcaption></figure>

Snyk 웹훅을 사용하는 데 **문제**가 발생하는 경우, 도움을 받기 위해 솔루션 엔지니어 또는 기술 성공 매니저에게 **문의**하십시오.