export function checkType<T>(val: T): val is T {
    if (val)
        return true
    else
        return false;
}