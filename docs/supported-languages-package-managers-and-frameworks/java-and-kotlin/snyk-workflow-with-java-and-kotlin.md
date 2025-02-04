# Java 및 Kotlin을 사용한 Snyk 워크플로

Snyk 팀은 Snyk를 워크플로에 통합하기 위한 플러그인을 개발했습니다:

* [**Gradle 플러그인**](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) **(커뮤니티 프로젝트)**
* [**Maven 플러그인**](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)

## 유효성 검사, 모니터링, 경보 및 Gatling

다음 능력은 모든 Snyk 사용자를 위해 제공됩니다:

### **Git 통합을 통해**

Snyk을 사용하여 코드에 제출된 변경 사항과 오픈 소스 패키지를 병합하기 전에 검증하기 위해 [PR 확인](../../scan-with-snyk/pull-requests/pull-request-checks/)을 실행할 수 있습니다. Snyk은 또한 정기적으로 기본 브랜치에서 재시험하고 경고를 알릴 수 있으며 결과를 표시할 수 있습니다.

이러한 결과는 다음에서 볼 수 있습니다:

* Snyk Code를 사용한 Code
* Snyk Open Source를 사용한 Open Source
  * 알려진 취약점 확인 (Snyk Open Source)
    * 알려진 취약성을 해결하기 위한 수정 풀 리퀘스트 생성 (Maven)
  * 라이선스 규정 준수 확인 (Snyk Open Source)(Maven)
  * 의존성 업데이트 - 기술적 부채 해결을 위한 업데이트 지정 (Snyk Open Source) (Maven)

Git 통합을 통해 매일 다음을 모니터링할 수 있습니다:

* Snyk Infrastructure as Code를 사용한 Infrastructure as Code

### **CI/CD 통합을 통해**

Snyk은 정책 위반에 대한 테스트 중 빌드 검사 실패로 QA 게이트를 제공할 수 있습니다.

Snyk은 다음과 같은 유연한 능력을 제공합니다:

* [Gradle 플러그인](https://snyk.io/blog/gradle-plugin-by-snyk-gradle-dependencies-scanning/) **(커뮤니티 프로젝트)**
* [Maven 플러그인](https://snyk.io/blog/snyk-maven-plugin-integrated-security-vulnerability-scanning-for-developers/)
* Jenkins, Circle CI 및 기타 플랫폼용 전용 플러그인 (관련 마켓플레이스 참조)
* [Github Actions](https://snyk.io/blog/building-a-secure-pipeline-with-github-actions/) 사용
* 대부분의 CI/CD 시스템에서 Snyk CLI 사용 가능(예시 참조: [GitHub 링크](https://github.com/snyk-labs/snyk-cicd-integration-examples))
  * 옵션이나 [snyk-filter](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/snyk-filter.md) 도구를 사용하여 기준에 따라 빌드 실패 처리
  * [컨테이너화된](https://hub.docker.com/r/snyk/snyk) 버전 사용 가능
* 파트너 플랫폼으로: Azure, Bitbucket 및 AWS는 Snyk과 함께 사용하기 위한 내장 파이프/구성요소를 가지고 있습니다.
  * Java에 대한 참고: Bitbucket Cloud에서 Git 통합을 사용하거나 사전 패키지화된 Bitbucket 파이프 대신 CLI를 사용하는 것이 좋습니다.

## 프로덕션 모니터링

* (Snyk 엔터프라이즈 플랜 전용) Kubernetes 통합을 사용하여 프로덕션에서 사용되는 컨테이너 이미지 및 해당 오픈 소스 또는 Linux 기반 패키지를 모니터링하여 알려진 취약점을 알리는 Snyk
* (모든 플랜) 프로덕션 통합이 없는 경우, [snyk monitor](../../snyk-cli/commands/monitor.md) CLI 명령을 사용하여 스냅샷을 찍고 제품으로 푸시되는 항목을 모니터링하세요.
