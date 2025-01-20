# 그룹과 조직 간 전환

Snyk은 Snyk 웹 UI에 로그인할 때 기본적으로 귀하의 기본 조직을 표시합니다. 또한 Snyk은 CLI를 사용하여 프로젝트를 테스트할 때 기본 조직 설정을 사용합니다. 자세한 정보는 [조직 관리](organizations/create-and-delete-organizations.md)를 참조하십시오.

## 그룹 변경

회사에 여러 그룹이 있는 경우 보고 계신 그룹을 명심해야 합니다. 각 그룹에는 다른 조직이 포함되어 있고 다른 설정이 있습니다.

다른 그룹으로 이동하려면 그룹 전환기 (![Switcher](<../../.gitbook/assets/image (4) (3) (2).png>) 아이콘)를 클릭하고 그룹을 선택하십시오:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-04-25 at 10.06.46.png" alt="그룹 변경"><figcaption><p>그룹 변경</p></figcaption></figure></div>

개인 조직을 보거나 추가하려면 **미그룹**을 선택하십시오.

## 웹 UI에서 조직 변경

{% hint style="warning" %}
SSO 이메일이 Google로 제공된 경우 해당 이메일을 통해 Google 소셜 제공자를 통해 로그인하면 SSO를 사용하여 로그인하는 것과 같지 않습니다. Google 소셜 옵션을 사용하면 기본 무료 사용자 권한이 있는 별도의 무료 계정이 생성됩니다. 서로 다른 계정에 속한 조직 간에 전환할 수 없습니다.
{% endhint %}

또한 보고 계신 조직을 알아야 합니다. 조직에는 다른 프로젝트가 포함되어 있습니다.

GitHub 통합을 통해 프로젝트를 추가한 경우 해당 프로젝트는 보고 계신 조직, 즉 선택한 조직에 추가됩니다.

다른 조직으로 이동하려면 조직 전환기를 클릭하고 조직을 선택하십시오:

<div align="left"><figure><img src="../../.gitbook/assets/Screenshot 2023-03-13 at 10.31.14.png" alt="조직 변경"><figcaption><p>조직 변경</p></figcaption></figure></div>

## CLI에서 조직 변경

1. 기본 조직만 있는 경우 `snyk test` 또는 `snyk monitor`를 실행하여 추가하거나 업데이트하는 모든 프로젝트는 자동으로 기본 조직과 연결됩니다.
2. 여러 조직이 있는 경우 `snyk config set org=ORG_ID`를 실행하여 새 프로젝트가 연결될 조직을 설정할 수 있습니다.
3. 개별 `snyk monitor` 실행의 전역 구성을 재정의하려면 `snyk test --org=ORG_ID` 또는 `snyk monitor --org=ORG_ID`를 실행하십시오.

기본 `<ORG_ID>`는 [계정 설정](https://app.snyk.io/account)에서 중요 조직으로 현재 선호되는 조직입니다.

자세한 정보는 [CLI에서 사용할 조직 선택하기](../../snyk-cli/scan-and-maintain-projects-using-the-cli/how-to-select-the-organization-to-use-in-the-cli.md)를 참조하십시오.
