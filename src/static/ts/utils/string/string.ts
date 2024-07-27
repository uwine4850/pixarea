export function removePrefix(str: string, prefix: string): string {
    if (str.startsWith(prefix)) {
        return str.substring(prefix.length);
    }
    return str;
}