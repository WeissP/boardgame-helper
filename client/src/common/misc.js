const objectMap = (obj, fn) =>
    Object.fromEntries(
        Object.entries(obj).map(
            ([k, v]) => [k, fn(v, k)]
        )
    )

const updateArray = (f, x, index) => {
    f((prev) => prev.map((el, i) => (i !== index ? el : x)))
}

export { objectMap, updateArray }
