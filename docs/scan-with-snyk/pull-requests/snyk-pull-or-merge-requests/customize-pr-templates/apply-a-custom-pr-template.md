# 사용자 지정 PR 템플릿 적용

## API를 사용하여 사용자 정의 PR 템플릿 생성 및 관리하기

API 엔드포인트 [그룹에 대한 풀 리퀘스트 템플릿 만들기 또는 업데이트](https://apidocs.snyk.io/?#post-/groups/-group_id-/settings/pull_request_template)을 사용하여 사용자 정의 PR 템플릿을 생성할 수 있습니다. 사용자 정의 속성을 포함하는 JSON 페이로드가 포함된 API 요청을 보냅니다. 이 요청은 조직 또는 프로젝트 내의 모든 프로젝트에서 사용될 그룹 수준의 풀 리퀘스트 템플릿을 구성합니다. Snyk API를 사용하여 생성된 풀 리퀘스트 템플릿은 언제든지 업데이트할 수 있으며 그룹의 모든 프로젝트는 최신 변경사항으로 자동 업데이트됩니다.

PR 템플릿의 API 구성은 그룹 수준에서만 가능합니다.

API 요청을 사용하여 사용자 정의 템플릿을 업로드하면 해당 그룹의 모든 Snyk PR이 이 형식을 채택하게 되어 사용자 정의 가능한 속성에 대한 기본 Snyk 템플릿이 비활성화됩니다. 문자열만 허용되며 목록과 숫자는 허용되지 않습니다.

사용자 정의 템플릿에서 누락된 사용자 정의 가능한 속성이 있으면 PR을 열 때 Snyk는 이러한 속성에 대한 기본값으로 되돌아갑니다.

다음 **API를 사용하여 사용자 정의할 수 있는 속성**은 다음과 같습니다:

* `title` - PR 제목 사용자 정의
* `commit_message` - PR 커밋 메시지 사용자 정의
* `description` - PR 설명 사용자 정의

PR에 대해 브랜치 이름을 사용자 정의할 수는 없습니다. PR의 브랜치 이름은 Snyk의 기본값을 사용합니다.

그룹의 사용자 정의 PR 템플릿을 검색하려면 엔드포인트 [그룹에 대한 풀리퀘스트 받기](https://apidocs.snyk.io/?#get-/groups/-group_id-/settings/pull_request_template) 사용하세요. 이것은 템플릿을 변경하거나 문제해결할 때 유용합니다.

템플릿을 삭제하려면 [그룹에 대한 풀리퀘스트 템플릿 삭제](https://apidocs.snyk.io/?#delete-/groups/-group_id-/settings/pull_request_template) 엔드포인트를 사용하세요.

## YAML PR 템플릿 파일을 사용자 정의하기

### YAML 파일 생성하기

[머스터치(mustache) 문법](https://mustache.github.io)을 사용하여 YAML 템플릿을 수동으로 생성하고 파일을 프로젝트나 저장소에 추가하세요.

프로젝트에 사용자 정의 템플릿이 업로드되면, 해당 프로젝트의 Snyk PR은 이 형식을 채택하여 사용자 정의 속성을 위한 기본 Snyk 템플릿을 비활성화합니다. 문자열만 허용되며 목록과 숫자는 허용되지 않습니다. 템플릿에서 누락된 사용자 정의 가능한 속성이 있으면 PR을 열 때 Snyk는 이러한 속성에 대한 기본값으로 되돌아갑니다.

#### YAML 여러 줄 연산자

YAML 여러 줄 연산자를 사용할 수 있습니다. 다음 형식을 따라 여러 줄에 걸쳐 자세한 설명을 만들 수 있습니다:

```yaml
description: |
  This pull request comes from Snyk
  For more information see the project page 
  If you have more questions reach out to a member of the security team

```

파이프 연산자는 새 줄 문자를 유지합니다. 더하기 기호, `>`,를 사용하여 모든 줄을 끝에 새 줄 문자와 함께 공백으로 연결합니다. 콜론을 사용하려면 여러 줄 연산자 `|` 또는 `>`를 사용하거나 줄을 이중 인용부호로 묶어야 합니다:

```yaml
commitMessage: "snyk: this is a security pull request"
```

#### YAML을 위한 사용자 정의 가능한 속성

다음 속성들은 사용자 정의가 가능합니다:

* `title` - PR 제목 사용자 정의
* `commitMessage` - PR 커밋 메시지 사용자 정의
* `description` - PR 설명 사용자 정의

PR에 대해 브랜치 이름을 사용자 정의할 수는 없습니다. PR의 브랜치 이름은 Snyk의 기본값을 사용합니다.

### YAML 사용자 정의 PR 템플릿 사용하기

프로젝트(리포지토리)에 snyk\_pull\_request\_template.yaml이라는 이름의 YAML 파일을 수동으로 업로드할 수 있습니다. 이 통합의 종류에 따라 방법이 다릅니다.

* GitHub/ GitHub Enterprise - `/.github/snyk_pull_request_template.yaml`
* GitLab - `/.gitlab/snyk_pull_request_template.yaml`
* Azure DevOps - `/.azuredevops/snyk_pull_request_template.yaml`
* 기타 (예: BitBucket) - `/.config/snyk_pull_request_template.yaml`

여러 리포지토리에 대해 사용자 정의 템플릿을 사용하려면 각 리포지토리에 YAML 사용자 정의 템플릿 파일을 추가하세요.

## 사용자 정의 PR 템플릿 가져오는 브로커 구성

만약 [Snyk 브로커](../../../../enterprise-setup/snyk-broker/)를 사용 중이라면, 버전이 4.188.0 이상인 브로커를 사용하고 `ACCEPT_CUSTOM_PR_TEMPLATES` 환경 변수를 사용하여 브로커가 사용자 정의 PR 템플릿을 가져올 수 있도록 설정해야 합니다.

이를 위해 `ACCEPT=/path/to/custom.json`을 제거하고 다음 환경 변수를 브로커 컨테이너나 배포에 추가해야 합니다:

```
ACCEPT_CUSTOM_PR_TEMPLATES=true
```
