# 컨테이너 이미지 스캔

Snyk Container는 컨테이너 레지스트리 스캔을 기반으로 컨테이너 이미지에서 취약점을 찾아 수정하는 데 도움을 줍니다.

Snyk Container를 사용하여 컨테이너 이미지를 스캔할 수 있습니다:

* [Snyk 웹 UI](use-snyk-container/)에서
* [Snyk CLI](../../snyk-cli/scan-and-maintain-projects-using-the-cli/snyk-cli-for-snyk-container/)를 통해
* [브로커](../../enterprise-setup/snyk-broker/snyk-broker-container-registry-agent/integrate-with-self-hosted-container-registries-broker.md)로 (자체 호스트된 컨테이너 레지스트리용)

## **Snyk Container를 웹 UI에서 사용하기 위한 전제조건**

Snyk Container를 사용하여 컨테이너 이미지를 스캔하기 전에 다음을 확인하세요:

* Snyk 계정을 생성하거나 로그인합니다.
* Docker Hub와 같은 지원되는 컨테이너 레지스트리와 통합을 설정합니다. [컨테이너 보안 통합](container-registry-integrations/)을 참조하세요.

자세한 내용은 [시작하기](../../getting-started/)를 참조하세요.

## 컨테이너 이미지에서 취약점 보기

**프로젝트** 탭에서 가져온 Snyk 프로젝트에 대한 취약점 결과를 볼 수 있습니다. 가져온 프로젝트는 **대상(Targets)**&#xC73C;로 그룹화됩니다.

{% hint style="info" %}
조직에 가져온 모든 저장소 및 컨테이너 레지스트리 이미지의 이력을 볼 수 있습니다. 자세한 내용은 [가져오기 로그](../../snyk-admin/snyk-projects/import-log.md)를 참조하세요.
{% endhint %}

해당 프로젝트에 대한 취약성 정보를 보려면 대상 목록에서 가져온 프로젝트를 선택하세요.

프로젝트 항목을 클릭하여 발견된 취약점의 자세한 정보와 도입된 위치, 수정 방법, 취약점에 대한 기타 정보를 볼 수 있습니다.

## 컨테이너 이미지에서 취약점 수정하기

컨테이너 이미지의 취약성을 수정하려면:

1. **수정 PR 열기**를 클릭하여 Snyk 권장에 따라 PR을 엽니다.
2. 이미지를 업그레이드하거나 다시 빌드하세요.

업데이트된 이미지를 푸시한 후, Snyk이 자동으로 새 이미지를 다시 스캔합니다.

취약성 수정에 대한 자세한 내용은 [Snyk 웹 UI에서 이미지 분석 및 수정](use-snyk-container/analyze-and-fix-container-images.md)을 참조하세요.
