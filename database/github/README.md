# github

This github database is used to write(set) and read(get) files to a github repository. The URL scheme doesn't require a hostname, it defaults to `github.com`.

Authenticated URL: `github://user:personal-access-token@owner/repo/path#ref`

Unauthenticated URL: `github://owner/repo/path#ref`

| URL Query  | Description                                                                 |
|------------|-----------------------------------------------------------------------------|
| user | (optional) The username of the user connecting                              |
| personal-access-token | (optional) An access token from GitHub (https://github.com/settings/tokens) |
| owner | the repo owner/organization                                                 |
| repo | the name of the repository                                                  |
| path | path in repo to read and write                                              |
| ref | (optional) can be a SHA, branch, or tag                                     |

You must use an authenticated URL if you want to,
1. Write/Set
2. The repo doesn't exist. Repo will be created if it doesn't exist
