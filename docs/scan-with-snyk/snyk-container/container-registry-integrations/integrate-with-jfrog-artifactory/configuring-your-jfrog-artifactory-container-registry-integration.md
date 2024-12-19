# JFrog Artifactory 컨테이너 레지스트리 통합 구성

이 페이지의 지침은 하나의 Artifactory 인스턴스를 컨테이너 레지스트리로 설정하고 Snyk 조직과의 통합을 활성화하는 방법을 설명합니다.

## Artifactory 컨테이너 레지스트리 통합을 위한 전제 조건

* Snyk에서 구성 중인 조직의 관리자여야 합니다.
* Artifactory와 통합하기 위해 사용자 자격 증명을 제공해야 합니다.
* Docker를 실행해야 합니다. Snyk는 이 통합을 위해 Docker 저장소 및 Docker 패키지 유형을 지원합니다.
* 자체 호스팅된 Artifactory 인스턴스를 사용 중인 경우 [Snyk Broker - 컨테이너 레지스트리 에이전트](../../../../enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/)를 참조하십시오.

## Artifactory 컨테이너 레지스트리 통합 구성

1. [Snyk 계정](https://app.snyk.io)에 로그인합니다.
2. **Integrations**로 이동합니다. 통합 목록에서 **Artifactory**를 선택합니다.\
   구성 페이지가 열립니다.
3. 자격 증명을 입력합니다:
   * **사용자 이름과 비밀번호 -** Artifactory 로그인 자격 증명을 사용합니다. SSO 구성을 사용 중인 경우, Artifactory에서 액세스 토큰을 생성하고 해당 토큰 세부 정보를 사용하여 로그인해야 합니다.
   * **컨테이너 레지스트리 이름 -** 전체 레지스트리 URL: `<subdomain>.jfrog.io/artifactory/api/docker/<repo-name>`

<figure><img src="https://user-images.githubusercontent.com/112600/144875482-078b715e-2834-469b-9983-7e88a65f175e.png" alt="" width="375"><figcaption><p>Artifactory 계정 자격 증명</p></figcaption></figure>

4. **변경 사항 저장**을 클릭합니다. 확인이 표시됩니다.

{% hint style="info" %}
통합을 설정하려면 Artifactory 자격 증명에 적절한 Artifactory 저장소의 읽기 권한이 있어야 합니다.
{% endhint %}

Snyk는 연결 값과 페이지를 테스트하고, 이 값을 입력한 통합 세부 정보를 화면 상단에 표시하여 페이지를 다시로드합니다. 화면 상단에 있는 확인 메시지는 세부 정보가 저장되었음을 나타냅니다. 연결에 실패하면 알림이 표시됩니다.