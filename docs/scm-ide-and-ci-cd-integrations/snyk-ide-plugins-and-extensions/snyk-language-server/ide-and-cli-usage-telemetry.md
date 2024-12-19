# IDE 및 CLI 사용 텔레메트리

Snyk Language Server는 각 성공적인 테스트 후에, Snyk의 소유 분석 서비스를 통해 사용 텔레메트리를 수집합니다. 데이터는 제3자 서비스를 거치지 않습니다.

요청 본문에는 여러 카테고리의 데이터가 포함됩니다:

- **통합의 버전 및 유형**, 예를 들어 Language Server 버전, IDE 유형 및 버전, 플러그인 유형 및 버전
- **로컬 환경**: 운영 체제, 아키텍처
- **프로젝트 데이터**, 예를 들어 리포지토리 git URL, 브랜치 이름
- **성능 데이터**: 스캔 소요 시간(밀리초)
- **이벤트 세부 정보**: 이벤트 유형, 타임스탬프, 상태, 발견된 수

API 경로에는 고유한 Snyk 조직 ID가 포함됩니다. 따라서 데이터에는 조직 정보가 주석으로 달려 있습니다.

사용자의 개인 식별 가능한 정보(PII)는 요청의 일부가 아닙니다. Snyk 내부 분석 서비스 API는 인증 토큰이 있어야만 합니다. 들어오는 이벤트 데이터는 백엔드 서비스 내에서 사용자 데이터와 병합되어, Snyk가 특별 보고서 페이지를 작성할 수 있게 합니다: [개발자 IDE 및 CLI 사용](../../../manage-issues/reporting/available-snyk-reports.md#developer-ide-and-cli-usage).

GitHub 리포지토리에서 오픈 소스 로직을 확인하고 분석할 수 있습니다: [snyk/go-application-framework](https://github.com/snyk/go-application-framework/blob/main/internal/api/clients/analytics\_v20240307.go). 이것은 보안 감사가 필요할 때 유용합니다.