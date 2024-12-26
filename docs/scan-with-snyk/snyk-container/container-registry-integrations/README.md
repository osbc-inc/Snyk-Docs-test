# 컨테이너 레지스트리 통합

## Snyk Container 통합 개요

Snyk Container은 컨테이너 보안을 모니터링하기 위한 통합을 제공합니다. 이러한 통합은 Snyk 사용자 및 고객을 위한 다양한 작업 흐름을 지원합니다.

사용하는 통합은 요구 사항과 작업 흐름에 따라 다릅니다. 하나의 통합부터 시작하여 나중에 다른 통합으로 이동하거나 팀이 성장함에 따라 통합을 조합하여 사용할 수 있습니다.

예를 들어, 이미지를 빌드하고 개발팀에 빠른 피드백을 제공하도록 Snyk CLI를 사용하고 나중에 Kubernetes 통합을 사용하여 프로덕션에서 실행되는 이미지에 대한 보증을 제공하는 것이 일반적입니다.

주요 Container 통합은 다음과 같습니다:
- **CLI:** 로컬 조사 또는 빌드한 이미지를 테스트하는 데 사용합니다. 이 통합은 빠른 피드백을 제공하고 기계에서 초기 테스트 및 CI의 게이트키퍼 역할로 사용할 수 있습니다. 또한, Snyk의 CD에서 스냅숏을 캡처하고 새로 발견된 취약점을 식별하는 도구로 기능합니다. 자세한 내용은 [Snyk CLI를 사용한 컨테이너 보안](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/) 참조.
- **SCM:** Snyk는 Git 저장소에서 Dockerfile을 직접 감지하고 덜 취약한 기본 이미지로 업데이트하기 위한 권장 사항을 제공할 수 있습니다. 자세한 내용은 [Dockerfile 스캔하기](../../../scan-with-snyk/snyk-container/scan-your-dockerfile/) 참조.
- **컨테이너 레지스트리:** 수많은 이미지를 테스트하는 데 사용하거나 여러 CI 파이프라인을 수정할 수 없는 경우에 사용합니다.
- **Kubernetes:** 작업 부하를 모니터링하고 작업이 어떻게 구성되어 실행되는지에 대한 추가 컨텍스트를 제공합니다. 자세한 내용은 [Kubernetes 통합 개요](../kubernetes-integration/overview-of-kubernetes-integration/) 참조.

{% hint style="info" %}
클라우드 호스팅된 컨테이너 레지스트리의 경우, Snyk는 2GB보다 큰 크기의 이미지를 가져오거나 스캔하지 않습니다. 이 크기를 초과하는 이미지를 스캔하려면 [Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)를 사용하십시오.
{% endhint %}

##  통합 선택

귀하의 요구 사항 및 맥락에 따라, 컨테이너 이미지의 새로운 취약점을 모니터링하기 위해 Snyk는 다음 중 하나의 통합 지점을 선택하는 것을 권장합니다: CLI, 컨테이너 레지스트리 통합 또는 Kubernetes 통합. 이러한 테스트 방법 중 하나를 선택함으로써 취약점 식별 시 최상의 결과를 보장하고 소음을 줄일 수 있습니다.

Snyk CLI는 CLI 옵션으로 매우 사용자 정의가 가능하기 때문에 다른 모든 통합보다 넓은 범위를 가지고 있습니다. 지속적 배포 파이프라인의 일부로 `snyk container monitor`를 실행하여 배포 중인 컨테이너 이미지의 스냅숏을 캡처할 수 있습니다. 더 많은 정보는 [Snyk Container용 Snyk CLI](../../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)를 참조하십시오.