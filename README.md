# Time Stat
Keep stats on your time.

## Dimensions

1. in-general
2. today
3. this week
4. this month
5. this year
6. this day of the week
7. this 10-minute time
8. this 10-minute time on this day of the week

## Resolution

- minutes up to 24 hours (1440)
- second up to 1 minute (60)
- 1,500 possible outcomes

## Algorithm

1. Timer creates singleton with start time
2. Timer writes end time and..
3. Transactional insertion of reaper task with guid
4. Reaper task updates dimensions 1-8 if guid doesn't match (idempotence)
5. Deletes timer (reset)

Each dimension is a Bayesian probability distribution over the 1,500 possible
outcomes ([algorithm](http://www.databozo.com/2013/09/15/Bayesian_updating_of_probability_distributions.html)).

Each dimension also records the number of contributing samples, average, median and the standard deviation.  Dimensions 2-5 accumulate history (one Datastore entry per day, etc...)  Dimensions 1 and 6-8 are singletons and accumulate over the lifetime of the timer.

## Limitations

Timers are namespaced by who is timing and what they are timing.  A user cannot run multiple timers simultaneously, even if they are timing different things.
