# Project to analyze github requests

## Intro
Initial goal is to get current PR size and lyfecycle metrics to set goals for future improvements.
According to LinearB Labs' [research](https://www.youtube.com/watch?v=r6v5R7zkbgE), 50% of PRs are idle 81.3% of lifespan (based on a sample of 733.4K PRs, 3.9 M comments and 25.6K Developers).
The expected outcome is to speed up process' cycle time (not individuals) and to increase review depth (average number of comments per PR).


## Pull Size Request Best Practices
- [ ] Break Work into tasks that can be completed in a single sitting
- [ ] Start and finish PRs same day (PRs under 15 min review time get picke up 20x faster)
- [ ] Find a reviewer within one hour
- [ ] Make changes right after comments are provided
- [ ] Keep those big refactors to separate pull requests
- [ ] Establish working agrements with your team (e.g.: max PR size is 200 LOC)
- [ ] Keep those big refactors to separate pull requests

## TODO
- [ ] Parse commits object to calculate time to merge
- [ ] Calculate size
