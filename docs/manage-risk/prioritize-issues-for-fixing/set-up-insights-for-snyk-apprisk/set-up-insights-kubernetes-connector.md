# Insights 설정: Kubernetes 커넥터

{% hint style="info" %}
Insights 기능은 Snyk AppRisk Pro와 함께만 사용할 수 있습니다.

Snyk는 가장 효과적인 통합을 달성하고 계속해서 확장되는 기능 세트에 액세스하기 위해 [Snyk 런타임 센서](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)를 설치하는 것을 권장합니다.
{% endhint %}

## Snyk AppRisk용 Kubernetes 커넥터란 무엇인가요?

한 가지 목표는 네트워크 구성을 통해 공개적으로 접근 가능한 워크로드에 대한 리스크 요인을 식별하는 것입니다. 이를 위해 Snyk는 배포된 이미지가 어떤 워크로드에 배치되었고 구성되어 있는지 이해해야 합니다.&#x20;

따라서 Snyk는 다음 정보를 수집해야 합니다:

- 배포된 이미지 및 해당 ID 및 SHA 목록.
- 사용 중인 네트워킹 서비스와 해당 구성과 같은 서비스.

Kubernetes 커넥터는 이 정보를 수집하기 위해 Kubernetes 클러스터에 배포된 에이전트입니다.&#x20;

## Snyk AppRisk용 Kubernetes 커넥터 시작하기

### Snyk AppRisk용 Kubernetes 커넥터 전제조건

Kubernetes 클러스터에 Kubernetes 커넥터를 배포하기 전에 다음 사항을 준비해야 합니다:

- **Snyk 조직**: 수집된 Kubernetes 정보가 전송되어 저장될 Snyk 조직. 이는 새로운 조직이어야 하며, Snyk 프로젝트가 포함된 것과 동일할 필요는 없지만, 동일한 Snyk 그룹에 있어야 합니다.
- **특별히 생성된 Snyk 서비스 계정**: Kubernetes 커넥터와 함께 사용할 서비스 계정. 서비스 계정 생성 방법에 대한 지침은 [서비스 계정](../../../enterprise-setup/service-accounts/)에서 확인하세요. 역할 및 권한에 대해 Snyk는 다음을 권장합니다:
  - 이 서비스 계정을 위한 새로운 특정 역할 생성
  - 최소 권한 접근법 사용, 새로운 특정 역할에 **Kubernetes 리소스 발행**에 필요한 유일한 권한을 부여합니다.

### 단계 1: Snyk 조직 생성

