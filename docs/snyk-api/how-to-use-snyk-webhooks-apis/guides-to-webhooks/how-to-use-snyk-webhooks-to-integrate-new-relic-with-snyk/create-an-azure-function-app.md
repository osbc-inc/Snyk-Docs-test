# Azure Function App 생성

Azure Function App을 생성하고 구성하는 일반적인 프로세스를 따르세요. 자세한 정보는 [Microsoft Azure Functions 문서](https://learn.microsoft.com/en-us/azure/azure-functions/)를 참조하십시오.

다음은 Azure Function 코드의 예시입니다.

```csharp
#r "Newtonsoft.Json"

using System.Net;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.String;
using Newtonsoft.Json;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using Newtonsoft.Json.Linq;
using System.Web;
using System.Text.RegularExpressions;

public static async Task<IActionResult> Run(HttpRequest req, ILogger log)
{
    log.LogInformation("C# HTTP trigger function processed a request.");
    string responseMessage = "유효한 페이로드를 받지 못했습니다!";

    try
    {
        string name = req.Query["name"];

        string requestBody = await new StreamReader(req.Body).ReadToEndAsync();
        dynamic data = JsonConvert.DeserializeObject(requestBody);
        //log.LogInformation("data: " + requestBody);

        if (data.project != null)
        {
            string count = data.newIssues.Count.ToString();
            string projectName = data.project.name;
            string[] projectNameParts = projectName.Split(':');
            string containerImage = projectName;
            if (projectNameParts.Length > 1)
            {
                containerImage = projectNameParts[1] + ":" + data.project.imageTag;
            }
            string repoURL = data.project.name;
            string artifactURL = data.project.name;
            string entityLookupValue = repoURL;
            string entityType = "Repository";
            int idxRepoURLBranch = repoURL.IndexOf("(");
            if (data.project.origin == "github")
            {
                int idxRepoURLProject = repoURL.IndexOf(":");
                string package = "";
                if (idxRepoURLBranch >= 0)
                {
                    package = repoURL.Substring(idxRepoURLProject + 1, repoURL.Length - idxRepoURLProject - 1);
                    log.LogInformation("package: " + package);
                    entityLookupValue = "https://github.com/" + repoURL.Substring(0, idxRepoURLBranch);
                    repoURL = "https://github.com/" + repoURL.Substring(0, idxRepoURLBranch);

                    if (data.project.branch != "")
                    {
                        repoURL = repoURL + "/blob/" + data.project.branch + "/" + package;
                    }
                }
                else
                {
                    entityLookupValue = "https://github.com/" + repoURL.Substring(0, idxRepoURLProject);
                    repoURL = "https://github.com/" + repoURL.Substring(0, idxRepoURLProject);
                }
            }
            else if (data.project.origin == "docker-hub")
            {
                entityType = "ContainerImage";
                entityLookupValue = data.project.imageId;
                if (idxRepoURLBranch >= 0)
                {
                    artifactURL = "https://hub.docker.com/repository/docker/" + repoURL.Substring(0, idxRepoURLBranch);
                    repoURL = "https://hub.docker.com/repository/docker/" + repoURL.Substring(0, idxRepoURLBranch);

                    if (data.project.branch != "")
                    {
                        repoURL = repoURL + "/tree/" + data.project.branch;
                    }
                }
                else
                {
                    int idxRepoURLProject = repoURL.IndexOf(":");
                    artifactURL = "https://hub.docker.com/repository/docker/" + repoURL.Substring(0, idxRepoURLProject);
                    repoURL = "https://hub.docker.com/repository/docker/" + repoURL.Substring(0, idxRepoURLProject);
                }
            }
            else if (data.project.origin == "azure-repos")
            {
                var AZURE_DEVOPS_ORG = Environment.GetEnvironmentVariable("AZURE_DEVOPS_ORG");
                int idxRepoURLProject = repoURL.IndexOf("/");
                string package = "";
                if (idxRepoURLBranch >= 0)
                {
                    package = repoURL.Substring(idxRepoURLProject + 1, repoURL.Length - idxRepoURLProject - 1);
                    log.LogInformation("package: " + package);
                    entityLookupValue = "https://dev.azure.com/" + AZURE_DEVOPS_ORG + "/" + repoURL.Substring(0, idxRepoURLProject);
                    repoURL = "https://dev.azure.com/" + AZURE_DEVOPS_ORG + "/" + repoURL.Substring(0, idxRepoURLBranch);

                    if (data.project.branch != "")
                    {
                        repoURL = repoURL + "/blob/" + data.project.branch + "/" + package;
                    }
                }
                else
                {
                    entityLookupValue = "https://dev.azure.com/" + AZURE_DEVOPS_ORG + "/" + repoURL.Substring(0, idxRepoURLProject);
                    repoURL = "https://dev.azure.com/" + AZURE_DEVOPS_ORG + "/" + repoURL.Substring(0, idxRepoURLProject);
                }
            }

            log.LogInformation(projectName + ", data.newIssues.Count: " + count);
            responseMessage = "새로운 문제가 발견되지 않았습니다. 처리할 내용이 없습니다!";

            name = name ?? data?.name;
            string browseUrl = data.project.browseUrl;
            int x = 0;

            StringBuilder sb = new StringBuilder();
            sb.Append("{\"findings\":[");

            if (data.newIssues.Count > 0)
            {
                log.LogInformation("새로운 문제가 발견되었습니다!");

                for (int i = 0; i < data.newIssues.Count; i++)
                {
                    string id = data.newIssues[i].id.ToString();
                    string issueType = data.newIssues[i].issueType;
                    string pkgName = data.newIssues[i].pkgName;
                    int priorityScore = data.newIssues[i].priorityScore;
                    string title = data.newIssues[i].issueData.title;
                    string issueId = data.newIssues[i].issueData.id;
                    string issueVendorId = issueId;
                    if (data.newIssues[i].issueData.identifiers.CVE.Count > 0)
                    {
                        issueId = data.newIssues[i].issueData.identifiers.CVE[0];
                    }
                    else if (data.newIssues[i].issueData.identifiers.CWE.Count > 0)
                    {
                        issueId = data.newIssues[i].issueData.identifiers.CWE[0];
                    }
                    double cvssScore = data.newIssues[i].issueData.cvssScore;
                    string severity = data.newIssues[i].issueData.severity.ToString().ToUpper();
                    string issueSeverity = data.newIssues[i].issueData.severity;
                    string descr = data.newIssues[i].issueData.description.ToString();
                    if (data.newIssues[i].issueData.description.ToString().Length >= 256)
                    {
                        descr = data.newIssues[i].issueData.description.ToString().Substring(0, 256);
                    }
                    descr = descr.Replace("\n", "").Replace("\r", "");
                    bool remediationExists = data.newIssues[i].fixInfo.isFixable;
                    string remediationRecommendation = "";
                    if (remediationExists)
                    {
                        remediationRecommendation = "upgrade " + pkgName + " to " + data.newIssues[i].fixInfo.fixedIn[0];
                    }

                    if (i > 0)
                    {
                        sb.Append(",");
                    }

                    sb.Append("{");
                    sb.Append("  \"artifactURL\": \"" + artifactURL + "\",");
                    sb.Append("  \"containerImage\": \"" + containerImage + "\",");
                    sb.Append("  \"cvss.score\": \"" + cvssScore + "\",");
                    sb.Append("  \"cvssScore\": \"" + cvssScore + "\",");
                    sb.Append("  \"disclosureUrl\": \"" + browseUrl + "\",");
                    sb.Append("  \"entityLookupValue\": \"" + entityLookupValue + "\",");
                    sb.Append("  \"entityType\": \"" + entityType + "\",");
                    sb.Append("  \"imageBaseImage\": \"" + data.project.imageBaseImage + "\",");
                    sb.Append("  \"imageId\": \"" + data.project.imageId + "\",");
                    sb.Append("  \"imagePlatform\": \"" + data.project.imagePlatform + "\",");
                    sb.Append("  \"imageTag\": \"" + data.project.imageTag + "\",");
                    sb.Append("  \"issueCountsBySeverityCritical\": " + data.project.issueCountsBySeverity.critical + ",");
                    sb.Append("  \"issueCountsBySeverityHigh\": " + data.project.issueCountsBySeverity.high + ",");
                    sb.Append("  \"issueCountsBySeverityMedium\": " + data.project.issueCountsBySeverity.medium + ",");
                    sb.Append("  \"issueCountsBySeverityLow\": " + data.project.issueCountsBySeverity.low + ",");
                    sb.Append("  \"issueId\": \"" + issueId + "\",");
                    sb.Append("  \"issueInstanceKey\": \"" + repoURL + "\",");
                    sb.Append("  \"issueSeverity\": \"" + issueSeverity + "\",");
                    sb.Append("  \"issueType\": \"Library Vulnerability\",");
                    sb.Append("  \"issueVendorId\": \"" + issueVendorId + "\",");
                    sb.Append("  \"message\": \"" + descr + "\",");
                    sb.Append("  \"pkgName\": \"" + pkgName + "\",");
                    sb.Append("  \"projectName\": \"" + projectName + "\",");
                    sb.Append("  \"priorityScore\": " + priorityScore + ",");
                    sb.Append("  \"remediation.exists\": \"" + remediationExists + "\",");
                    sb.Append("  \"remediationExists\": \"" + remediationExists + "\",");
                    sb.Append("  \"remediationRecommendation\": \"" + remediationRecommendation + "\",");
                    sb.Append("  \"severity\": \"" + severity + "\",");
                    sb.Append("  \"snykIssueType\": \"" + issueType + "\",");
                    sb.Append("  \"snykOrigin\": \"" + data.project.origin + "\",");
                    sb.Append("  \"source\": \"Snyk\",");
                    sb.Append("  \"title\": \"" + title + "\"");
                    sb.Append("}");
                }
            }

            sb.Append("]}");

            string payload = sb.ToString();
            payload = payload.Replace(System.Environment.NewLine, ". ");

            if (payload != "{\"findings\":[]}")
            {
                log.LogInformation("payload: " + payload);
                var content = new StringContent(payload, Encoding.UTF8, "application/json");
                var NEW_RELIC_SECURITY_URL = Environment.GetEnvironmentVariable("NEW_RELIC_SECURITY_URL");
                var NEW_RELIC_LICENSE_KEY = Environment.GetEnvironmentVariable("NEW_RELIC_LICENSE_KEY");

                var url = NEW_RELIC_SECURITY_URL;
                var headerKey = NEW_RELIC_LICENSE_KEY;

                using var client = new HttpClient();
                client.DefaultRequestHeaders.Clear();
                client.DefaultRequestHeaders.Accept.Add(new System.Net.Http.Headers.MediaTypeWithQualityHeaderValue("application/json"));
                client.DefaultRequestHeaders.TryAddWithoutValidation("Content-Type", "application/json");
                client.DefaultRequestHeaders.Add("Api-Key", headerKey);
                var response = await client.PostAsync(url, content);

                string result = response.Content.ReadAsStringAsync().Result;
                log.LogInformation("response.StatusCode: " + response.StatusCode);
                if (response.StatusCode == HttpStatusCode.OK)
                {
                    x++;
                }
                log.LogInformation("result: " + result);

                // 결과를 요약하여 출력
                string output = "성공적으로 " + x + "개의 문제를 처리했습니다.";
                log.LogInformation(output);
                responseMessage = output;
            }
        }
        else
        {
            log.LogInformation("프로젝트를 찾을 수 없습니다!");
        }
    }
    catch (Exception ex)
    {
        StringBuilder sb = new StringBuilder();
        sb.Append("{");
        sb.Append("  \"eventType\": \"SnykFindingsErrors\",");
        sb.Append("  \"message\": \"" + ex.Message + "\"");
        sb.Append("}");

        var content = new StringContent(sb.ToString());

        content.Headers.ContentType = new MediaTypeHeaderValue("application/json");

        var NEW_RELIC_INSIGHTS_URL = Environment.GetEnvironmentVariable("NEW_RELIC_INSIGHTS_URL");
        var NEW_RELIC_LICENSE_KEY = Environment.GetEnvironmentVariable("NEW_RELIC_LICENSE_KEY");

        var url = NEW_RELIC_INSIGHTS_URL;
        using var client = new HttpClient();
        client.DefaultRequestHeaders.Add("Api-Key", NEW_RELIC_LICENSE_KEY);
        var response = await client.PostAsync(url, content);

        string result = response.Content.ReadAsStringAsync().Result;
        log.LogInformation("response.StatusCode: " + response.StatusCode);

        responseMessage = "실행 중 오류가 발생했습니다";
    }

    return new OkObjectResult(responseMessage);
}
```

다음 [Azure Function 환경 변수 구성](configure-azure-function-environment-variables.md).