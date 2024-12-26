# IaC+ 및 클라우드 컴플라이언스 보고 보기

{% hint style="info" %}
**릴리스 상태**\
+는 폐쇄 베타 버전이며 새로운 고객의 참여를 더 이상 받지 않습니다.\
사용 가능한 기능에 대한 자세한 내용은 [현재 IaC로 시작하기](https://docs.snyk.io/scan-using-snyk/snyk-iac/getting-started-with-current-iac)를 참조하십시오.
{% endhint %}

## 컴플라이언스 보고 개요

는 클라우드 서비스에 대한 관련 컴플라이언스 표준 및 컨트롤을 위한 [컴플라이언스 보고](../../manage-issues/reporting/available-snyk-reports.md#cloud-compliance-issues-report) 및 [클라우드 이슈](getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/) 트리지를 지원합니다. 이 정보를 통해, 개발자들은 문제를 해결하여 클라우드 환경을 컴플라이언스에 맞게 조정하고 감사자는 적절한 증거를 볼 수 있습니다.

* 컴플라이언스 보고를 보는 방법은 [클라우드 컴플라이언스 이슈 보고서](../../manage-issues/reporting/available-snyk-reports.md#cloud-compliance-issues-report)를 참조하십시오.
* 컴플라이언스 표준 및 컨트롤에 따라 클라우드 이슈를 필터링하여 트리지하는 방법은 [이슈 필터링](getting-started-with-iac+-and-cloud-scans/manage-iac+-and-cloud-issues/view-iac+-and-cloud-issues-in-the-snyk-web-ui.md#filter-issues)을 참조하십시오.

컴플라이언스 표준 및 컨트롤, 보안 규칙 등에 대한 정의는 [중요 개념](getting-started-with-iac+-and-cloud-scans/key-concepts-for-iac+-and-cloud.md)을 참조하십시오.

## 지원되는 컴플라이언스 표준

| 컴플라이언스 표준                 | 상태              |
| ----------------------------------- | ------------------- |
| AWS Well Architected (2020-07-02)   | 일반적으로 사용 가능 |
| CSA Cloud Controls Matrix (v3.0.1)  | 일반적으로 사용 가능 |
| CSA Cloud Controls Matrix (v4.0.5)  | 베타                |
| CIS Kubernetes Benchmark (v1.6.1)   | 일반적으로 사용 가능 |
| CIS AWS Benchmark (v1.2.0)          | 일반적으로 사용 가능 |
| CIS AWS Benchmark (v1.3.0)          | 일반적으로 사용 가능 |
| CIS AWS Benchmark (v1.4.0)          | 일반적으로 사용 가능 |
| CIS AWS Benchmark (v1.5.0)          | 일반적으로 사용 가능 |
| CIS AWS Benchmark (v2.0.0)          | 일반적으로 사용 가능 |
| CIS Google Cloud Benchmark (v1.1.0) | 일반적으로 사용 가능 |
| CIS Google Cloud Benchmark (v1.2.0) | 일반적으로 사용 가능 |
| CIS Google Cloud Benchmark (v1.3.0) | 일반적으로 사용 가능 |
| CIS Azure Benchmark (v1.1.0)        | 일반적으로 사용 가능 |
| CIS Azure Benchmark (v1.3.0)        | 일반적으로 사용 가능 |
| CIS Azure Benchmark (v1.4.0)        | 일반적으로 사용 가능 |
| CIS Controls (v7.1)                 | 일반적으로 사용 가능 |
| CIS Controls (v8.0)                 | 베타                |
| GDPR (2016)                         | 일반적으로 사용 가능 |
| HIPAA (2013)                        | 일반적으로 사용 가능 |
| ISO/IEC 27001 (2013)                | 일반적으로 사용 가능 |
| NIST SP 800-53 (Rev4)               | 일반적으로 사용 가능 |
| PCI DSS (v3.2.1)                    | 일반적으로 사용 가능 |
| SOC 2 (2017)                        | 일반적으로 사용 가능 |