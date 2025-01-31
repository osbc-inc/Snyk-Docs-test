# 통합을 구성하여 Kubernetes 구성 파일에서 보안 문제 찾기(현재 IaC)

{% hint style="info" %}
이 페이지는 현재 IaC에만 적용됩니다.
{% endhint %}

Snyk은 소스 코드 저장소에 저장된 Kubernetes 구성을 테스트하고 모니터링하여 [Kubernetes 환경을 보호하는 방법](https://snyk.io/learn/kubernetes-security/)에 대한 정보, 팁 및 트릭을 제공합니다. 이를 통해 프로덕션으로 푸시되기 전에 잘못된 구성을 잡아내고 취약점에 대한 해결책을 제공할 수 있습니다.

## 지원되는 Git 저장소 및 Kubernetes 파일 형식

Snyk은 통합 Git 저장소에서 가져올 때 JSON 및 YAML 형식의 Kubernetes 구성 파일을 스캔합니다.

## Snyk를 구성하여 Kubernetes 구성 파일 스캔

### **Kubernetes 구성 파일을 스캔하기 위한 사전 요구 사항**

* 조직의 관리자 액세스
* Git 저장소 액세스 및 권한 부여\
  자세한 내용은 [Git 저장소 (SCM) 통합](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 참조하세요.

### **Snyk를 구성하여 Kubernetes 구성 파일 스캔**

* Snyk 웹 UI ([app.snyk.io](https://app.snyk.io))에 로그인하고 관리하려는 관련 그룹 및 조직으로 이동합니다.\
  통합은 조직별로 관리됩니다.
* Infrastructure as Code 설정에서 Snyk가 인프라스트럭처를 코드 파일로 감지하도록 설정을 전환합니다.
* 필요한 경우, 예제에서 Kubernetes 탭의 **인프라스트럭처를 코드로 감지하기 위한 설정**을 검토하고 조정합니다.\
  스캔할 파일 유형인 클라우드 형성, 테라폼 둘 중 하나를 선택하고 각 배포에 대해 심각도 수준을 선택하세요.

<figure><img src="https://docs.snyk.io/~gitbook/image?url=https%3A%2F%2F2533899886-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F-MdwVZ6HOZriajCf5nXH%252Fuploads%252F6nr3X5R9GQQEv74L0olw%252Fimage.png%3Falt%3Dmedia%26token%3D8464b11f-447f-4715-a7fa-9d0d9671baa5&#x26;width=768&#x26;dpr=4&#x26;quality=100&#x26;sign=419597f9&#x26;sv=2" alt="인프라스트럭처를 코드로 감지하기 위한 설정 선택"><figcaption><p>인프라스트럭처를 코드로 감지하기 위한 설정 선택</p></figcaption></figure>

사용 가능한 테스트 횟수는 계정 요금제에 따라 달라집니다. 자세한 내용은 [요금제 및 가격 책정](https://snyk.io/plans/) 페이지를 참조하세요.
