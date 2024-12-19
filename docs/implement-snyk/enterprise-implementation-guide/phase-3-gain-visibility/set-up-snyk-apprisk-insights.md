# Snyk AppRisk Insights 설정하기

## 위험 요소에 따라 문제 우선순위 설정

리스크 기반 우선순위 및 Snyk AppRisk를 위한 Kubernetes 커넥터를 설정하고 배포하는 방법에 대해 자세히 알아보세요. 이는 Snyk AppRisk 메뉴의 [**Issues**](../../../getting-started/snyk-web-ui.md#view-and-prioritize-issues) 페이지를 사용합니다.

리스크 기반 우선순위 또는 인사이트는 Snyk AppRisk가 당신의 애플리케이션의 컨텍스트를 이해하고 보안 문제를 더 잘 우선순위를 매기는 데 도움을 주는 기능입니다.

[Snyk 리스크 기반 우선순위](../../../manage-risk/prioritize-issues-for-fixing/#prioritization-based-on-risk) 제품은 다음과 같은 여러 가지 리스크 요소에 중점을 두어 취약점을 분석합니다:

- **배포됨(Deployed)**: 내 코드가 배포된 컨테이너 이미지에 있는가?
- **운영 체제 조건(OS condition)**: 이 취약점이 내 운영 체제에 적용되는가?
- **공개 지향(Public facing)**: 내 컨테이너 이미지가 인터넷으로 향한 구성된 경로가 있는가?
- **로드된 패키지(Loaded package)**: 이미지의 종속성인 써드파티 패키지가 로드되었는가?

Snyk 리스크 기반 우선순위의 목표는 애플리케이션이 어떻게 배포되고 구성되었는지를 이해하여 여러분의 애플리케이션에 대한 위험으로 보이는 문제들을 기반으로 우선순위를 매기도록 도와주는 것입니다.

Snyk 리스크 기반 우선순위 작업 방식에 대해 더 많은 세부 정보와 핵심 개념에 대한 이해를 얻으려면 [Assets](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#assets) 및 [Risk factors](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/#risk-factors) 페이지에 초점을 맞춘 [작동 방식](../../../manage-risk/prioritize-issues-for-fixing/assets-and-risk-factors-for-snyk-apprisk/) 페이지를 참조하세요.

이 수준의 목적은 Snyk 컨테이너 문제 또는 취약점에 컨텍스트를 제공하는 것입니다. 클러스터에 Kubernetes 커넥터를 배포한 후, 컨테이너가 배포되었고 공개적으로 노출되었는지 식별할 수 있게 됩니다. 이는 컨테이너 취약점을 우선순위를 매길 수 있도록 도와줍니다.

Kubernetes 커넥터를 설정하는 방법에 대해 더 자세히 알아보려면 [우선순위 설정: Kubernetes 커넥터](../../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-kubernetes-connector.md) 페이지를 참조하세요.

## 소스 코드를 컨테이너 이미지와 연결시키기

Snyk Code 및 Open Source 프로젝트를 컨테이너 프로젝트와 연관시킴으로써 리스크 기반 우선순위의 전체 기능을 활용하세요. 이는 컨텍스트 기반 위험 요소를 사용하여 Snyk Open Source 및 Snyk Code 문제에 우선순위를 지정할 수 있게 해줍니다.

선택한 애플리케이션에 필요한 연결 설정을 수행하는 방법에 대해 알아보려면 [우선순위 설정: Snyk Open Source, Code, 그리고 Container 프로젝트 연결](../../../manage-risk/prioritize-issues-for-fixing/set-up-insights-for-snyk-apprisk/set-up-insights-associating-snyk-open-source-code-and-container-projects.md) 페이지를 확인하세요.