# Snyk Broker와의 연결

이 페이지는 Snyk과 브로커 클라이언트 간의 연결과 허용된 요청에 대한 세부 정보를 제공합니다.

## Snyk 브로커와의 수신 및 발신 연결 사용

브로커 클라이언트는 내부 네트워크에서 실행됩니다.

### Snyk에서 브로커 클라이언트로의 수신 연결

Snyk에서 브로커 클라이언트로 직접적인 수신 연결은 없습니다.

브로커 클라이언트는 [https://broker.snyk.io](https://broker.snyk.io)로 외부 연결을 생성합니다. 이는 웹소켓 연결을 설정하여 브로커 서버와의 통신을 허용합니다.

따라서 Snyk IP 주소를 허용할 필요가 없습니다. 대신, 브로커 클라이언트의 IP/포트를 허용할 수 있습니다.

### 브로커 클라이언트로의 발신 연결

브로커 클라이언트가 발신 연결을 초기화하여 웹소켓을 설정합니다.

브로커 클라이언트가 초기화한 웹소켓 연결이 설정된 후에, Snyk는 웹소켓을 통해 브로커 클라이언트로 수신 요청을 보낼 수 있습니다.

따라서 Snyk-특정 IP 주소나 다른 외부 IP 주소로부터 브로커 클라이언트로의 수신 연결을 허용할 필요가 없습니다.

## 허용된 요청

### **Snyk 브로커용 승인된 데이터 목록**

브로커 클라이언트는 수신 및 발신 요청에 대한 승인된 데이터 목록을 유지합니다. 이러한 승인된 목록에 있는 데이터만 요청할 수 있습니다. 이는 Snyk가 리포지토리를 모니터링하는 데 필요한 최소한의 액세스 권한만을 허용합니다.

### 허용된 수신 요청

에 대해 다음 요청이 허용됩니다:

* Snyk.io는 종속성 매니페스트 파일 및 `.snyk` 정책 파일만 가져오고 조회할 수 있습니다.
* 다른 소스 코드는 보거나 추출되거나 수정되지 않습니다.
* 취약점 정책에 포함된 추가적인 `.snyk` 파일을 지원하고, 무시 지시사항을 확인하기 위해 추가할 수 있습니다.

와 는 전체 저장소에 엑세스해야 합니다.

자세한 내용은 [Snyk가 데이터를 처리하는 방법](../../working-with-snyk/how-snyk-handles-your-data.md) 을 참조하세요.

### 허용된 발신 요청

브로커 클라이언트 설정을 구성하면, Git 리포지토리 웹훅이 자동 Snyk 스캔을 활성화하도록 설정되어 있습니다. 개발자가 새로운 풀 리퀘스트를 제출하거나 머지 이벤트를 트리거하면 자동적으로 Snyk 스캔이 실행됩니다.

웹훅 알림은 브로커 클라이언트를 통해 Snyk으로 전달되며, Snyk 작업과 관련된 이벤트에 대해서만 발생합니다: 브랜치로 푸시 및 풀 리퀘스트를 엽니다. 이 때 이벤트 데이터에는 종속성 매니페스트 파일 또는 `.snyk` 정책 파일이 포함되어야 합니다.

### 기본 승인된 데이터 목록 및 `accept.json` 파일

시기상조에 따라, 브로커 배포에서 [`accept.json` 파일을 추가하고 구성](snyk-broker-infrastructure-as-code-detection/)해야 할 수 있습니다. 이렇게 하면 브로커를 시작할 때 ACCEPT 규칙을 적용할 수 없게 됩니다.
