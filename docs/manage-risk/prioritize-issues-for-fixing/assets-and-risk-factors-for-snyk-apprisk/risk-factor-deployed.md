# 리스크 요소: 배포됨

{% hint style="info" %}
배포된 리스크 요소는 Snyk AppRisk Pro 사용자에게만 제공됩니다.
{% endhint %}

어떤 배포된 코드든 여러분의 애플리케이션 및 비즈니스의 악용 가능성을 증가시킵니다.

배포된 코드의 이해와 배포된 위치가 여러분에게 실행 중인 코드의 리스크 표면적을 줄이기 위한 조치 전략을 채택할 수 있도록 도와줍니다.

## 통합 유형

배포된 리스크 요소는 Kubernetes 커넥터, Snyk 런타임 센서, 그리고 제3자 통합과 함께 작동합니다.

### Kubernetes **커넥터**

Snyk AppRisk는 Kubernetes 클러스터에서 실행 중인 이미지와 생성된  프로젝트 사이의 일치를 찾아 컨테이너 이미지가 배포되었음을 결정합니다.

Snyk AppRisk는 Kubernetes 상태 정보를 사용하여 현재 실행 중인 Docker 이미지 식별자를 추출합니다. Kubernetes 컨테이너의 상태에는 Kubernetes 런타임에서 실행 중인 이미지 이름이 있습니다. 알려진 Docker 이미지 데이터베이스에서 일치하는 이름을 찾습니다. 이미지 이름이 일치하는 경우, Snyk은 그래프에서 이 정보를 표시할 수 있습니다. 그래프는 이미지와 컨테이너 간의 관계를 보여줍니다.

<figure><img src="https://lh6.googleusercontent.com/BoYMeFGbzjUmNmXbmtrklBcl9LLm9S94mwJWkrFA_5E5WIO07BsS3Zv-fbGBlXkNAx4oGnbBtzFijWTxUQbsnlzJI2QqprUJWPevpwBybhmwtzQayYnmW6_Qvhddgz1_vdy-NDZgQKUQhmxnY54xkrI" alt="배포된 이미지의 취약성"><figcaption><p>배포된 이미지의 취약성</p></figcaption></figure>

