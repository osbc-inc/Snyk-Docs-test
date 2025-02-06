# BitBucket Cloud - Flow 및 기술

## Flow

1. Snyk로부터 모니터링 중인 프로젝트들을 가져옵니다 (만약 `skipSnykMonitoredRepos` 플래그가 **설정되어 있지 않고** `SNYK_TOKEN`이 export되었을 경우).
2. 자격 증명이 액세스할 수 있는 `one`/`some`/`all` 작업 공간을 SCM에서 가져와 작업 공간 목록을 생성합니다.
3. 가져온/제공된 작업 공간 아래의 `한 개`/`모든` 레포지토리를 가져옵니다.
4. Snyk에 의해 모니터링되지 않는 레포지토리들을 제거하고 Repo 목록을 만듭니다 (만약 `skipSnykMonitoredRepos` 플래그가 **설정되어 있지 않고** `SNYK_TOKEN`이 export되었을 경우).
5. 모니터링되지 않는 레포지토리들을 쉽게 가져올 수 있도록 Snyk 계정으로 가져올 레포지토리들을 위한 가져오기 파일을 생성합니다 (만약 `importConfDir` 플래그가 설정된 경우)
6. 가져온/제공된 레포지토리들에 대한 커밋들을 가져와 기여자 목록을 작성합니다.
7. 커밋 수를 기여자별로 계산합니다.
8. 제외 파일에 지정된 기여자들을 제거합니다 (만약 `exclusionFilePath` 플래그가 설정되어 있고 유효한 텍스트 파일 경로가 제공된 경우).
9. 결과를 출력합니다.

## 사용된 Bitbucket Cloud API 엔드포인트

* BB Cloud에서 레포지토리를 가져오기 위해, 작업 공간이 **제공되지 않았을 경우**: `/api/2.0/repositories`
* BB Cloud에서 레포지토리를 가져오기 위해, 작업 공간/들이 **제공되었을 경우**: `/api/2.0/repositories/{Workspace}`
* 가져온/제공된 레포지토리 목록에 대한 커밋들을 가져오기 위해: `/api/2.0/repositories/{Workspace}/{Repo}/commits`
