# Project Conventions

## Repository Scope

- Keep changes for this repository within `background/`.
- Redirect root-level, server, and mobile work to the corresponding project conversation.

## Git Operations

- Do not perform any Git operation unless the user explicitly requests it. This includes read-only inspection commands such as `status`, `log`, and `diff`, as well as `fetch`, `checkout`, `branch`, `merge`, `commit`, `push`, and similar operations.
- Commit only when the user explicitly asks. Never commit or push without explicit user instruction.
- When Git work is authorized, follow Git Flow principles.
- Write all commit messages in English.

## Language and Naming

- Write all documentation in English, including README files and other document files.
- Use English for identifiers and names, including variables, functions, types, files, and resources. Follow the casing and formatting conventions established by the project and implementation language.
- Use Turkish only for user-facing language inside the application. Do not use Turkish for identifiers, technical names, documentation, or commit messages.
- Apply Clean Code Architecture naming principles so names communicate meaning and responsibility while remaining consistent with existing project and language conventions.
- Name functions clearly according to their action and responsibility.

## Code Structure

- Keep functions short, focused, understandable, and limited to one coherent responsibility. Avoid excessive context and unrelated responsibilities; use engineering judgment rather than a rigid line-count rule.

## Go Project Layout

- For Go projects created or changed in this workspace, follow the patterns from [golang-standards/project-layout](https://github.com/golang-standards/project-layout).
- Put executable entry points under `cmd/<app>` and private application code under `internal/`.
- Add optional directories such as `pkg`, `api`, or `deployments` only when the project needs them.
