# 도구: jira-tickets-for-new-vulns

`jira-tickets-for-new-vulns`은 Snyk에서 모니터된 프로젝트를 동기화하고 새로운 문제 및 이미 생성되지 않은 이슈에 대한 Jira 티켓을 자동으로 열 수 있도록 제공합니다.

매 X분/시간마다 `Cron`을 실행하여 이슈를 해결하세요. 이 도구는 정기적으로 실행되거나 선택한 트리거(웹훅)로 실행됩니다.

## 설치

[릴리스 페이지](https://github.com/snyk-tech-services/jira-tickets-for-new-vulns/releases)에서 이진 파일을 사용합니다.

## 사용법 - 빠른 시작

```
./snyk-jira-sync-<yourplatform> 
    -orgID=<SNYK_ORG_ID>                     // 설정 아래에서 찾을 수 있음
    -token=<API Token>                       // Snyk API 토큰. 서비스 계정이 작동합니다.
    -jiraProjectKey=<키>                     // Jira 티켓이 열릴 프로젝트 키
```

**확장된 옵션**

```
./snyk-jira-sync-<yourplatform> 
    --orgID=<SNYK_ORG_ID>                                                // 설정 아래에서 찾을 수 있음
    --projectID=<SNYK_PROJECT_ID>                                        // 선택 사항. 제공하지 않으면 조직의 모든 프로젝트를 동기화합니다.
                                                                        // 프로젝트 ID는 project->settings에서 찾을 수 있습니다.
    --api=<API 엔드포인트>                                              // 선택 사항. 사설 인스턴스의 경우 https://<인스턴스>/api로 설정
    --token=<API 토큰>                                                   // Snyk API 토큰. 서비스 계정이 작동합니다.
    --jiraProjectID=<12345>                                              // Jira 티켓이 열릴 프로젝트 ID
    --jiraProjectKey=<키>                                                // Jira 티켓이 열릴 프로젝트 키
    --jiraTicketType=<작업|버그|....>                                     // 선택 사항. 열릴 티켓 유형. 기본값은 버그입니다. 아래의 'Notes' 섹션을 참조하세요.
    --severity=<critical|high|medium|low>                                // 선택 사항. 티켓을 열기 위한 심각도 임계값. 기본값은 low입니다.
    --maturityFilter=[mature,proof-of-concept,no-known-exploit,no-data]  // 선택 사항. 성숙도 수준만 포함합니다. 쉼표로 구분됩니다.
    --type=<all|vuln|license>                                            // 선택 사항. 티켓을 열기 위한 이슈 유형. 기본값은 모두입니다.
    --assigneeId=<123abc456def789>                                       // 선택 사항. 티켓을 할당할 Jira 사용자 ID입니다. 참고: assigneeName과 assigneeId는 동시에 사용하지 마십시오.
    --assigneeName=<계정이름>                                            // 선택 사항. 티켓을 할당할 Jira 사용자 이름입니다. 참고: assigneeName과 assigneeId는 동시에 사용하지 마십시오.
    --priorityIsSeverity                                                 // 선택 사항. 심각도에 따라 티켓 우선 순위를 설정합니다 (기본값: Low|Medium|High|Critical=>Low|Medium|High|Highest)
    --labels=<이슈 레이블1>,이슈 레이블2                                 // 선택 사항. JIRA 티켓 레이블을 설정합니다
    --priorityScoreThreshold=[0-1000]                                    // 선택 사항. 최소 우선 순위 점수 임계값
    --dryRun=<true|false>                                                // 선택 사항. 결과는 도구가 실행되는 위치에 json 파일에서 찾을 수 있습니다
    --debug=<true|false>                                                 // 선택 사항. 디버그 모드를 활성화합니다
    --ifUpgradeAvailableOnly=<true|false>                                // 선택 사항. 업그레이드 가능한 문제에 대해서만 티켓을 생성합니다
    --configFile                                                         // 루트가 아닌 경우 jira.yaml 경로
```

## 제약 사항

이 도구는 인프라 구성 프로젝트를 지원하지 않습니다. 코드 및 오픈 소스 프로젝트에 대해서만 문제를 엽니다. 다른 프로젝트 유형은 무시됩니다.

**우선 순위는 심각도입니다**

문제의 Jira 티켓 우선 순위를 설정할 수 있는 옵션이 있습니다. 기본값은 다음과 같습니다:

| 문제 심각도 | Jira 우선 순위 |
| -------------- | ------------- |
| critical       | Highest       |
| high           | High          |
| medium         | Medium        |
| low            | Low           |

기본값을 무시하고 값을 설정하려면 `SNYK_JIRA_PRIORITY_FOR_XXX_VULN` 환경 변수를 사용하세요.

> 예: 심각한 심각성은 Jira에서 Hot Fix 우선 순위를 받아야 합니다.
>
> export SNYK\_JIRA\_PRIORITY\_FOR\_CRITICAL\_VULN='Hot Fix'

## 소스로부터 설치

리포지토리를 클론하고 빌드하세요:

> `go run main.go jira.go jira_utils.go vulns.go snyk.go snyk_utils.go`

**문제를 보고하세요.**

## 의존성

JSON 파싱을 쉽게 만들어주는 https://github.com/michael-go/go-jsn/jsn github.com/tidwall/sjson github.com/kentaro-m/blackfriday-confluence gopkg.in/russross/blackfriday.v2

## 로그 파일

도구가 실행된 위치에 만들어진 모든 생성된 티켓이 나열된 로그 파일이 있습니다.

```
{
  "projects": {
    "123": [
      {
        "Summary": "test/goof:package.json - 원격 코드 실행 (RCE)",
        "Description": "\r\n \\*\\*\\*\\* 문제 세부 정보: \\*\\*\\*\\*\n\r\n cvss 점수:  8.10\n exploit 성숙도:  proof\\-of\\-concept\n 심각성:  높음\n 패키지 버전: 3.0.0\\]\n\r\n*영향 받은 경로:*\n\\- \"snyk\"@\"1.228.3\" =\u003e \"proxy\\-agent\"@\"3.1.0\" =\u003e \"pac\\-proxy\\-agent\"@\"3.0.0\" =\u003e \"pac\\-resolver\"@\"3.0.0\"\n\r\n[Snyk에서이 문제 보기|https://app.snyk.io/org/test/project/123]\n\n[이 문제에 대해 자세히 알아보기|https://snyk.io/vuln/SNYK-JS-PACRESOLVER-1589857]\n\n",
        "JiraIssueDetail": {
          "JiraIssue": {
            "Id": "10001",
            "Key": "FPI-001"
          },
          "IssueId": "SNYK-JS-PACRESOLVER-1589857"
        }
      },
      {
        "Summary": "test/goof:package.json - 프로토타입 오염",
        "Description": "\r\n \\*\\*\\*\\* 문제 세부 정보: \\*\\*\\*\\*\n\r\n cvss 점수:  6.30\n exploit 성숙도:  proof\\-of\\-concept\n 심각성:  중간\n 패키지 버전: 4.2.0\\]\n\r\n*영향 받은 경로:*\n\\- \"snyk\"@\"1.228.3\" =\u003e \"configstore\"@\"3.1.2\" =\u003e \"dot\\-prop\"@\"4.2.0\"\n\r\\- \"snyk\"@\"1.228.3\" =\u003e \"update\\-notifier\"@\"2.5.0\" =\u003e \"configstore\"@\"3.1.2\" =\u003e \"dot\\-prop\"@\"4.2.0\"\n\r\n[Snyk에서이 문제 보기|https://app.snyk.io/org/test/project/123]\n\n[이 문제에 대해 자세히 알아보기|https://snyk.io/vuln/SNYK-JS-DOTPROP-543499]\n\n",
        "JiraIssueDetail": {
          "JiraIssue": {
            "Id": "10001",
            "Key": "FPI-001"
          },
          "IssueId": "SNYK-JS-DOTPROP-543499"
        }
      },
    ]
  }
}
```

## Jira.yaml

다음은 구성 파일 구조의 예시입니다. Jira 프로젝트에 사용자 정의 필수 필드가 구성된 경우 이러한 필드를 구성 파일에 추가해야 합니다. 사용자 정의 필드가 예상하는 키와 값 모두를 구성 파일의 customMandatoryFields 키에 추가하십시오:

```
schema: 1
snyk: 
    orgID: a1b2c3de-99b1-4f3f-bfdb-6ee4b4990513 # <SNYK_ORG_ID> 
    projectID: a1b2c3de-99b1-4f3f-bfdb-6ee4b4990514 # <SNYK_PROJECT_ID>
    severity: critical # <critical|high|medium|low>
    maturityFilter: mature # <mature,proof-of-concept,no-known-exploit,no-data>
    type: all # <all|vuln|license>
    priorityScoreThreshold: 10
    api: https://myapi # <API endpoint> 기본값
    ifUpgradeAvailableOnly: false # <true|false>
jira:
    jiraTicketType: 작업 # <Task|Bug|....>
    jiraProjectID: 12345
    assigneeId: 123abc456def789
    assigneeName: AccountName
    priorityIsSeverity: true # <true|false>
    label: label1 # <IssueLabel1>,<IssueLabel2>
    jiraProjectKey: testProject
    priorityIsSeverity: false # <true|false> (기본값: Low|Medium|High|Critical=>Low|Medium|High|Highest)
    customMandatoryFields:
        key: 
            value: 5
```

**참고:**

- 토큰은 구성 파일에 포함되어 있지 않아야 합니다.
- 명령줄 인수가 구성 파일을 재정의합니다. 예: 위의 구성 파일을 사용하여 `./snyk-jira-sync-macOs -Org=1234 -configFile=true -token=123`를 실행하면 도구가 사용하는 조직 ID는 a1b2c3de-99b1-4f3f-bfdb-6ee4b4990513이 아니라 1234가 됩니다.
- 기본값은 '확장된 옵션'을 참조하세요.
- Jira 설정에서 구성된 문제 유형을 사용했는지 확인하십시오. 기본값은 Bug입니다. 사용 중인(또는 기본값) 유형이 Jira 설정에 있는지 확인하세요.