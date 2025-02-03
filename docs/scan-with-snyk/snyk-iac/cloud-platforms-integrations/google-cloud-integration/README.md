# Google 클라우드 통합

Snyk은 [Google Cloud](https://cloud.google.com/) 프로젝트와 통합하여 클라우드 구성에서 문제점을 찾고 취약성을 우선순위로 정할 수 있는 클라우드 컨텍스트를 생성합니다.

다음 방법을 사용하여 Google Cloud 계정을 Snyk에 등록할 수 있습니다:

* [Snyk 웹 UI](google-cloud-integration-web-ui/)
* [Snyk API](google-cloud-integration-api/)

## Google Cloud 통합 전제 조건

Google Cloud 통합을 추가하려면 다음이 필요합니다:

* Snyk 엔터프라이즈 [플랜](https://snyk.io/plans/)
* 적절한 기능 플래그가 Snyk 연락처에 의해 할당된 새로운 Snyk 조직
* Snyk 그룹 관리자 또는 조직 관리자 역할
* 읽기 전용 Google 서비스 계정을 생성할 권한이 있는 [Google Cloud](https://cloud.google.com/) 프로젝트 및 관련 자격 증명에 액세스 권한
* Snyk을 위해 Google 서비스 계정을 생성하기 위해 Google 자격 증명으로 구성된 [Terraform CLI](https://www.terraform.io/downloads) 액세스
* **API 전용:** Snyk API를 사용하는 조직 수준의 [서비스 계정](../../../../enterprise-setup/service-accounts/)으로 Org Admin 역할 필요
* **API 전용:** [curl](https://curl.se/), [HTTPie](https://httpie.io/), 또는 [Postman](https://www.postman.com/)과 같은 API 클라이언트
* **API 전용 (선택 사항)**: [jq](https://stedolan.github.io/jq/)로 서비스 계정 Terraform 템플릿을 포함하는 JSON을 변환하려는 경우
