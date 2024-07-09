export function RemoveUrlParameters(paramsToRemove: string[]) {
    const url = new URL(window.location.href);
    paramsToRemove.forEach(param => url.searchParams.delete(param));
    window.history.replaceState({ path: url.href }, '', url.href);
}
