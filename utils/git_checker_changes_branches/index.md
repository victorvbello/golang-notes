## Git checker changes branches

#### Github Toke `TOKEN_GITHUB`
> Add your [Personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
#### Output
```
token <TOKEN_GITHUB>
organization name <ORGANIZATION_NAME>
Process [ORG/REPO_NAME_1]
Process [ORG/REPO_NAME_2]
.
.
.
Process [ORG/REPO_NAME_N]
----
Repos ready to Prod
[ORG/REPO_NAME_1] https://github.com/ORG/REPO_NAME_1/compare/main...develop
[ORG/REPO_NAME_2] https://github.com/ORG/REPO_NAME_2/compare/main...develop
.
.
.
[ORG/REPO_NAME_N] https://github.com/ORG/REPO_NAME_N/compare/main...develop
----
```

#### Run

```sh
go run git_checker_changes_branches/main.go "<ORGANIZATION_NAME>" "<TOKEN_GITHUB>"
```