Kubernetes는 이미지 관리에 대해 [구체적](https://kubernetes.io/docs/concepts/containers/images/#image-names)으로 명시되어 있습니다. Snyk는 Snyk가 알고 있는 이미지를 매핑하는 데 이와 같은 논리를 사용합니다. 로 이미지를 스캔할 때마다, Snyk는 이미지 이름 및 이미지 ID에 대한 정보를 수집합니다. 이 정보를 사용하여 이미지를 Kubernetes 정보와 매핑합니다.

{% hint style="info" %}
Snyk는 [Kubernetes](https://kubernetes.io/docs/concepts/containers/images/#image-names)와 Docker [이미지](https://docs.docker.com/engine/reference/commandline/images/)에 문서화된 지정된 명명 표준을 준수하여 Kubernetes와 일관성을 유지합니다.&#x20;
{% endhint %}

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.01.48.png" alt="명명 표준"><figcaption><p>명명 표준</p></figcaption></figure>

다음 예제를 고려하십시오.

<table><thead><tr><th width="267.3333333333333">Kubernetes 매니페스트에서 제공된 이름 </th><th>일치에 사용된 이름</th><th>변경 (예/아니오)</th></tr></thead><tbody><tr><td>gcr.io/my-company/my-app:production</td><td>gcr.io/my-company/my-app:production</td><td>아니요</td></tr><tr><td>gcr.io/my-company/my-app:latest</td><td>gcr.io/my-company/my-app:latest</td><td>아니요</td></tr><tr><td>gcr.io/my-company/my-app</td><td>gcr.io/my-company/my-app:latest</td><td>예 - 최신 태그가 추가됨</td></tr><tr><td>my-app</td><td>docker.io/my-app/my-app:latest</td><td>예 - Docker 공개 레지스트리로 기본 설정, 최신 태그가 추가됨</td></tr></tbody></table>

일치는 다음의 우선 순위를 사용하며, 첫 번째 단계는 적어도 하나의  프로젝트에서 통과해야 하며, 그 다음 단계에서 일치를 더 검증합니다.&#x20;

1. 이미지 이름 일치, 예를 들어 `gcr.io/my-company/my-app:latest`.
2. 이미지 다이제스트 일치.
3. 이미지 다이제스트로  프로젝트 그룹화.

다음 예제를 고려하십시오.

#### **예제 1:  CLI 사용**

결과: 이미지가 성공적으로 일치되고 리스크 요소가 적용됨

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.04.31.png" alt="이미지 일치"><figcaption><p>이미지 일치</p></figcaption></figure>

컨테이너 이미지를  CLI를 사용하여 스캔하며, 레지스트리를 포함한 이미지의 전체 이름을 참조합니다. 이미지가 빌드된 후 클러스터에 배포되기 전에 이를 수행하는 것을 추천합니다.

다음은 스캔의 예시입니다:

&#x20;`$ snyk container monitor gcr.io/my-company/my-app:latest`

이 이미지는 다음과 같이 클러스터에 배포됩니다.

`spec:`

&#x20; `containers:`

&#x20; `- name: my-app`

&#x20;   `image: gcr.io/my-company/my-app:latest`

이렇게 함으로써 Insights는 이미지 이름을 성공적으로 일치시키고 이  프로젝트에 연관된 모든 문제에 배포된 리스크 요소를 적용할 수 있습니다.

#### **예제 2:  CLI 및 컨테이너 레지스트리 사용**

결과: 이미지가 성공적으로 일치되고 리스크 요소가 적용됨

<figure><img src="../../../.gitbook/assets/Screenshot 2023-07-12 at 02.05.31.png" alt="이름 일치"><figcaption><p>이름 일치</p></figcaption></figure>

부분 이름을 참조하여 컨테이너 이미지를 스캔하며, 컨테이너 레지스트리를 생략합니다.

다음은 스캔의 예시입니다:

`$ snyk container monitor my-app:lates`

Insights는 이 프로젝트를 일치시키지 못하며 이름이 일치하지 않습니다.

이 이미지는 다음과 같이 클러스터에 배포됩니다.

`spec:`

&#x20; `containers:`

&#x20; `- name: my-app`

&#x20;   `image: gcr.io/my-company/my-app:latest`

또한, 동일한 이미지가 컨테이너 레지스트리에서도 스캔됩니다.

이로써, 컨테이너 레지스트리를 통해 전체 이름을 사용한 프로젝트가 생성되어 일치가 가능해집니다.

이것은 이미지 다이제스트도 포함합니다.

그러면 Insights는 이미지 다이제스트로 모든  프로젝트를 그룹화하게 되어 컨테이너 레지스트리를 통해 부분 이름을 가진 CLI 프로젝트에 대해 배포된 리스크 요소를 적용할 수 있습니다.

{% hint style="info" %}
Snyk는 CLI 명령어에서 이미지의 전체 이름을 지정할 것을 권장합니다. 이를 수행할 수 없는 경우, Snyk는 두 번째 통합을 통해 동일한 이미지를 스캔할 것을 권장합니다.&#x20;
{% endhint %}

### Snyk 런타임 센서

Snyk 런타임 센서는 Kubernetes 클러스터 내에서 활성 컨테이너를 지속적으로 모니터링하여 배포된 리스크 요소를 적용합니다. 이를 위해 실시간 Kubernetes 이벤트를 캡처하고 클러스터 상태의 시간당 스냅숏을 찍습니다. 이러한 실시간 데이터 및 스냅숏을 사용하여 현재 배포된 이미지를 식별합니다. 이 정보를 Snyk 프로젝트 및 알려진 취약점과 비교하면, Snyk 런타임 센서는 정확하게 배포된 리스크를 평가합니다. 이를 통해 보안 인사이트가 최신이며 현재 배포 환경을 반영하도록 합니다.

### 제3자 통합

배포된 리스크 요소는 외부 소스의 데이터와 함께 Kubernetes 클러스터를 실시간 모니터링해야 하는 제3자 통합에까지 확장됩니다. 이를 통해 다양한 환경에서 배포된 리스크를 종합적으로 평가할 수 있습니다. 제3자 도구와 통합하여 Snyk는 다양한 소스의 취약점을 상호 참조하여 배포된 리스크 요소가 가장 정확하고 현재의 위협 환경을 반영하도록 합니다. 이 종합적인 방식을 통해 Snyk는 각 통합된 플랫폼에서 실행 가능한 인사이트를 제공하고 견고한 보안 포지션을 유지할 수 있도록 합니다.

## Snyk Insights 배포된 리스크 요소에 대한 기술적 상세

Kubernetes 커넥터는 Kubernetes 이벤트를 계속 모니터링합니다. 이러한 이벤트는 지속적으로 Snyk 플랫폼에 스트리밍됩니다.&#x20;

정기적으로 매 시간, 데이터 파이프라인은 클러스터 상태의 조정을 수행하여 스냅숏을 생성합니다. 이 스냅숏은 클러스터에서 실행 중인 이미지를 결정하는 데 사용됩니다.

동일한 간격으로, 데이터 파이프라인은 모든 Snyk 프로젝트 및 데이터 원본의 스냅숏을 찍고 패키지와 이미지를 추출합니다. 이 스냅숏은 Snyk가 특정 고객을 위해 알려진 이미지와 패키지를 결정하는 데 사용됩니다.&#x20;

두 스냅숏은 비교되어 해당 시점에서 배포된 사실을 확인하는 증거 그래프가 생성됩니다.&#x20;