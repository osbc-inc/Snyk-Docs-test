# 요청할 스코프

스코프는 사용자의 계정에서 특정 작업을 수행하는 데 Snyk 앱이 갖는 권한을 정의합니다. 사용자가 Snyk 계정에 대한 앱의 접근을 승인하면, 앱이 요청하는 스코프 목록을 보고 연결을 승인할지 여부를 결정할 수 있습니다.

Snyk 앱이 요청할 스코프를 결정할 때는 앱이 수행할 작업을 고려해야 합니다. 가능한 모든 스코프를 요청하는 것이 좋아 보일 수 있지만, 사용자들은 필요한 것보다 더 많은 권한을 요청하는 앱을 설치하기를 거부할 수 있습니다. 또한 사용자가 앱을 설치해도, 앱이 요청하는 스코프와 일치하는 모든 권한이 없으면 인증 프로세스를 완료할 수 없습니다.

다음은 **사용 가능한 스코프**를 나열합니다.

{% hint style="info" %}
`org.read`는 필수 스코프이며 항상 포함되어야 합니다.
{% endhint %}

| 스코프                       | 설명                                                                         |
| ----------------------------- | --------------------------------------------------------------------------- |
| org.read                    | 조직 정보 및 설정 보기                                                        |
| org.edit                    | 조직 정보 및 설정 편집                                                        |
| org.report.read             | 조직 내 보고서 보기                                                          |
| org.project.create          | 새로운 프로젝트 추가                                                          |
| org.project.read            | 프로젝트 정보 및 설정 보기 및 조직 대상 보기                                |
| org.project.edit            | 프로젝트 정보 편집                                                          |
| org.project.delete          | 프로젝트 영구 제거 및 조직 대상 영구 제거                                   |
| org.project.status          | 프로젝트 활성화 및 비활성화                                                   |
| org.project.test            | 프로젝트 테스트                                                               |
| org.project.ignore.create   | 새로운 프로젝트 무시 생성                                                    |
| org.project.ignore.read     | 프로젝트 무시 정보 보기                                                      |
| org.project.ignore.edit     | 프로젝트 무시 구성                                                          |
| org.project.ignore.delete   | 프로젝트 무시 영구 제거                                                     |
| org.project.tag.edit        | 프로젝트 태그 생성, 적용 및 제거                                             |
| org.project.pr.create       | 프로젝트용 수정 풀 요청 생성                                                 |
| org.project.pr.skip         | 보안 테스트 실패한 풀 요청을 성공적으로 표시하여 체크를 건너뛰기          |
| org.project.jira.issue.read | 지라 이슈 정보 보기                                                          |
| org.project.jira.issue.create | 새 지라 이슈 생성                                                          |
| org.package.test            | Snyk에서 지원하는 생태계에서 패키지 테스트                                |

{% hint style="info" %}
Snyk 앱을 만든 후에는 스코프를 업데이트할 수 없습니다. 앱 개발 프로세스 중에 필요한 스코프가 변경된 경우, 새로운 스코프 목록으로 새로운 Snyk 앱을 만들고 앱 설정의 `clientId` 및 `clientSecret`을 대체하세요. 이미 Snyk 앱을 설치한 사용자가 있다면, 사용자는 새로운 앱을 Snyk 계정과 연결해야 합니다.
{% endhint %}