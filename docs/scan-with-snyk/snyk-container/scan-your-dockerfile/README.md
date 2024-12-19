# Dockerfile을 스캔하세요

Snyk Container를 사용하면 Dockerfile을 분석하고 Dockerfile에서의 베이스 이미지를 스캔할 수 있습니다.

## Dockerfile 분석을 위한 전제 조건

Dockerfile 분석을 사용하기 전에 다음 사항을 확인해야 합니다:

* Snyk를 통합한 계정에 해당 Dockerfile 저장소가 포함되어 있어야 합니다.
* Dockerfile을 포함한 Git 저장소를 위한 통합을 구성해야 합니다. 자세한 정보는 [Dockerfile 분석을 위한 지원되는 소스 코드 관리자(SCM)](./#supported-scms-for-dockerfile-analysis)를 참조하십시오.
* 이미 모니터링을 위해 관련 컨테이너 프로젝트를 가져왔어야 합니다.

## Dockerfile 분석을 위한 지원되는 소스 코드 관리자(SCM)

Snyk는 다음 서비스에 대해 Dockerfile 분석을 지원합니다:

* GitHub
* GitHub Enterprise
* GitLab
* Bitbucket Cloud
* Bitbucket Cloud App
* Bitbucket Server
* Azure Repos