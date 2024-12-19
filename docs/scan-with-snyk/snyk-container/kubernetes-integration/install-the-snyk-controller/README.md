# Snyk 컨트롤러 설치

## Snyk 컨트롤러를 설치하기 위한 선행 조건

{% hint style="info" %}
**기능 가용성**\
Snyk 컨트롤러는 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 정보는 [플랜 및 가격](https://snyk.io/plans/)을 참조하세요.
{% endhint %}

Snyk 컨트롤러를 설치하기 전에 다음이 필요합니다:

- Snyk 조직의 관리자 계정이 있어야 합니다.
- [emptyDir](https://kubernetes.io/docs/concepts/storage/volumes/#emptydir)로 클러스터에 50GB 이상의 저장 공간이 있어야 합니다.
- 클러스터에서 실행할 때 요구되는 RAM 요구 사항은 아래 코드에 표시된 것보다 크거나 같아야 합니다.

  ```
  requests: cpu: "250m" memory: "400Mi"
  limits: cpu: "1" memory: "2Gi"
  ```

- Kubernetes 클러스터에는 Kubernetes `linux/amd64` 워커 노드가 있어야 합니다.
- Kubernetes 클러스터가 HTTPS를 통해 Snyk와 통신할 수 있어야 합니다.
- [Kubernetes 통합 기능 활성화](../overview-of-kubernetes-integration/enable-the-kubernetes-integration.md)하여 **통합 ID**를 가져옵니다.
- **그룹** 또는 **조직** **서비스 계정 토큰**을 생성해야 합니다. 자세한 정보는 [서비스 계정](../../../../enterprise-setup/service-accounts/)을 참조하세요. 다양한 역할에서 통합이 데이터를 게시할 수 있습니다:
   - 그룹 관리자
   - 조직 관리자
   - **쿠버네티스 리소스 발행** 권한을 가진 조직 사용자 지정 역할
- 로컬에 [Helm](https://helm.sh/docs/intro/install/)을 설치해야 합니다.

{% hint style="info" %}
Snyk 컨트롤러 Kubernetes 통합은 클러스터에 배포되도록 설계되었습니다. Fargate, Google Cloud Run, Azure Container Instances와 같은 서버리스 FaaS 배포 옵션에서 테스트되지 않았습니다. 서버리스 플랫폼에 배포하는 것은 신뢰할 수 없으며 권장되지 않으며 지원되지 않습니다.
{% endhint %}

Snyk 컨트롤러는 기본적으로 US-01 데이터 센터를 사용합니다. [대체 데이터 센터](../../../../working-with-snyk/regional-hosting-and-data-residency.md)에서 Snyk를 사용 중인 경우, 특정 배포를 위해 환경 변수를 통해 업스트림 엔드포인트 `integrationApi` URL을 변경해야 합니다:

- AU: [https://api.au.snyk.io/v2/kubernetes-upstream](https://api.au.snyk.io/v1/kubernetes-upstream)
- EU: [https://api.eu.snyk.io/v2/kubernetes-upstream](https://api.eu.snyk.io/v1/kubernetes-upstream)
- Helm 예시:

```
--set integrationApi=https://api.au.snyk.io/v2/kubernetes-upstream
```