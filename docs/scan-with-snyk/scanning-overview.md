# 스캐닝 개요

{% hint style="info" %}
귀하의 [가격제도](../implement-snyk/enterprise-implementation-guide/trial-limitations.md)에 따라 계정의 스캔이 제한될 수 있습니다. 자세한 내용은 [테스트로 간주되는 것은 무엇입니까?](../working-with-snyk/what-counts-as-a-test.md)를 참조하십시오.
{% endhint %}

Snyk은 개발자를 우선시하여 직접 IDE, 워크플로우, 자동화 파이프라인에 통합하여 보안 전문성을 도구 상자에 추가함으로써 개발 작업을 보호합니다. 이 접근 방식을 통해 다음을 수행할 수 있습니다:

- Snyk를 사용하여 후속 강제보다는 초기 인증에 집중할 수 있습니다.
- 코드를 커밋하기 전에 Project에서 작업하는 동안 스캔을 실행합니다. 이를 통해 변경을 요구하는 문제를 초기에 발견하여 재작업을 최소화할 수 있습니다.
- 각 패키지를 작성하기 전에 추가하고 테스트합니다.
- 코드의 주요 부분을 작성한 후 계속 작업하기 전에 이를 스캔합니다.

## 안전한 코드 설계 방법 학습

다음 자료는 모든 사용자를 위해 제공됩니다:

- [Snyk Advisor](https://snyk.io/advisor): 건강한 오픈 소스 패키지 또는 기본 이미지를 선택하여 코드 개발을 돕습니다.
- [Snyk Learn](https://learn.snyk.io/): 코드를 안전하게 작성하는 방법을 배우고 Snyk 사용 방법에 대한 교육을 제공합니다.

## 코드 작성 및 배포

- [Snyk CLI](../snyk-cli/)을 사용하여 로컬 머신에서 스캔할 수 있습니다. 이는 오픈 소스 및 정적 코드뿐만 아니라 컨테이너 및 인프라 코드 구성, Terraform 플랜 파일 등의 변수로 템플릿화된 복잡한 파일을 스캔하는 데 유용합니다.
- [Snyk IDE 플러그인](../scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/)을 사용하여 개발 환경에서 오픈 소스 패키지, 퍼스트파티 코드, 인프라 코드 (IaC) 쿠버네티스 배포 파일을 생성하는 동안 Project를 만들면서 테스트할 수 있습니다.
- [Git 통합](../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 사용하여 코드 및 배포된 응용 프로그램의 Git 저장소에서 보안을 강화할 수 있습니다.
- [CI/CD 통합](../scm-ide-and-ci-cd-integrations/snyk-ci-cd-integrations/)을 사용하여 통합 및 배포 파이프라인에서 빌드 실패를 발생시켜 코드에 있는 취약점을 방지할 수 있습니다.

## 운영 중인 코드 모니터링

코드를 운영 환경으로 통합하기 전에 오픈 소스 및 컨테이너 Project에 도입된 문제를 식별하려면 [`snyk monitor`](../snyk-cli/commands/monitor.md) 또는 [`snyk container monitor`](../snyk-cli/commands/container-monitor.md) CLI 명령을 사용하여 이러한 Project의 취약점을 모니터링하여 운영 환경으로 밀어넣기 전에 문제를 식별할 수 있습니다.

자세한 내용은 [정기적 간격으로 Project 모니터링](../snyk-cli/scan-and-maintain-projects-using-the-cli/monitor-your-projects-at-regular-intervals.md)을 참조하십시오.

## Snyk를 사용하여 문제 관리 및 수정

응용 프로그램을 처음 스캔할 때 수백 개 이상의 문제가 발견된다면 문제의 우선순위를 정하는 것이 중요합니다. 자세한 내용은 [문제의 우선순위 설정](../manage-risk/prioritize-issues-for-fixing/)를 참조하십시오.

Snyk은 문제를 반응적 및 예방적으로 해결하기 위한 기능을 제공합니다:

- **예방적인 방법**
  - [Snyk Advisor](https://snyk.io/advisor)를 사용하여 설계를 시작할 수 있는 더 나은 패키지를 식별합니다.
  - CLI 및 IDE 플러그인을 사용하여 개발 중에 테스트합니다.
  - 패키지를 추가한 후 설치되었는지 확인하고 코드를 작성하기 전에 보안 스캔을 실행합니다.
- **수정 조언**
  Snyk는 이 조언을 통해 패키지 매니페스트에 업데이트가 필요한 최상위 패키지를 계산하거나 안전한 코드 줄을 업데이트하는 방법을 제공하며 결과 화면에 조언을 표시합니다.
- **자동화**
  - Snyk는 새로운 취약성이 감지되면 자동 수정 풀 리퀘스트를 생성할 수 있습니다.
  - 패키지의 새 버전이 나올 때 종속성 업데이트 관련 풀 리퀘스트를 생성할 수 있습니다. 이는 기술 부채를 해결하기 위해 의존 관계를 업데이트하는 PR 터치를 제공하여 도와줍니다.

## 배포 및 롤아웃 권장 사항

#### 소규모 기업

스타트업, 소규모 팀, 개인 및 오픈 소스 유지 관리자는 보통 Git를 사용하여 응용 프로그램을 도입하여 몇 분이면 결과를 얻고 거의 즉시 문제에 대응을 시작합니다. 소규모 팀은 민첩성을 가지고 있어 자신들의 워크플로우에 가장 잘 어울리는 것을 파악하는 데 혜택을 봅니다. &#x20;

자세한 내용은 [팀 구현 가이드](../implement-snyk/team-implementation-guide/)를 참조하십시오.

#### 대규모 기업

수백 개의 애플리케이션을 개발하는 대기업의 경우 개발자 참여와 채택을 보장하고 롤아웃 경험을 확보하기 위해 더 느린 접근 방식을 권장합니다.

자세한 내용은 [엔터프라이즈 구현 가이드](../implement-snyk/enterprise-implementation-guide/)를 참조하십시오.