# Java 지원에 대한 자세한 정보

## Maven Bill of Materials (BOM)

Maven은 함께 작동하는 의존성 버전을 알려주는 [Bill of Materials (BOM) POM 파일](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html#bill-of-materials-bom-poms)을 지원합니다.

## BOM 파일의 내용 및 사용

BOM 파일에는 다음이 포함됩니다:

- `pom` 패키징 유형: `<packaging>pom</packaging>`.
- `dependencyManagement` 섹션.

외부 프로젝트는 의존성 관리를 쉽게 만들기 위해 BOM 파일을 제공할 수 있습니다. 여기 일반적인 예시가 있습니다:

- [spring-data-bom](https://github.com/spring-projects/spring-data-bom) - Spring 팀이 Spring Data 프로젝트를 위한 BOM을 제공합니다.
- [jackson-bom](https://github.com/FasterXML/jackson-bom) - Jackson 프로젝트가 Jackson 의존성을 위한 BOM을 제공합니다.

다음은 BOM 파일의 예시입니다:

{% code title="예시 1 - BOM 파일" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <groupId>snyk</groupId>
    <artifactId>snyk-bom</artifactId>
    <version>1.0</version>
    <packaging>pom</packaging>
    <name>Snyk Bill Of Materials</name>
    
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>log4j</groupId>
                <artifactId>log4j</artifactId>
                <version>1.2.12</version>
            </dependency>
            <dependency>
                <groupId>commons-logging</groupId>
                <artifactId>commons-logging</artifactId>
                <version>1.1.1</version>
            </dependency>
        </dependencies>
    </dependencyManagement>
</project>
```
{% endcode %}

`dependencyManagement` 섹션에는 의존성 요소가 포함됩니다. 각 의존성은 Maven이 전환(및 직접) 의존성을 선택하기 위해 참조하는 것입니다.

`dependencyManagement` 섹션에서 의존성을 정의하면 해당 의존성이 프로젝트의 의존성 트리에 추가되지 않습니다. 이것은 오직 참조 조회용으로 사용됩니다.

이 BOM은 부모로 프로젝트 POM에 가져올 수 있습니다. `log4j` 버전을 명시할 필요가 없습니다. BOM에서 상속받습니다:

{% code title="예시 2 - 프로젝트 POM" %}
```xml
<project ...>
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>snyk</groupId>
        <artifactId>snyk-bom</artifactId>
        <version>1.0</version>
    </parent>
    
    <groupId>snyk</groupId>
    <artifactId>snyk-project</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <packaging>jar</packaging>
    <name>Snyk Project</name>
    
    <dependency>
        <groupId>log4j</groupId>
        <artifactId>log4j</artifactId>
    </dependency>
</project>
```
{% endcode %}

## Snyk가 BOM 처리하는 방법

### 의존성 해결

Snyk는 BOM의 `dependencyManagement` 조회에 있는 버전을 부모로 가져온 프로젝트 POM에서 선언된 어떤 의존성에 적용합니다.

Snyk가 BOM 파일을 스캔할 때, `dependencyManagement` 내용은 해당 파일의 의존성으로 간주되지 않습니다. 이것들은 오직 조회용입니다.

다음은 이전 예시 파일 각각을 어떻게 분석하고 처리하는지 설명합니다.

- BOM 파일 - Snyk는 이 파일에 대해 어떤 의존성도 포함하지 않기 때문에 Snyk 프로젝트를 만들지 않습니다.
- 프로젝트 POM - Snyk는 `v1.2.12.`를 가진 `log4j`의 단일 의존성이 있는 프로젝트를 생성합니다. Snyk는 부모 BOM의 규칙을 적용하여 `log4j`의 올바른 버전을 식별합니다. `commons-logging` 의존성은 직접 프로젝트 POM에서 선언되지 않았기 때문에 포함되지 않습니다.

{% hint style="info" %}
BOM이 `dependencyManagement` 외부에 직접 의존성을 가지고 있는 경우, Snyk는 해당 BOM에 대한 프로젝트를 만듭니다.
{% endhint %}

### 수정 PRs

Snyk는 Fix PR 기능을 통해 취약한 패키지를 업그레이드하는 권장 사항을 제공합니다.

Fix PR은 문제가 보고된 POM 파일 내에서 관리되는 버전의 의존성에 대해서만 작성할 수 있습니다.

만약 버전이나 의존성이 부모 BOM에서 관리되고 있다면, Snyk는 취약한 경로를 수정할 수 있는 것을 확인할 수 있지만 해당 수정을 적용할 수 없습니다.

## 패키지 레지스트리 통합 (Artifactory/Nexus) - Maven

Artifactory 및 Nexus 패키지 레지스트리 통합은 Snyk 엔터프라이즈 요금제 사용자에게 제공됩니다.

- {{Snyk 오픈 소스}}는 Artifactory 또는 Nexus를 통해 개인 패키지를 통해 전환 의존성을 해결합니다.
- Snyk는 공개적으로 사용 가능한 인스턴스에 사용자 이름과 비밀번호를 사용하거나 Snyk 브로커를 사용하여 네트워크의 사설 서버에 연결할 수 있습니다.
- {{Snyk 오픈 소스}}는 Artifactory 및 Nexus 모두와 보안 테스트를 위해 레지스트리와 대화하거나 지역 확인자로써 통합을 제공합니다. [Nexus 저장소 관리자 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/nexus-repository-manager-connection-setup/) 및 [Artifactory 레지스트리 설정](../../scan-with-snyk/snyk-open-source/package-repository-integrations/artifactory-package-repository-connection-setup/)을 참조하세요.

{% hint style="info" %}
Snyk 엔터프라이즈 사용자가 아니고 Artifactory 또는 Nexus를 사용하는 경우, 분석은 빌드 시스템이 종속성을 검색하고 로컬로 보유하기 때문에 CLI를 사용하여 최상으로 실행됩니다.
{% endhint %}

## 문제 해결

도움이 필요하면 [Snyk 지원에 문의하십시오](https://support.snyk.io).