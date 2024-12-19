# JetBrains 플러그인에 대한 문제 해결

{% hint style="warning" %}
Snyk 플러그인은 유통업체에서 End Of Life (EOL) 상태에 도달한 모든 운영 체제에서 지원되지 않습니다.&#x20;
{% endhint %}

## 자세한 로그 가져오기: 디버그 로그

각 JetBrains IDE는 로그에 쉽게 액세스할 수 있습니다.

다음 단계를 따라 디버그 수준으로 로깅을 설정하십시오.

플러그인 로그를 가져오려면 **도움말** > **Finder에서 로그 보기** (Mac) 또는 **탐색기에서 로그 보기** (Windows)로 이동하십시오.

<figure><img src="../../../.gitbook/assets/image (487).png" alt="Finder에서 로그 보기"><figcaption><p>Finder에서 로그 보기</p></figcaption></figure>

IDE에서 로그 수준을 디버그로 변경할 수 있습니다:

빠르게 `Shift Shift`를 누르고 **액션** 탭을 선택하십시오. 그런 다음 `debug`를 검색하십시오. 또는 메뉴에서 디버그 로그 설정을 선택하십시오 (Jetbrains Rider에는 사용 불가).

<figure><img src="../../../.gitbook/assets/image (488).png" alt="액션 탭"><figcaption><p>액션 탭</p></figcaption></figure>

<figure><img src="../../../.gitbook/assets/image (489).png" alt="액션 검색"><figcaption><p>액션 검색</p></figcaption></figure>

- `{{Snyk Code}}`는 일반 디버그 로깅을 활성화합니다 (플러그인 버전 2.6.0까지만).
- `{{Snyk Code}}RequestLogging`은 {{Snyk Code}} API와 통신할 때 HTTP 요청의 상세 프로토콜을 활성화합니다 (플러그인 버전 2.6.0까지만).
- `Snyk Language Server`는 백그라운드에서 언어 서버의 디버그 로깅을 활성화합니다.

<figure><img src="../../../.gitbook/assets/image (490).png" alt="Snyk Language Server 구성"><figcaption><p>Snyk Language Server 구성</p></figcaption></figure>

## 신뢰할 수 있는 루트 인증서 문제

플러그인이 잘못된 구성 때문에 네트워크 문제를 겪으면 IDE가 사용자 지정 인증서를 해결하는 방법에 대해 [JetBrains 문서](https://www.jetbrains.com/help/idea/ssl-certificates.html)를 참조하십시오.

## IntelliJ에서 {{Snyk Code}} 체크박스 비활성화

때때로 JetBrains IntelliJ 플러그인의 {{Snyk Code}} 체크박스가 비활성화됩니다. 몇 가지 가능한 이유는 다음과 같습니다:

- **네트워크 또는 프록시 설정:** 네트워크 또는 프록시 설정이 올바르게 구성되지 않은 경우 체크박스가 비활성화될 수 있습니다. MITM 프록시 및 인증서 대체 여부를 확인하십시오. 또한 CLI 또는 cURL과 같은 다른 도구를 사용하여 엔드포인트 API 및 deeproxy에 연결할 수 있는지 확인하십시오.&#x20;
- **잘못된 엔드포인트 주소:** {{Snyk Code}} 플러그인 구성에서 엔드포인트 주소가 잘못되어 있는 경우 체크박스가 비활성화됩니다. 이를 해결하려면 엔드포인트 주소가 올바른지 확인하고 지침에 따라 IntelliJ를 재시작하십시오.
- **{{Snyk Code}}가 서버 측에서 비활성화됨:** Snyk 조직 설정에서 {{Snyk Code}}가 비활성화된 경우 체크박스가 비활성화됩니다. 이를 해결하려면 IntelliJ 설정에서 표시된 지침을 따르십시오. IDE를 재시작하십시오.
- **JetBrains 로그를 확인하십시오:** 추가 정보는 [IDE 로그 파일 찾기](https://intellij-support.jetbrains.com/hc/en-us/articles/207241085-Locating-IDE-log-files)를 참조하십시오.

## 정의되지 않은 Python 버전

Snyk은 종속성을 스캔하고 찾기 위해 Python을 사용합니다.&#x20;

여러 Python 버전을 사용하는 경우 올바른 Python 명령을 실행하는 데 `-command` 옵션을 사용하십시오. 플러그인은 프로젝트와 관련된 Python 버전을 감지하지 못합니다.

## {{Snyk Container}} 및 Kubernetes JetBrains 통합 작동 방식

JetBrains 플러그인은 Kubernetes 작업량 파일을 스캔하고 사용된 이미지를 수집합니다. 플러그인이 컨테이너 이미지를 올바르게 스캔하고 있는지 문제 해결하려면 다음을 확인할 수 있습니다:

- 프로젝트의 Kubernetes YAML 파일에 이미지 정의가 있는지 확인하십시오. 이미지가 `imageValue:imageKey` 형식으로 매핑된 이미지 이름으로 지정되어 있는지 확인하십시오. 예를 들어, `image:nginx:1.17.1`와 같이 이미지 yaml 속성에 이미지가 지정되어 있는지 확인하십시오.
- 컨테이너 이미지가 로컬로 성공적으로 빌드되었거나 컨테이너 레지스트리에 푸시되었는지 확인하십시오. 또한 컨테이너 이미지를 Kubernetes YAML 파일에서 참조하기 전에 이를 확인하는 것이 좋습니다.

에러가 발생할 경우 [지원팀에 문의](https://snyk.zendesk.com/agent/dashboard)하십시오.

각 발견된 이미지에 대해 Snyk CLI를 사용하여 테스트를 수행하십시오.

- 이미지에 대한 테스트를 수행할 때 CLI는 Docker 데몬에서 이미지를 다운로드합니다(로컬에 이미지가 없는 경우).
- Snyk는 컨테이너 스캔 범위를 확대할 예정이므로 지원되기를 원하는 Dockerfile 또는 워크플로 같은 더 많은 파일이 있다면 특성 요청을 [Snyk 지원팀에 제출](https://support.snyk.io)하십시오.