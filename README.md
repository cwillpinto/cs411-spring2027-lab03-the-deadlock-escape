# CS411 Spring 2027 · Lab 3: The Deadlock Escape

Read the included assignment handout. `Transfer` deliberately takes locks in caller order, so opposing transfers can deadlock. Preserve the public API and account balance invariant while implementing one consistent lock order.

```sh
go test -run '^TestTransfersSequential$' ./...
go test -run '^TestTransfers$' -count=50 -timeout=10s ./...
go test -race ./...
```

The stress test can time out before the repair. Keep the timeout option on exploratory runs. Public tests are representative, not exhaustive.

## Work in your own GitHub repository

The instructor starter is an upstream source. Keep your completed work in a repository under your own GitHub account, then submit the repository URL and the evidence listed in the assignment handout through Blackboard.

If Blackboard links a template repository, choose **Use this template**, create the new repository under your account, and clone that new repository. Check `git remote -v` before pushing; `origin` must be your repository.

For a direct clone, first create an empty repository under your account. Replace the placeholders below with the starter link supplied in Blackboard and your own empty repository URL:

```sh
git clone <instructor-starter-url>
cd <cloned-directory>
git remote rename origin upstream
git remote add origin <your-empty-repository-url>
git remote -v
git push -u origin HEAD
```

Commit and push your completed work to your own `origin`. Do not push to the instructor's upstream repository. Use the exact submission files and verification commands in the assignment handout.
