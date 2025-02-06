# GitHub - Flow 및 기술

## Flow <a href="#flow" id="flow"></a>

1. SCM에서 액세스 권한이 있는 조직을 가져와 orgs 목록을 작성합니다.
2. 가져온/제공된 orgs에 속하는 repo 중 `one`/`all`를 가져옵니다.
3. 가져온/제공된 repo/s의 커밋을 가져와 기여자 목록을 작성합니다.
4. 기여자별로 repo의 커밋 수를 계산합니다.
5. 제외 파일(exclusion file)에 지정된 기여자들을 제거합니다 (`the exclusionFilePath` 플래그가 설정되고 유효한 텍스트 파일 경로가 제공된 경우).
6. 결과를 출력합니다.

## 사용된 GitHub API 엔드포인트 <a href="#azure-api-endpoints-used" id="azure-api-endpoints-used"></a>

* GitHub에서 orgs를 가져오는 데 사용되는 엔드포인트: `/user/orgs`
* 가져온/제공된 orgs 목록과 관련된 repo 목록을 얻는 데 사용되는 엔드포인트: `/orgs/{Org}/repos`
* 가져온/제공된 repo 목록의 커밋을 가져오는 데 사용되는 엔드포인트: `repos/{Org}/{Repo}/commits?since={threeMonthsDate}`
