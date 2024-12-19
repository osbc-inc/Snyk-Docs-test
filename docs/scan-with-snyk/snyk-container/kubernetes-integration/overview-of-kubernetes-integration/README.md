# 쿠버네티스 통합 개요

{% hint style="info" %}
**기능 가용성**\
쿠버네티스 통합은 엔터프라이즈 플랜에서만 사용할 수 있습니다. 자세한 내용은 [플랜 및 가격](https://snyk.io/plans/)을 참조하십시오.
{% endhint %}

Snyk은 쿠버네티스와 통합하여 실행 중인 작업 부하를 가져와 스캔할 수 있습니다. 이를 통해 해당 작업 부하를 덜 안전하게 만들 수 있는 이미지와 구성에서 취약점을 식별할 수 있습니다. 작업 부하가 가져오되면 Snyk은 계속하여 그들을 모니터링하고 새 이미지가 배포되고 작업 부하 구성이 변경됨에 따라 추가적인 보안 문제를 식별합니다.

## 쿠버네티스 통합 작동 방식

쿠버네티스 통합은 다음 프로세스를 따릅니다:

1. 관리자가 클러스터에 컨트롤러를 설치하고, 적합한 권한이 있는 서비스 계정 토큰과 고유한 통합 ID를 사용하여 통합을 인증합니다. 필요한 권한 및 자세한 내용은 [Snyk 컨트롤러 설치 전 요구 사항](../install-the-snyk-controller/#prerequisites-for-installing-the-snyk-controller)을 참조하십시오.
2. 다음 중 하나의 옵션으로 컨트롤러를 설치합니다:
   * [Snyk Helm을 사용하여 컨트롤러 설치 (Azure 및 Google Cloud Platform)](../install-the-snyk-controller/install-the-snyk-controller-with-helm-azure-and-google-cloud-platform.md)
   * [Amazon Elastic Kubernetes Service (Amazon EKS)에 Snyk 컨트롤러 설치](../install-the-snyk-controller/install-the-snyk-controller-on-amazon-elastic-kubernetes-service-amazon-eks.md)
   * [OpenShift 및 OperatorHub를 사용하여 Snyk 컨트롤러 설치](../install-the-snyk-controller/install-the-snyk-controller-with-openshift-4-and-operatorhub.md)
3. 컨트롤러는 쿠버네티스 API와 통신하여 클러스터에서 실행 중인 작업 부하(예: 배포, ReplicationController, CronJob 등)를 확인하고 연결된 이미지를 찾아 클러스터에서 직접 취약점을 스캔합니다.
4. 협력자가 가져오는 각 작업 부하에 대해 Snyk은 각 이미지에서 발견된 취약점과 작업 부하에서 식별된 구성 문제 요약을 표시합니다.
5. Snyk은 가져온 작업 부하를 계속 모니터링하고 새로 발견된 취약점을 프로젝트에 영향을 미칠 때마다 보고합니다.
6. 구성에 따라 취약점이 발견되면 Snyk은 이메일이나 Slack을 통해 즉시 통지하여 즉각 조치를 취할 수 있게 합니다.

{% hint style="info" %}
데이터베이스의 건강을 유지하기 위해 Snyk은 변경되거나 업데이트되지 않은 작업 부하 관련 정보를 8일 동안 제거합니다. 이는 작업 부하를 다시 스캔할 때 오류가 발생할 수 있습니다.

이미지와 해당 프로젝트가 제거되면, 그리고 8일 안에 메타데이터가 여전히 Snyk 데이터베이스에 남아 있는 동안 동일한 작업 부하를 다시 가져오는 경우, 프로젝트를 다시 생성할 수 있습니다.
{% endhint %}

## Snyk 및 쿠버네티스 통합 아키텍처

<figure><img src="../../../../.gitbook/assets/System Diagram-Kubernetes integration (1).jpg" alt="쿠버네티스 통합 아키텍처 다이어그램"><figcaption><p>쿠버네티스 통합 아키텍처 다이어그램</p></figcaption></figure>

이 다이어그램에 표시된 프로세스를 기반으로:

1. 고객의 Snyk 조직이 쿠버네티스 통합을 활성화합니다.
2. 고객이 Snyk 컨트롤러를 Kubernetes 클러스터에 설치합니다.
3. Snyk 컨트롤러가 이미지 정보를 읽고 컨테이너 레지스트리에서 이미지를 가져옵니다.
4. Snyk 컨트롤러가 이미지를 스캔합니다.
5. Snyk 컨트롤러가 스캔 결과를 분석할 문제를 식별하기 위해 Snyk 플랫폼에 전송합니다.
6. 고객이 Snyk 플랫폼에서 취약점 문제를 확인합니다.

## 약관

{{Snyk Container}} 쿠버네티스 통합은 Red Hat UBI (Universal Base Image)를 사용합니다.

이 애플리케이션을 다운로드하거나 사용하기 전에 [Red Hat 웹사이트](https://www.redhat.com/en/about/agreements)에서 찾을 수 있는 Red Hat 구독 약관에 동의해야 합니다. 이 약관에 동의하지 않는 경우 애플리케이션을 다운로드하거나 사용하지 마십시오.

Red Hat 엔터프라이즈 약정(또는 Red Hat과의 합의된 다른 약정)에서 컨테이너와 관련된 가입 서비스를 규정하는 조항이 있는 경우, 기존 약정이 지배 약정이 될 것입니다.

## 쿠버네티스 통합 문제 해결

쿠버네티스 통합이나 Snyk 컨트롤러에 대한 문제 해결 정보 및 일반적인 온보딩 오류에 대한 자세한 내용은 [쿠버네티스 통합 문제 해결](https://support.snyk.io/s/article/Kubernetes-Integration-troubleshooting) 지원 페이지를 참조하십시오.