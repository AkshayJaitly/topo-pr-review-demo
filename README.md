# topo-pr-review-demo

Demo repo for the **pr-review-topo** Claude skill, which reviews a PR in
topological (dependency) order — foundational files first, dependents last.

The open PR ("feat: request-scoped logging + ctx-aware auth") contains two
deliberate bugs that are easy to miss with a bottom-up review but stand out
when files are reviewed foundation-first.
