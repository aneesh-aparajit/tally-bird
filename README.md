# TallyBird
> Implementing and learning about HyperLogLog

- HyperLogLog (HLL) is a __probabilistic algorithm__ for estimating the __cardinality__ (unique count) of a large dataset. This is a approximation, with a very low error on ideal case.
- __Usecases__: This algorithm can be used to answer these questions:
    - How many unique visits has this page had on this day?
    - How many unique users have played this soung?
    - How many unique users have viewed this video?

- __Problem Statement__: How many unique visitors in the exhibition?
    - Writing down all the visitor's full names?
    - Writing down the last digitsof their phone number?
    - Both the above methods require a lot of work.
    - Can we achieve the task by only finger counting? YES

## Flajolet-Martin Algorithm
- __Idea__: Track the longest sequence of leading zeroes you have seen in the 6 digits of phone numbers.
- __Estimation__: The estimation of how many unique users will be close to 10^L, where L is the longest sequence of leading zeroes you found in all the numbers.
    - Instead of counting the zeroes in decimal, we can do it in the binary data.
- We hash the entry , because the criteria we choose to keep track of may not be IID (identical and independantly distributed).
    - If we are using binary outputs, let's say we have a total of `x` leading zeroes, in that case, we will say that roughtly $2^x$ users visited the site.
    - There is also correction factor ($\phi=0.77351$)

$$
\text{Cardinality} = \frac{2^L}{\phi}
$$


## Todo
- [x] Implement Vanilla HyperLogLog
- [ ] Implement Presto
