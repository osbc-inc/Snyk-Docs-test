---
description: 각 SCM 스크립트의 플래그
---

# 플래그

| SCM             | 자격 증명           | 조직             | 저장소        | 제외 파일 경로         | Json     | Snyk 모니터링된 저장소 건너뛰기 | 가져오기 파일 폴더 경로 | 가져오기 파일의 저장소 유형 | 추가적인 플래그                                |
| --------------- | ------------------- | --------------- | ----------- | --------------------- | -------- | -------------------------- | ----------------------- | ------------------------- | ----------------------------------------------- |
| **Azure**       | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    | `"org" [required]`                              |
| **BB Cloud**    | `"user"/"password"` | `"workspaces"`  | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    |                                                 |
| **BB Server**   | `"token"`           | `"projectKeys"` | `"repo"`    | `"exclusionFilePath"` | `"json"` | `"skipSnykMonitoredRepos"` | `"importConfDir"`       | `"importFileRepoType"`    | `"url" [required]`                              |
| **GitHub**      | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"` | `"json"` |                            |                         |                           |                                                 |
| **GitHub Ent.** | `"token"`           | `"orgs"`        | `"repo"`    | `"exclusionFilePath"` | `"json"` |                            |                         |                           | `"url" [required], "fetchAllOrgs" [optional]**` |
| **GitLab**      | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"` | `"json"` |                            |                         |                           |                                                 |
| **GitLab Ent.** | `"token"`           | `"groups"`      | `"project"` | `"exclusionFilePath"` | `"json"` |                            |                         |                           | `"url" [required]`                              |

{% hint style="info" %}
플래그 이름은 SCM 내의 해당 이름과 일치합니다.\
Bitbucket Server, GitHub Enterprise 및 GitLab Enterprise와 같은 "비공개" SCM은 명령에 플래그로 설정된 호스트 URL이 필요합니다.
{% endhint %}

{% hint style="info" %}
`"fetchAllOrgs"` 플래그는 GitHub Enterprise에 고유하며 GHE의 Orgs에 대한 두 가지 액세스 레벨을 구별합니다:\
1\. 플래그 있음 - 제공된 토큰이 액세스 권한을 갖는 모든 org를 가져옴.\
2\. 플래그 없음 - 토큰을 제공한 **사용자**가 일부 작업 권한(예: 읽기 권한 등)을 가지고 있는 org만 가져옴.
{% endhint %}

{% hint style="info" %}
snyk-api-import 도구를 사용하여 생성된 가져오기 파일을 사용하는 방법에 대한 자세한 내용은 [생성 및 사용 중인 가져오기 파일](creating-and-using-the-import-file.md)을 참조하십시오.
{% endhint %}
