# API 요청 속도 제한 제어

## Azure DevOps

Azure DevOps는 고유한 "TSTU" 개념을 사용하여 API 호출 속도를 제한하는 방법을 가지고 있습니다. 자세한 내용은 [가이드](https://docs.microsoft.com/en-us/azure/devops/integrate/concepts/rate-limits?view=azure-devops)에서 설명되어 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 매 초 최대 두 번의 호출로 엄격한 제한을 적용합니다.

## Bitbucket Cloud

Bitbucket Cloud에서의 API 요청 속도 제한은 인증된 사용자당 1시간에 1,000회 호출입니다. 자세한 내용은 [가이드](https://support.atlassian.com/bitbucket-cloud/docs/api-request-limits/)에서 설명되어 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 매 시간 최대 1,000회의 호출로 엄격한 제한을 적용하고 429 응답("너무 많은 호출")을 다루기 위한 추가 규제 메커니즘을 적용합니다.

## Bitbucket Server

Bitbucket Server에서는 시스템 관리자가 API 요청 속도 제어에 대한 완전한 제어권을 가집니다. 자세한 내용은 [가이드](https://confluence.atlassian.com/bitbucketserver/improving-instance-stability-with-rate-limiting-976171954.html)에서 설명되어 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 매 시간 최대 1000회의 호출로 중간 제한을 적용하고 429 응답("너무 많은 호출")을 다루기 위한 추가 규제 메커니즘을 적용합니다.

## GitHub

GitHub에서는 인증된 사용자당 1시간에 5,000회의 API 호출이 가능합니다. 자세한 내용은 [가이드](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps)에서 설명되어 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 매 시간 최대 4,500회의 호출로 엄격한 제한을 적용하고 429 응답("너무 많은 호출")을 다루기 위한 추가 규제 메커니즘을 적용합니다.

## GitHub Enterprise

GitHub Enterprise에서는 인증된 사용자당 1시간에 5,000회의 API 호출이 가능합니다. 자세한 내용은 [가이드](https://docs.github.com/en/developers/apps/building-github-apps/rate-limits-for-github-apps)에서 설명되어 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 초당 최대 3회의 호출로 1시간에 총 10,800회의 호출로 엄격한 제한을 적용하고 429 응답("너무 많은 호출")을 다루기 위한 추가 규제 메커니즘을 적용합니다.

## GitLab 및 GitLab Server

GitLab에서는 인증된 사용자당 1분에 300회의 API 호출이 가능합니다. 자세한 내용은 [가이드](https://docs.gitlab.com/ee/user/gitlab\_com/index.html#gitlabcom-specific-rate-limits)에서 확인할 수 있으며, GitLab Server에 대한 내용은 [가이드](https://docs.gitlab.com/ee/user/admin\_area/settings/rate\_limits\_on\_raw\_endpoints.html)에서 확인할 수 있습니다.

`snyk-scm-contributors-count` 도구는 속도 제한을 다루기 위해 매 분 최대 120회의 호출로 엄격한 제한을 적용하고 429 응답("너무 많은 호출")을 다루기 위한 추가 규제 메커니즘을 적용합니다.

{% hint style="info" %}
GitLab Server에서 API 속도 제어는 관리자가 설정 가능하며, 자세한 내용은 [가이드](https://docs.gitlab.com/ee/user/admin\_area/settings/rate\_limits\_on\_raw\_endpoints.html)에서 확인할 수 있습니다.
{% endhint %}