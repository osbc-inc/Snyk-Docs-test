# Snyk Code Local Engine

{% hint style="info" %}
**릴리스 상태**

Snyk Code Local Engine은 이른 액세스 상태이며 엔터프라이즈 계획에서만 이용할 수 있습니다. 그룹이나 조직에서 이를 설정하려면 Snyk 계정 팀에 문의하십시오.
{% endhint %}

Snyk Code Local Engine (SCLE)은 코드를 인터넷에 업로드하지 않고 스캔을 로컬에서 수행할 수 있는 Snyk Code Engine의 완전히 포함된 버전입니다. 로컬 엔진을 사용하면 스캔이 로컬에서만 수행되고 스캔 결과가 Snyk로 업로드되어 Snyk 웹 UI에서 확인할 수 있습니다.

이 고수준 아키텍처 다이어그램은 구성 요소와 그들의 상호작용을 보여줍니다. Snyk는 Git 저장소의 요청을 Snyk Code Local Engine을 통해 스캔하고, 그 결과를 Snyk로 반환합니다.

<figure><img src="../../.gitbook/assets/Screen Shot 2021-11-11 at 2.36.41 PM.png" alt="Snyk Code Local Engine 고수준 아키텍처"><figcaption><p>Snyk Code Local Engine 고수준 아키텍처</p></figcaption></figure>

## Snyk Code Local Engine 시스템 요구 사항

Snyk Code Local Engine을 배포하려면 다음과 같은 핵심 요구 사항이 있습니다:

* Kubernetes 버전 1.21.0 - 1.28.0:
  * _권장:_ 전용 Kubernetes 클러스터
  *   클러스터에서 \*.snyk.io로의 웹소켓을 지원하는 아웃바운드 HTTPS 연결

      이 연결은 모든 흐름 (CLI, IDE, SCM 및 PR Checks)에 필요합니다.
  * Kubernetes – 다음 중 하나:
    * 관리되는 공용 클라우드 Kubernetes 서비스 - EKS, AKS, GKE\
      \- 또는 -
    * 관리되지 않는 Kubernetes (셀프 호스팅 클러스터)
  * PR Checks와 Snyk CLI 지원에는 Kubernetes 클러스터로의 네트워크 액세스가 필요합니다:
    * Kubernetes 클러스터에서 Git 및 CI/CD 도구로
    * Snyk CLI를 실행하는 사용자로부터 Kubernetes 클러스터로
* Helm 버전 3.8.0 이상
* x86 CPU가 있는 환경

### Snyk Code Local Engine 리소스 요구 사항

Snyk Code Local Engine은 모듈식이며 특정 통합을 실행하거나 필요에 따라 모두 실행할 수 있습니다. 또한 원하는대로 확장할 수 있으며 단일 또는 여러 노드로 구성할 수 있습니다.

각 Kubernetes 노드는 실행을 위해 적어도 다음의 자유 리소스가 필요합니다:

* 55GB RAM
* 14 Core CPU
* 50GB 임시 저장소

최소 노드의 크기는 환경에 필요한 것에 따라 다를 수 있지만, 이것들이 필요한 최소 자유 리소스입니다.

다양한 Snyk Code Local Engine 버전의 총 필요 리소스는 다음 목록에서 식별됩니다. 실제 메모리 및 저장소 소비는 사용량 및 스캔된 저장소의 크기에 따라 다릅니다. 최소 총 필요 리소스는 다중 노드로 나누어질 수 있습니다.

| 배포 옵션                | 필요한 리소스                                                           | 사용 사례                                                                   |
| -------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------------- |
| CLI                  | <p>- 165GB RAM<br>- 60 Core CPU<br>- 55GB Ephemeral Storage</p>   | - 파이프라인에서 SCLE 실행                                                       |
| IDE                  | <p>- 165GB RAM<br>- 60 Core CPU<br>- 55GB Ephemeral Storage</p>   | - 개발자들이 IDE에서 사용                                                        |
| SCM                  | <p>- 170GB RAM<br>- 65 Core CPU<br>- 65GB Ephemeral Storage</p>   | - 모니터링을 위해 저장소 가져오기                                                     |
| SCM 및 PR Checks      | <p>- 200GB RAM<br>- 90 Core CPU<br>- 160GB Ephemeral Storage</p>  | <p>- 모니터링을 위해 저장소 가져오기<br>- 새 취약점을 위해 모든 PR 스캔</p>                      |
| SCM, PR Checks 및 CLI | <p>- 220GB RAM<br>- 100 Core CPU<br>- 160GB Ephemeral Storage</p> | <p>- 파이프라인에서 SCLE 실행<br>- 모니터링을 위해 저장소 가져오기<br>- 새 취약점을 위해 모든 PR 스캔</p> |
| 전체 배포 (모든 기능)        | <p>- 330GB RAM<br>- 140 Core CPU<br>- 220GB Ephemeral Storage</p> | - 상기 모든 사항                                                              |

## CLI 및 IDE와 Snyk Code Local Engine 사용하기

Snyk CLI와 IDE를 Snyk Code Local Engine과 함께 사용하려면 Snyk 계정 팀에 Snyk Code Local Engine URL을 제공해야 합니다. 당신의 조직을 위해 실행 중인 Snyk Code Local Engine의 URL입니다.\
\
CSM이 조직을 위해 URL을 구성한 후, **Settings** --> **Snyk Code** 에서 확인할 수 있습니다:

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252Fe38hwQgZJS5T6xEfJq7g%252FSnyk%2520Code%2520Local%2520Engine%2520settings%2520showing%2520Local%2520Engine%2520URL.png%3Falt%3Dmedia%26token%3D61421bb4-2c3d-492a-8f30-7bfe867e0a0c&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=630b0817&#x26;sv=2" alt="Snyk Code Local Engine 설정 Local Engine URL"><figcaption><p>Snyk Code Local Engine 설정 Local Engine URL</p></figcaption></figure>

## 로컬 엔진 구성 및 배포

환경에서 로컬 엔진을 구성하고 배포하는 방법은 로컬 엔진 설치 패키지의 Readme 파일에 안내되어 있습니다. 자세한 정보와 도움이 필요한 경우 Snyk 계정 팀에 문의하십시오.
