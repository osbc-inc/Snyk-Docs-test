# Customize PR templates

당신이 [SCM 통합](../../../../scm-ide-and-ci-cd-integrations/snyk-scm-integrations/)을 이용해 가져온 Snyk Open Source와 Snyk Container 프로젝트를 수정하거나 업그레이드할 때, Snyk는 당신의 저장소에 대해 풀 리퀘스트(PR)를 생성합니다.

Snyk는 제목, 설명 및 커밋 메시지에 대한 기본 PR 템플릿을 가지고 있습니다. 이들은 무엇이 변경되는지, 어떤 이슈가 해결되는지 등을 보여줍니다.

당신은 풀 리퀘스트를 제출하는 데 자체 표준 및 관행을 가질 수 있습니다. 예를 들어, 만약 Snyk로부터의 풀 리퀘스트인 경우, 제목이 `SNYK:`로 시작되도록 요구할 수 있습니다.

다음의 방법으로 사용자 지정 PR 템플릿을 설정할 수 있습니다:

* API 요청 - 그룹 내 모든 PR을 업로드된 사용자 지정 템플릿으로 설정합니다. [사용자 정의 템플릿을 사용한 API 요청 생성](apply-a-custom-pr-template.md#create-an-api-request-with-a-custom-template)을 참조하세요.
* [YAML 업로드 ](apply-a-custom-pr-template.md#customize-using-a-yaml-pr-template-file)- 특정 저장소에 대한 사용자 지정 템플릿을 설정합니다.

템플릿이 설정된 후에는 사용자 지정 PR 템플릿 기능이 활성화됩니다.

다음과 같은 **우선 순위 순서**로 사용자 지정 템플릿이 적용됩니다:

* 먼저 Snyk는 특정 저장소에 대해 설정된 것, 즉 YAML 업로드를 적용합니다.
* 해당 템플릿에서 누락된 부분이 있거나 프로젝트에 YAML 파일이 존재하지 않으면, Snyk는 API 요청을 통해 설정된 그룹 템플릿을 확인합니다.
* 그룹 템플릿에 누락된 것이 있거나 사용자 지정 템플릿을 찾을 수 없는 경우, Snyk는 기본 템플릿을 적용합니다.
