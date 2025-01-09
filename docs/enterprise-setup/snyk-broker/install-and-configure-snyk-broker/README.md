---
description: Broker와의 SCM 통합은 , , (Dockerfile), 부분 확인 필요
---

# Install and configure Snyk Broker

## Snyk Broker 사용 방법

Snyk Broker는 특별한 통합을 통해 Snyk과 중계 역할을 하는 오픈 소스 도구로, [snyk.io](http://snyk.io/)가 인터넷에 접근할 수 없는 저장소의 코드를 스캔하고 결과를 반환할 수 있도록 지원합니다. Broker와의 SCM 통합은 , , (Dockerfile), 를 지원합니다. 자세한 내용은 [Snyk Broker와의 통합](../#integrations-with-snyk-broker)을 참조하십시오.

Snyk Broker에 대한 포괄적인 정보 및 작동 방식, 배포 방법, 커밋 서명, 업그레이드 및 문제 해결 방법은 전체 [Snyk Broker 사용 설명서](../)를 참조하십시오.

## **배포 옵션**

<table data-card-size="large" data-view="cards" data-full-width="false"><thead><tr><th></th><th></th><th></th><th data-hidden data-card-cover data-type="files"></th><th data-hidden data-card-target data-type="content-ref"></th></tr></thead><tbody><tr><td><strong>빌트인</strong> <a href="https://github.com/snyk/snyk-broker-helm"><strong>Broker Helm Chart</strong></a>를 사용하여 Snyk Broker를 설치합니다. 자세한 내용은 <a href="install-and-configure-broker-using-helm.md">Helm을 사용하여 Broker 설치 및 구성</a>을 참조하십시오.</td><td></td><td></td><td><a href="../../../.gitbook/assets/helmkube.png">helmkube.png</a></td><td><a href="install-and-configure-broker-using-helm.md">install-and-configure-broker-using-helm.md</a></td></tr><tr><td><strong>Snyk 제공하는</strong> <strong>도커 이미지를 사용하여</strong> <strong>Snyk Broker를 설치합니다</strong>. 자세한 내용은 <a href="install-and-configure-broker-using-docker.md">도커를 사용하여 Broker 설치 및 구성</a>을 참조하십시오.</td><td></td><td></td><td><a href="../../../.gitbook/assets/Docker-Logo (1).jpg">Docker-Logo (1).jpg</a></td><td><a href="install-and-configure-broker-using-docker.md">install-and-configure-broker-using-docker.md</a></td></tr></tbody></table>

또한 [각 Github 릴리스](https://github.com/snyk/broker/releases)에 대해 사용 가능한 이진 파일을 사용할 수 있습니다. 그러나 이 방법은 Snyk 제품 지원 범위를 벗어나며 특정 사용 사례만을 위해 제공됩니다. 컨테이너 버전을 강력히 권장합니다.

## 개발자를 위한 추가 정보

업그레이드가 필요한 경우 [Snyk Broker 클라이언트 업그레이드](../upgrade-the-snyk-broker-client.md)를 참조하십시오.

문제 해결 정보는 [Broker 문제 해결](../troubleshooting-broker.md) 페이지에서 제공됩니다.

라이센스인 [Apache License, Version 2.0](https://github.com/snyk/broker/blob/master/LICENSE)를 확인할 수 있습니다.

풀 리퀘스트를 제출하려면 [기여 방법](https://github.com/snyk/broker/blob/master/.github/CONTRIBUTING.md)을 참조하십시오.

Broker에 대한 구체적인 정보는 [보안](https://github.com/snyk/broker/blob/master/SECURITY.md)을 참조하십시오.
