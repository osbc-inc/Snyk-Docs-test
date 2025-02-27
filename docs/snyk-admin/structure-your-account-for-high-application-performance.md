# 높은 애플리케이션 성능을 위한 계정 구조화

Snyk을 사용하는 최상의 경험을 보장하기 위해, 사용자 계정, 그룹, 조직 및 프로젝트에 대한 결정을 내릴 때 다음 가이드라인을 고려하십시오.

## 사용자 계정의 구조

조직과 그룹에 많은 수의 사용자가 있을 수 있습니다.

Snyk은 각 조직에 2,000명을 초과하는 사용자가 없도록 당신의 조직을 구조화하는 것을 권장합니다.

조직에 **2,000명 이상의 사용자가 있다면**, 성능 문제가 발생할 수 있습니다. 대규모 사용자를 로드해야 하는 경우 대시보드 및 그룹 구성원 관리 페이지의 성능이 저하됩니다.

사용자가 **특정 그룹에 많은 수의 멤버십을 가지고 있다면**, Snyk 웹 UI 및 CLI 및 API를 통한 모든 요청이 성능이 떨어지게 됩니다. 대부분의 요청에서 액세스 및 권한을 확인하기 위해 계산 및 쿼리가 발생하기 때문입니다.

## 그룹의 구조

**Snyk 사용자 중 소수의 고객은 여러 그룹을 가지고 있습니다**. 예를 들어, 서로 완전히 분리된 비즈니스 부문을 유지하기 위해 여러 그룹을 설정하는 경우가 있습니다. 그러나 복수의 그룹을 고려하는 사람들은 본인의 계정을 여러 그룹으로 구성하는 제한 사항을 이해해야 합니다.

구체적으로, 각 그룹은 독립된 엔티티입니다. 이에 따라 다음과 같은 결과가 나타납니다:

* 그룹의 기능이 서로 연결되어 있지 않습니다.
* 교차 그룹 보고서가 없습니다.
* 사용자, 프로젝트 및 조직은 그룹 간에 공유할 수 없습니다.
* 여러 그룹 간의 SSO 관리가 더 어렵습니다.
* 서비스 계정은 여러 그룹에 걸쳐 사용할 수 없습니다.
* API를 통해 여러 그룹의 데이터를 가져오려면 여러 번의 호출이 필요합니다.

여러 그룹이 필요한 경우, Snyk 계정팀과 상의하시길 바랍니다.

## 조직의 구조

Snyk 웹 UI 또는 Snyk API를 사용하여 그룹에 많은 수의 조직을 생성할 수 있습니다. 그러나 그룹에 2,000개 이상의 조직이 있다면 성능 문제가 발생할 수 있습니다.

**Snyk이 대규모의 조직을 로드하는 경우, 다음과 같은 결과가 발생합니다**:

* 그룹 관리자 및 그룹 수준 알림에 대한 성능이 저하됩니다.
* 그룹 수준 서비스 계정 생성이 실패할 수 있습니다.

## 프로젝트의 구조

조직에 많은 수의 프로젝트를 가져올 수 있습니다.

Snyk은 각 조직을 **10,000개 이상의 프로젝트로 제한**할 것을 권장하며, 각 조직에 25,000개 이상의 프로젝트를 허용하지 않습니다.

10,000개 이상의 프로젝트가 필요한 경우, **대량의 프로젝트가 프로젝트 목록, 알림 보내기, 대시보드 표시 및 사용 페이지 표시와 같은 경험에 미치는 영향**을 고려해야 합니다. 또한, 대규모의 프로젝트를 가진 조직을 삭제하는 것이 어렵습니다.

그룹 내 모든 조직을 통한 전체 프로젝트 수에 제한은 없지만, 수십만 프로젝트가 있는 경우 그룹 수준 보고 및 종속성 및 라이선스 문제 표시가 느려집니다.
