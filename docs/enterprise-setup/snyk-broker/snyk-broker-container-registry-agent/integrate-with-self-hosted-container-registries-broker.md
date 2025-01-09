# 자체 호스팅 컨테이너 레지스트리와 통합

{% hint style="info" %}
**기능 가용성**

Snyk 브로커와 함께 사용하는 Self-hosted 컨테이너 레지스트리는 엔터프라이즈 플랜에서만 사용 가능합니다.

자세한 정보는 [요금제 및 가격정보](https://snyk.io/plans/)를 참조하세요.
{% endhint %}

Snyk은 인터넷에 접근할 수 없는 자체 호스트된 개인 컨테이너 레지스트리와 통합하여 해당 레지스트리에 있는 컨테이너 이미지를 보다 안전하게 보호할 수 있도록 도와줄 수 있습니다.

자체 호스트된 컨테이너 레지스트리를 활성화하고 구성하려면 Snyk 팀에 문의하십시오.

## Self-hosted 컨테이너 레지스트리의 솔루션 구성 요소

자체 호스트된 컨테이너 레지스트리와의 통합에는 다음 구성 요소가 포함됩니다:

* **브로커 서버** - Snyk SaaS 백엔드에서 실행
* **브로커 클라이언트 및 컨테이너 레지스트리 에이전트** - 두 개의 Docker 이미지가 인프라에 배포되어 컨테이너 레지스트리를 안전한 방식으로 샘플링하고 해당 정보를 Snyk로 전송하는 두 개의 별도 서비스를 생성합니다.

<figure><img src="../../../.gitbook/assets/mceclip0-8-.png" alt="Snyk 브로커 컨테이너 레지스트리 에이전트의 고수준 아키텍처"><figcaption><p>Snyk 브로커 컨테이너 레지스트리 에이전트의 고수준 아키텍처</p></figcaption></figure>

브로커 클라이언트는 에이전트에게 연결 세부 정보를 제공합니다. 에이전트는 이러한 세부 정보를 사용하여 컨테이너 레지스트리에 연결하고 이미지를 스캔한 다음 브로커 통신을 사용하여 콜백을 통해 스캔 결과를 전송합니다.

브로커 통신은 브로커 클라이언트가 브로커 ID를 사용하여 Snyk 환경 내에서 실행되는 브로커 서버에 연결할 때 발생합니다.

자세한 내용은 [Snyk 브로커 - 컨테이너 레지스트리 에이전트](./)를 참조하세요.

## **지원되는 컨테이너 레지스트리**

* Docker Hub (유형: docker-hub)
* GCR (유형: gcr)
* ECR (유형: ecr)
* Azure (유형: acr)
* Artifactory (유형: artifactory-cr)
* Harbor (유형: harbor-cr)
* Quay (유형: quay-cr)
* GitHub (유형: github-cr)
* Nexus (유형: nexus-cr)
* DigitalOcean (유형: digitalocean-cr)
* GitLab (유형: gitlab-cr)
* Google Artifact Registry (유형: google-artifact-cr)
