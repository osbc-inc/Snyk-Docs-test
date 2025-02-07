# 분류 정책 - 사용 사례

Policies 뷰에서 **Set Asset Class** 동작을 사용하여 고유 자산을 중요도에 따라 분류할 수 있습니다. 여기서 클래스 A가 가장 중요하고 클래스 D가 가장 중요하지 않은 것으로 취급됩니다.

자산 클래스를 설정할 수 있는 방법은 다음과 같습니다:

* 저장소 이름
* 자산 태그

{% hint style="info" %}
Snyk AppRisk는 GitHub 및 GitLab 주제를 자산 태그로 식별합니다.
{% endhint %}

분류 정책을 사용하여 애플리케이션에 비즈니스 컨텍스트를 제공할 수 있습니다. 분류 정책을 설정하면 모든 자산이 자동으로 분류됩니다.

분류 정책을 사용하기 시작했다면 일단 Class D 자산에 초점을 맞추는 것이 권장됩니다. 이는 그들이 가장 중요하지 않기 때문입니다.

다음 예시는 `sandbox`, `test`, `to-delete`를 이름에 포함하는 자산을 필터링하는 것입니다. Snyk AppRisk에서 GitHub 및 GitLab 주제는 SCM 통합에서 가져와 저장소 자산에 적용되므로 SCM에서 리포지토리에 `PCI-Compliance`와 같은 주제가 추가된 경우, Snyk은 해당 태그를 가져와 해당 자산을 Class A로 분류할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Create policy.png" alt="AppRisk - 분류 정책을 위한 필터 설정"><figcaption><p>Snyk AppRisk - 분류 정책을 위한 필터 설정</p></figcaption></figure>

필터를 설정한 후, 해당 자산에 Class D 자산 분류를 적용해야 합니다.

<figure><img src="../../../../.gitbook/assets/Set action.png" alt="AppRisk - 분류 정책을 위한 동작 설정"><figcaption><p>Snyk AppRisk - 분류 정책을 위한 동작 설정</p></figcaption></figure>

동일한 정책 내에서 Class A, B 및 C 자산에 대해서도 유사한 패턴을 만들고 동작을 생성할 수 있습니다.

<figure><img src="../../../../.gitbook/assets/Set Class.png" alt="AppRisk - 분류 정책을 위한 여러 동작 설정"><figcaption><p>Snyk AppRisk - 분류 정책을 위한 여러 동작 설정</p></figcaption></figure>
