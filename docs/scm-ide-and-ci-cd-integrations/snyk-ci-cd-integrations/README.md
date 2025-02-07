# Snyk CI/CD 통합

{% hint style="info" %}
Snyk은 다음과 같은 이유로 CI/CD 통합을 위해 [Snyk CLI](https://github.com/snyk/cli)를 사용할 것을 권장합니다:

* [미리보기 채널](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#preview)을 사용하여 CLI의 진행 중인 기능을 테스트할 유연성을 갖게 됩니다.
* CLI는 규칙적인 주기로 [안정된 릴리스](../../snyk-cli/releases-and-channels-for-the-snyk-cli.md#stable)를 제공합니다.
* CLI를 통해 대규모로 Snyk을 배포하는 경우 [사용 사례를 확장할 수 있는 옵션](../../snyk-cli/scan-and-maintain-projects-using-the-cli/cli-tools/)을 제공합니다.
{% endhint %}

[CI/CD 통합을 사용하기로 결정하면,](../git-repository-and-ci-cd-integrations-comparisons.md) 일반적으로 통합을 단계적으로 채택하게 됩니다. 배포 방법을 선택하고 검사할 코드에 대한 전략을 구현합니다. [Snyk CI/CD 통합 배포 및 전략](snyk-ci-cd-integration-deployment-and-strategies/)을 참조하세요.

자세한 정보는 사용 중인 통합 페이지를 참조할 수 있습니다:

* [AWS CodePipeline 통합](aws-codepipeline-integration-by-adding-a-snyk-scan-stage/)
* [Azure 파이프라인 통합](azure-pipelines-integration/)
* [Bitbucket 파이프라인 통합](bitbucket-pipelines-integration-using-a-snyk-pipe/)
* [CircleCI 통합](circleci-integration-using-a-snyk-orb.md)
* [GitHub Actions 통합](github-actions-for-snyk-setup-and-checking-for-vulnerabilities/)
* [Jenkins 통합](jenkins-plugin-integration-with-snyk.md)
* [Maven 통합](maven-plugin-integration-with-snyk.md)
* [TeamCity 통합](teamcity-jetbrains-integration-using-the-snyk-security-plugin/)
* [IaC용 Terraform Cloud 통합](terraform-cloud-integration-for-snyk-iac-using-run-tasks/)
* [IaC용 Terraform Enterprise 통합](terraform-enterprise-integration-for-snyk-iac.md)

GitLab 파이프라인 통합의 경우, [파이프라인 설정](https://github.com/snyk-labs/snyk-cicd-integration-examples/blob/master/GitLabCICD/gitlab-npm.yml)을 참조하십시오.

CI/CD를 위한 바이너리 및 npm 통합의 추가 예제는 [GitHub CI/CD 예제](https://github.com/snyk-labs/snyk-cicd-integration-examples)를 참조하십시오.
