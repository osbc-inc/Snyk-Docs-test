# Snyk는 무엇인가요?

Snyk는 코드, 오픈 소스 의존성, 컨테이너 이미지 및 인프라 코드 구성물의 보안 취약점을 스캔, 우선순위를 정하고 수정할 수 있게 해주는 플랫폼입니다. Snyk 플랫폼은 리스크 중심적 접근 방식을 사용하여 중요한 문제에 집중하고 의미 있는 영향을 미치지 않는 취약점의 소음을 제거합니다.

보안 프로그램을 관리하고 지배하기 위해 Snyk는 보안 팀에게 모든 애플리케이션 자산을 효과적으로 보여주는 즉각적인 시야, 대규모 환경에서 자동화하고 확장할 수 있는 스마트 정책, 보안 프로그램의 성능을 측정하기 위한 분석 및 보고 기능을 제공합니다.

* **Snyk Open Source**와 **Snyk Code**: [지원하는 언어 및 프레임워크](supported-languages-package-managers-and-frameworks/) 참조.
* **Snyk Container**: [지원하는 운영 체제 배포판](scan-with-snyk/snyk-container/how-snyk-container-works/operating-system-distributions-supported-by-snyk-container.md/) 참조.
* **{{Snyk Infrastructure as Code}}**: [지원하는 IaC와 클라우드 제공 업체](scan-with-snyk/snyk-iac/supported-iac-languages-cloud-providers-and-cloud-resources/) 참조.
* **{{Snyk AppRisk 오퍼링들}}**: {{Snyk AppRisk}} 참조.

## Snyk의 개발자 지향 접근

Snyk는 개발자의 워크플로우에서 가시성 및 실행 가능한 통찰을 제공합니다. 개발자를 보안 관행의 일환으로 참여시켜 보안 응용프로그램을 구축하는 데 중점을 두는 이점은 현재 개발자들이 Kubernetes와 Terraform과 같은 기술을 사용하여 모든 코드를 조합하고, 컨테이너에서 해당 코드를 실행하고, 인프라 코드 구성을 사용하여 배포한다는 점입니다.

각 구성 요소가 빌드되고 유지되는 곳마다 강력한 보안 프로세스가 수행됩니다. Snyk는 개발자가 선호하는 방식을 사용하면서 산업 최선의 가이드라인을 준수하고 지원합니다. Snyk는 직접 IDE, 워크플로 및 자동화 파이프라인에 통합하여 보안 전문성을 제공합니다.

<figure><img src=".gitbook/assets/image (565).png" alt="Snyk 개발자 보안 플랫폼: 제품 및 개발자 경험"><figcaption><p>Snyk 개발자 보안 플랫폼: 제품 및 개발자 경험</p></figcaption></figure>

## Snyk를 워크플로우에 활용하기

* **코드 보호**: [Snyk Open Source](https://docs.snyk.io/scan-using-snyk/snyk-open-source)를 사용하여 오픈 소스 의존성의 취약점을 수정하고 [Snyk Code](https://docs.snyk.io/scan-using-snyk/snyk-code)를 사용하여 소스 코드의 취약점을 수정합니다.
* **컨테이너 보호**: [Snyk Container](https://docs.snyk.io/scan-using-snyk/snyk-container)를 사용하여 컨테이너 이미지와 Kubernetes 애플리케이션의 취약점을 수정합니다.
* **인프라 보호**: [{{Snyk Infrastructure as Code}} (IaC)](https://docs.snyk.io/scan-using-snyk/snyk-iac/scan-your-iac-source-code)를 활용하여 Terraform, CloudFormation, Kubernetes 및 Azure 템플릿의 잘못된 구성을 수정합니다. [IaC+](https://docs.snyk.io/scan-using-snyk/snyk-iac/iac+-code-to-cloud-capabilities)를 사용하여 Amazon Web Services 계정, Microsoft Azure 구독 및 Google Cloud 프로젝트의 잘못된 구성을 수정합니다.

## Snyk를 실행하는 방법 선택하기

다음과 같은 방법으로 Snyk를 실행할 수 있습니다:

* [**Web**](getting-started/snyk-web-ui.md): Snyk 웹 UI([app.snyk.io](https://app.snyk.io))를 통해 구성 설정, 발견된 문제 필터링 및 수정, 리포트 등의 기능을 제공하는 웹 기반 경험.
* [**CLI**](snyk-cli/): Snyk 명령줄 인터페이스를 사용하여 로컬 머신에서 취약점 스캔을 실행하고 Snyk를 파이프라인에 통합합니다.
* [**IDEs**](scm-ide-and-ci-cd-integrations/snyk-ide-plugins-and-extensions/): Snyk IDE 통합을 사용하여 개발 환경에 Snyk를 포함시킬 수 있습니다.
* [**API**](snyk-api/): Snyk API를 통해 Snyk를 프로그래밍 방식으로 통합하여 Snyk 보안 자동화를 특정 워크플로에 맞게 조정할 수 있습니다.

## Snyk는 무엇과 통합할 수 있나요?

소프트웨어 개발 프로세스에 대한 Snyk 통합은 소스 제어, IDE, CI/CD 및 기타 여러 가지 개발 및 보안 프로세스에 Snyk를 통합할 수 있게 합니다.

자세한 내용은 [Snyk와 통합하기](integrate-with-snyk/)를 참조하십시오.

## **Snyk 비용은 어떻게 되나요?**

Snyk에는 무료부터 엔터프라이즈까지 다양한 요금제가 있습니다. [Snyk 요금제](https://snyk.io/plans/)를 확인하세요.

Snyk는 플랫폼의 평가판을 제공하지만 특정 기능 제한이 있습니다. [평가판 제한 사항](https://docs.snyk.io/implement-snyk/enterprise-implementation-guide/trial-limitations/)을 확인하세요.

## 내 데이터는 어떻게 처리되나요?

Snyk 데이터 처리에 대한 자세한 내용은 [Snyk가 데이터를 처리하는 방법](working-with-snyk/how-snyk-handles-your-data.md)을 참조하십시오.