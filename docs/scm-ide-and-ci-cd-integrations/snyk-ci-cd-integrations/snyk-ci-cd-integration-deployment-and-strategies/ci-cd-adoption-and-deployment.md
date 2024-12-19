# CI/CD 채택 및 배포

Snyk 통합을 사용할 결정을 내릴 때 소스 제어 관리(SCM) 통합과 CI/CD 통합의 장점을 비교하십시오. [SCM (Git) 및 CI/CD 통합 배포 소개](../../git-repository-and-ci-cd-integrations-comparisons.md)를 참조하십시오.

## CI/CD 통합 채택의 전형적인 단계

개발팀은 일반적으로 다음 단계에서 Snyk를 채택합니다:

1. [취약점 노출](ci-cd-adoption-and-deployment.md#stage-1-expose-vulnerabilities-snyk-monitor) (`snyk monitor`)
2. [게이트 키퍼로서 Snyk 사용](ci-cd-adoption-and-deployment.md#stage-2-use-snyk-as-a-gatekeeper-snyk-test) (`snyk test`)
3. [지속적인 모니터링](ci-cd-adoption-and-deployment.md#stage-3-continuous-monitoring-snyk-test-and-snyk-monitor) (`snyk test` 및 `snyk monitor`)

### **단계 1: 취약점 노출 (`snyk monitor`)**

일반적인 접근 방식은 개발 프로세스 중에 취약점을 노출하는 데 Snyk 결과를 사용하는 것입니다. 이를 통해 팀 구성원들 사이의 취약점을 더 잘 파악할 수 있습니다.

파이프라인에 처음으로 Snyk를 구현할 때는 `snyk monitor` 명령어만 사용하는 것이 권장됩니다. Snyk CI 플러그인 중 하나를 사용하는 경우 해당 플러그인을 빌드 실패하지 않도록 구성하는 것이 좋습니다.

모든 프로젝트에는 취약점이 존재하며, Snyk를 빌드 실패로 설정하면 Snyk 때문에 모든 빌드가 실패합니다. 이로 인해 팀이 빨리 실패 메시지로 압도되는 문제가 발생할 수 있습니다.

`snyk monitor`를 사용하여 결과를 노출하여 프로세스를 방해하지 않고 정보를 제공합니다.

`snyk monitor`에 대한 자세한 정보는 [`monitor` 명령어 도움말](../../../snyk-cli/commands/monitor.md)을 참조하십시오.

### **단계 2: 게이트 키퍼로서 Snyk 사용 (`snyk test`)**

Snyk를 게이트 키퍼로 사용하면 새로운 취약점의 도입을 방지할 수 있습니다 (때로는 "출혈을 막기"로 알려집니다).

팀이 응용 프로그램의 취약점을 이해하고 개발 주기 초반에 이를 수정하는 프로세스를 개발한 후에야 Snyk를 빌드 실패로 구성하여 응용 프로그램에 취약점이 도입되지 않도록 할 수 있습니다.

빌드에 `snyk test`를 추가하거나 실패 기능을 활성화하여 Snyk가 빌드를 실패하도록 설정한 다음 결과를 콘솔에 출력할 수 있습니다. 결과를 개발자 또는 데브옵스 팀이 빌드를 계속할지 중단할지 결정하는 데 사용할 수 있습니다.

`snyk test`에 대한 자세한 정보는 [`test` 명령어 도움말](../../../snyk-cli/commands/test.md)을 참조하십시오.

### **단계 3: 지속적인 모니터링 (`snyk test`** 및 **`snyk monitor`)**

Snyk를 취약점이 감지될 때 빌드 실패하도록 구성한 후에는 프로젝트의 성공적인 빌드 스냅샷을 Snyk에 보내 계속 모니터링하도록 구성할 수 있습니다.

이를 위해 파이프라인을 구성하여 `snyk test`가 성공적인 종료 코드를 반환하면 `snyk monitor`를 실행하도록 구성합니다.

## CI/CD 배포 방법

{% hint style="info" %}
이러한 모든 방법은 모두 동일한 Snyk 엔진을 사용하므로 선택한 배포 방법과 관계없이 동일한 결과를 제공합니다. 따라서 동일한 인수나 옵션이 적용됩니다.
{% endhint %}

파이프라인 내에서 Snyk를 구성하는 다양한 방법이 있습니다. 환경 및 기호에 따라 방법을 선택하십시오. 모든 방법이 성공적인 실행으로 이어질 것을 기대할 수 있습니다.

### **Snyk 네이티브 플러그인 사용**

Snyk 네이티브 플러그인은 대부분의 일반적인 CI/CD 도구에서 사용할 수 있습니다. 이러한 플러그인을 사용하면 설정 및 시작이 가장 쉽습니다. 플러그인에는 UI에 가장 일반적인 매개변수와 옵션이 포함되어 있습니다.

### **npm 방법을 사용하여 Snyk CLI 배포**

로컬에서 [CLI를 설치](../../../snyk-cli/install-or-update-the-snyk-cli/)하는 것과 유사한 단계를 따릅니다. 파이프라인 스크립트에서 npm 명령어를 실행할 수 있어야 합니다. 이 방법은 CLI 경험과 완전히 일치하므로 문제 해결 및 구성이 쉽습니다.

### **Snyk CLI 이진 버전 배포**

이진 프로그램 설치의 장점은 로컬 환경에 종속성이 없다는 것입니다. 예를 들어 파이프라인에서 npm 명령을 실행할 수 없는 경우 유용합니다.

CLI 이진 파일은 [CLI GitHub 저장소](https://github.com/snyk/cli/tags)에서 사용할 수 있습니다.

Snyk는 리눅스, 윈도우 및 기타 버전을 제공합니다.

### **Snyk 컨테이너 배포**

Snyk 이미지 중 하나를 사용하여 파이프라인에서 Snyk를 배포할 수도 있습니다. [Dockerhub](https://hub.docker.com/r/snyk/snyk)에 Snyk 이미지가 있습니다.

## Snyk CI/CD 통합 예시

이 리포지토리에서 CI/CD 도구에 대한 다양한 이진 및 npm 통합의 예시를 보여줍니다: [CI/CD 예시](https://github.com/snyk-labs/snyk-cicd-integration-examples).