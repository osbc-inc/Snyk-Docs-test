# Kubernetes 구성 파일에서 보안 문제를 스캔하고 수정합니다 \(현재 IaC\)

{% hint style="info" %}
이 페이지는 현재 IaC에만 적용됩니다.
{% endhint %}

Snyk Infrastructure as Code는 매니페스트 파일을 보안 취약점 및 Kubernetes 구성 파일을 잘못된 구성 및 보안 문제를 스캔합니다. 구성 파일이 스캔된 후, Snyk는 관리자가 구현한 설정에 따라 잘못된 구성 사항을 보고하고 그에 따른 수정을 권장합니다.

## Kubernetes 파일에서 문제를 스캔하고 수정하기 위한 전제 조건

* 관리자는 [조직을 통합](../scan-terraform-files/configure-your-integration-to-find-security-issues-in-your-terraform-files-current-iac.md)하고 선호하는 Git 리포지토리와 구성 파일의 감지를 활성화해야 합니다.
* Snyk 계정이 있어야 하며, 구성 파일은 JSON 또는 YAML 형식이어야 합니다.

Snyk Infrastructure as Code는 다음을 지원합니다:

* 배포, 포드 및 서비스
* Cron 작업, 작업, StatefulSet, 복제본 집합, DaemonSet 및 ReplicationController

## Kubernetes 구성 파일을 스캔하고 수정하기

* 계정에 로그인하고 관리하려는 해당 그룹 및 조직으로 이동합니다.
* 클라우드 구성 파일 감지가 관리자에 의해 활성화되기 전에 테스트를 위해 리포지토리를 가져온 경우, 해당 JSON 또는 YAML 구성 파일을 가져오기 위해 해당 리포지토리를 다시 가져와야 합니다.

리포지토리를 다시 가져올 때 클라우드 구성 파일을 가져오려면 Snyk가 구성 파일을 가져오고 테스트하며 이미 가져온 애플리케이션 매니페스트 파일을 다시 테스트하여 테스트 시간을 "현재"로 표시합니다.

* 관심 있는 프로젝트 링크를 클릭하여 스캔 결과를 확인하고 구성 파일을 수정합니다:

<figure><img src="../../../../.gitbook/assets/image (19) (2) (1) (1) (1) (1) (2) (1).png" alt="Kubernetes Project detail"><figcaption><p>Kubernetes 프로젝트 상세</p></figcaption></figure>