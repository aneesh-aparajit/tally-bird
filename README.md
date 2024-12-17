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

he HyperLogLog (HLL) is a probablistic data structure that tracks the cardinality of large data sets. HyperLogLog is suited for scenarios like the above, where the goal is to count the number of unique items in a massive data stream without explicitly storing every item. HLL relies on a clever hashing mechanism and a compact data structure to provide accurate estimates of unique users while using only a fraction of the memory required by traditional methods. This makes HLL an essential tool in modern database analytics.

HLL provides probabilistic counting mechanism based on the following parameters:

- `b` - Number of initial bits in a binary representation of a hash value
- `m` - number of registers (or also called buckets) - can be considered as a memory block. They are equal to 2^b. (The terms "buckets" and "registers" can be used interchangeably when discussing HyperLogLog and tasks).
- `p` - leftmost position of 1 (MSBs position of 1)


Consider a simple example of how this algorithm works using the string "A great database is a great life". First, the string is hashed to produce a hash value, which is then converted into its binary representation. From the hash value (binary form), b bits are extracted, starting from the most significant bit(MSB). The register value is calculated from the extracted bits. (by default each register has a value of 0).

$$
\text{Cardinality} = \alpha*m*\frac{m}{\sum_{i=1}^N2^{-M[i]}}
$$

## Todo
- [x] Implement Vanilla HyperLogLog
- [ ] Implement Presto
