# Azure - Flow 및 기술

## Flow

1. Snyk에서 모니터링 중인 프로젝트를 가져옵니다 (`skipSnykMonitoredRepos` 플래그가 **설정되지 않았으며** `SNYK_TOKEN`이 내보내기된 경우).
2. 자격 증명이 액세스 할 수 있는 프로젝트 중 `one`/`some`/`all`를 SCM에서 가져와 프로젝트 목록을 생성합니다.
3. `one`/`all` 프로젝트 아래의 하나/모든 repo를 가져옵니다.
4. Snyk에 의해 모니터링되지 않는 리포지토리를 제거하고 Repo 목록을 생성합니다 (`skipSnykMonitoredRepos` 플래그가 **설정되지 않았으며** `SNYK_TOKEN`이 내보내기된 경우).
5. Unmonitored repos를 임포트하는 데 사용할 수 있도록 가져온 리포지토리에 대한 임포트 파일을 생성합니다 (`importConfDir` 플래그가 설정된 경우).
6. 가져온/제공된 repo/s의 커밋을 가져와 기여자 목록을 만듭니다.
7. 기여자별로 리포지토리의 커밋 수를 계산합니다.
8. 제외 파일에서 지정된 기여자를 제거합니다 (`exclusionFilePath` 플래그가 설정되고 유효한 텍스트 파일 경로가 제공된 경우).
9. 결과를 출력합니다.

## Azure API endpoints 사용

* Azure에서 프로젝트 가져오기: `{Org}/_apis/projects`
* 가져온/제공된 프로젝트 목록과 관련된 리포지토리 목록 가져오기: `{Project}/_apis/git/repositories`
* 가져온/제공된 repo 목록의 커밋 가져오기: `{Project}/_apis/git/repositories/{Repo}/commits?$searchCriteria.fromDate={ThreeMonthsDate}`