Kubernetes 커넥터를 위해 별도의 조직을 생성하는 경우, Snyk 문서의 [Snyk 조직 생성](../../../snyk-admin/groups-and-organizations/organizations/create-and-delete-organizations.md#create-an-organization) 단계를 따르세요. 새 Snyk 조직은 다른 Snyk 조직과 동일한 Snyk 그룹에 있어야 합니다.&#x20;

별도의 Snyk 조직을 생성하지 않는 경우, 다음 단계를 계속하세요.

### 단계 2: 새로운 역할 생성

이 문서의 단계를 따라 [새로운 역할 생성](../../../snyk-admin/user-roles/user-role-management.md#create-a-role)을 수행하세요.

이 예시는 **Kubernetes 커넥터**라는 새로운 역할을 생성하는 것을 보여줍니다.

<figure><img src="../../../.gitbook/assets/image (14) (1) (1).png" alt="Insights 역할을 위한 Kubernetes 커넥터 생성"><figcaption><p>Insights 역할을 위한 Kubernetes 커넥터 생성</p></figcaption></figure>

### 단계 3: 이 역할에 권한 할당

새롭게 생성된 역할로 이동하여 [편집](../../../snyk-admin/user-roles/user-role-management.md#edit-a-role)을 선택하세요. 역할 생성 직후에도 즉시 이 페이지로 이동됩니다.&#x20;

페이지 하단으로 스크롤하여 **Kubernetes 리소스 발행** 권한을 선택하고 변경 사항을 저장하려면 **역할 권한 업데이트** 버튼을 클릭하세요.&#x20;

<figure><img src="../../../.gitbook/assets/image (12) (1) (1).png" alt="Kubernetes 리소스 발행 권한"><figcaption><p>Kubernetes 리소스 발행 권한</p></figcaption></figure>

### 단계 4: 서비스 계정 생성 및 역할 할당

다음으로, Kubernetes 커넥터와 통합하기 위해 [새로운 서비스 계정](../../../enterprise-setup/service-accounts/)을 만드세요.

{% hint style="info" %}
Snyk는 해당  서비스 계정을 Kubernetes 에이전트를 사용하거나 생성한 Snyk 조직에 대해 생성하는 것을 권장합니다.
{% endhint %}

**Snyk 조직 -> 설정 -> 서비스 계정**으로 이동하세요.

원하는 이름으로 새 서비스 계정을 생성하고 드롭다운에서 이전 단계에서 만든 역할을 선택하세요.

<figure><img src="../../../.gitbook/assets/image (11) (2) (1).png" alt="Insights Kubernetes 에이전트 역할 선택"><figcaption><p>Insights Kubernetes 에이전트 역할 선택</p></figcaption></figure>

서비스 계정이 생성되면 API 토큰이 표시됩니다. API 토큰을 복사하고 안전한 위치에 저장하세요. 이 정보는 Helm 차트에서 에이전트를 구성하는 데 필요합니다.

### 단계 5: Kubernetes 클러스터에 Kubernetes 커넥터 설치

Snyk는 에이전트를 배포하기 위해 Helm Chart를 사용하는 것을 권장합니다. Helm Chart는 클러스터에서 에이전트가 실행되기 위한 관련 권한을 생성합니다. Helm Chart를 설치하는 사용자는 Kubernetes 클러스터에서 새로운 역할을 생성할 권한이 필요합니다. \
\
[최신 릴리스 버전을 배포하기 위해 Kubernetes-scanner GitHub repo의 지침](https://github.com/snyk/kubernetes-scanner)을 따르세요.

{% hint style="info" %}
Kubernetes 커넥터를 올바르게 설정했는지 확인하려면 **이슈** 페이지의 **Insights 설정** 탭으로 이동하여 **Kubernetes 리소스** 테이블을 확인하여 Insights가 액세스할 수 있는 데이터를 확인하세요.
{% endhint %}

## FAQ

#### **[Kubernetes 모니터](../../../scan-with-snyk/snyk-container/kubernetes-integration/overview-of-kubernetes-integration/)(Snyk 컨트롤러 또는 Snyk 모니터로도 불림), Snyk AppRisk용 Kubernetes 커넥터, 그리고** [**Snyk 런타임 센서**](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)** 간의 차이는 무엇인가요?**

- Kubernetes **모니터**는 Kubernetes 클러스터의 워크로드에서 이미지를 추출하고 취약점을 스캔합니다. Kubernetes 모니터는 **배포된** 리스크 요인을 보고합니다.
- Snyk AppRisk용 Kubernetes **커넥터**는 Kubernetes 클러스터의 워크로드 구성을 추출합니다. Kubernetes 커넥터는 **공개되는** 및 **배포된** 리스크 요인을 보고합니다.
- [Snyk 런타임 센서](../../snyk-apprisk/integrations-for-snyk-apprisk/snyk-runtime-sensor.md)는 Kubernetes 클러스터에서 배포를 모니터링하고 수집된 데이터를 Snyk로 보냅니다. Snyk 런타임 센서는 **배포된** 및 **로드된 패키지** 리스크 요인을 보고합니다.

{% hint style="warning" %}
리스크 요인은 구현한 통합 옵션에 따라 사용 가능합니다.
{% endhint %}

#### **우선순위 지정을 위해 모든 세 가지 통합 옵션이 필요한가요, 아니면 하나만으로 충분한가요?**

설치 옵션은 하나만 필요합니다. Snyk는 배포된 및 로드된 패키지 리스크 요인을 보고하는 Snyk 런타임 센서를 설치하는 것을 권장합니다.&#x20;

#### **이미 기존 에이전트를 사용하는 고객인데 Kubernetes 커넥터도 설치해야 하나요?**

배포된 및 공개적으로 접근 가능한 리스크 요인을 모두 볼 경우 Kubernetes 커넥터를 Kubernetes 클러스터에 설치해야 합니다.

기존 에이전트만 설치된 경우, Snyk는 배포된 리스크 요인 만 계산할 수 있습니다.

#### **Kubernetes 커넥터가 수집하는 워크로드 데이터는 무엇인가요?**

Snyk가 수집하는 [데이터 목록](https://github.com/snyk/kubernetes-scanner/blob/main/helm/kubernetes-scanner/values.yaml)은 Kubernetes 스캐너 GitHub 리포지토리에 있습니다.

#### **Kubernetes 커넥터를 삭제하면 데이터는 어떻게 되나요?**

Kubernetes 커넥터를 삭제하면 데이터와 리스크 요인이 48시간 동안 유지됩니다.

#### **Kubernetes 커넥터가 작동을 멈추면 데이터는 어떻게 되나요?**

만약 어떤 이유로 인해 Kubernetes 커넥터가 작동을 중지한다면 데이터와 리스크 요인은 48시간 동안 유지됩니다.