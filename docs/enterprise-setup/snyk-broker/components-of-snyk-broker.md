# Snyk Broker의 구성 요소

Broker 클라이언트와 서버는 코드로 향하는 터널로 함께 작동하여 [snyk.io](http://snyk.io/)에서 Git 리포지토리로 프록시를 통해 요청을 보내고, 모니터링된 리포지토리에서 필요한 파일(예: 매니페스트 파일)을 가져오며, Git 서비스에서 게시된 웹훅을 사용하여 결과를 가져옵니다. Broker 클라이언트를 설치하게 됩니다. Broker 서버는 Snyk에서 제공되며 설치가 필요하지 않습니다.

Broker 클라이언트는 내부 네트워크에서 실행되며 Git 토큰과 같은 중요 데이터를 네트워크 경계 내에 보관합니다. 터널을 통해 권한 부여된 데이터 목록 상의 요청만을 사용하여 스캔이 가능하게 합니다. 이는 Snyk가 리포지토리를 모니터링하는 데 필요한 최소한의 액세스 권한으로 접근 권한을 제한합니다. 자세한 정보는 [Snyk Broker용 승인된 데이터 목록](connections-with-snyk-broker.md#approved-data-list-for-snyk-broker)을 참조하십시오.

Snyk Code 에이전트, Snyk 컨테이너 레지스트리 에이전트 및 Infrastructure as Code용 Snyk Broker를 사용하는 데 필요한 구성 요소에 대한 정보는 [Broker 배포 구성 요소 정의](prepare-snyk-broker-for-deployment.md#define-your-broker-deployment-components)를 참조하십시오.

다음 다이어그램은 기본 구성 요소가 어떻게 작동하는지 설명합니다.

- 모든 데이터는 전송 및 정지 상태에서 암호화됩니다.
- 클라이언트와 서버 간 통신은 안전한 WebSocket 연결을 통해 이루어집니다.
- 통신이 발신적으로 시작되므로 수신 포트를 열 필요가 없습니다.
- 연결이 시작된 후, WebSocket 연결은 양방향입니다.

<figure><img src="../../.gitbook/assets/Snyk Broker diagram.png" alt="Snyk Broker WebSocket이 클라이언트에 의해 시작되고 HTTPS를 통해"><figcaption><p>Snyk Broker WebSocket이 클라이언트에 의해 시작되고 HTTPS를 통해</p></figcaption></figure>

Snyk와 Broker 클라이언트 간의 연결 및 허용된 요청에 대한 자세한 내용은 [Snyk Broker와의 연결](connections-with-snyk-broker.md)를 참조하십시오.