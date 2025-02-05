# 오픈 소스용 Ruby

## 오픈 소스 지원을 위한 Ruby

지원되는 패키지 관리자 및 기능에 대한 자세한 내용은 [Ruby 세부 정보](./)를 참조하십시오.

## 오픈 소스 및 라이선스

{% hint style="info" %}
**기능 가용성**\
귀하의 요금제에 따라 일부 기능을 사용할 수 없는 경우가 있습니다. 자세한 정보는 [요금제 및 가격](https://snyk.io/plans/)을 참조하십시오.

플랫폼별 패키지는 지원되지 않습니다. `Gemfile.lock`에 이러한 패키지가 있는 경우, 이로 인해 잘못된 수정 PR이 생성될 수 있습니다. 가능하다면 패키지의 플랫폼별 변형 대신 비플랫폼별 변형을 사용하십시오.
{% endhint %}

## Bundler 지원

Snyk는 [Bundler](https://bundler.io/)를 사용하여 종속성을 관리하는 Ruby 프로젝트를 CLI 및 Git 통합에서 테스트, 모니터링 및 수정하는 것을 지원하며 특정 종속성 버전을 [Ruby 취약점 데이터베이스](https://snyk.io/vuln?type=rubygems)와 비교합니다.

Snyk는 모든 Bundler 그룹을 테스트합니다. 테스트 또는 개발 그룹과 같은 특정 그룹을 제외하는 것은 불가능합니다.

## Snyk를 통해 지원되는 Ruby의 관리 파일

다음 관리 파일이 지원됩니다:

* `Gemfile`
* `Gemfile.lock`

올바르게 테스트, 모니터링 및 수정하기 위해 Snyk은이 두 파일이 모두 존재해야 합니다.

## **개인 Gem 소스**

Gemfile에서 개인 Gem 소스에 액세스가 필요한 경우, [Ruby 구성을 위한 개인 Gem 소스](../../scan-with-snyk/snyk-open-source/package-repository-integrations/private-gem-sources-for-ruby-configuration.md)를 참조하십시오.

Snyk CLI를 사용하는 경우 개인 Gem 소스를 사용하는 것이 정상적으로 작동해야 합니다.

개인 Gem 소스를 사용하여 Ruby 프로젝트에서 수정 PR을 생성하는 경우, Snyk 파일을 올바르게 업데이트하려면 Gem을 호스팅하는 서비스에 액세스할 수 있어야 할 수 있습니다.

## Ruby 프로젝트에서 취약점 수정

{% hint style="info" %}
Ruby 버전 < 3.2의 경우, 예를 들어 `ruby "2.7.7"`을 Gemfile에 고정하는 것을 Snyk은 지원하지 않습니다. 모든 포인트 버전을 포함하는 더 포용적인 버전 범위를 사용해야 합니다.`ruby "~> 2.7.x"`와 같이
{% endhint %}

Snyk는 Gemfile을 수정한 후 `bundle update`를 사용하여 취약한 젬을 업데이트하여 가능한 한 지정된 규칙을 준수하여 취약점을 수정할 수 있습니다.

일부 시나리오에서 Snyk는 모든 종속성을 취약하지 않은 버전으로 업그레이드할 수 없을 수 있습니다. 이 경우 Gemfile의 규칙을 업데이트하는 것을 고려하십시오.
