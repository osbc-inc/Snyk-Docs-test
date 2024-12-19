# Helm 차트 설치를 위한 고급 구성

{% hint style="info" %}
**기본 이외 지역을 위한 멀티테넌트 설정**\
기본 이외의 지역에서 Snyk 브로커를 설정할 때는 특정 URL을 사용하는 추가 환경 변수가 필요합니다. URL 및 예제에 대한 자세한 정보는 [지역 호스팅 및 데이터 저장소, 브로커 URL](https://docs.snyk.io/working-with-snyk/regional-hosting-and-data-residency#broker-urls)를 참조하세요.
{% endhint %}

Helm을 사용하여 Snyk 브로커를 설정할 때 고급 매개변수를 설정할 수 있습니다. 아래 페이지에서 설명된 대로 설정할 수 있습니다:

- [브로커 Helm 차트 설치를 위한 사용자 정의 추가 옵션](custom-additional-options-for-broker-helm-chart-installation.md)
- [Snyk 브로커 Helm 설치와 함께 인그레스 옵션](ingress-options-with-snyk-broker-helm-installation.md)
- [동일한 네임스페이스에 여러 브로커 배포](deploying-multiple-brokers-in-the-same-namespace.md)
- [Helm 차트 설치를 위한 서비스 계정](service-accounts-for-helm-chart-installation.md)
- [Helm 차트 설치를 위한 멀티테넌트 설정](multi-tenant-settings-for-helm-chart-installation.md)
- [Snyk 오픈 소스 스캔 (SCA) 대용량 매니페스트 파일, Helm 설정](https://docs.snyk.io/enterprise-setup/snyk-broker/install-and-configure-snyk-broker/advanced-configuration-for-helm-chart-installation/snyk-open-source-scans-sca-of-large-manifest-files-helm-setup)
- [Kubernetes 시크릿 및 Helm 차트 설치](kubernetes-secrets-and-helm-chart-installation.md)
- [이미지 저장소, 탭 및 이미지 Pull 시크릿](image-repository-tab-and-image-pull-secret.md)
- [보안하지 않은 다운스트림 모드](insecure-downstream-mode.md)
- [브로커 Helm 차트 설치를 위한 프록시 설정](proxy-settings-for-broker-helm-chart-installation.md)
- [Helm을 사용하여 문제 해결 및 자체 인증서 제공을 위한 매개변수](parameters-for-troubleshooting-and-providing-your-own-certificate-with-helm.md)
- [Docker 및 Helm을 사용한 자격 증거 풀링](../advanced-configuration-for-snyk-broker-docker-installation/credential-pooling-with-docker-and-helm.md)