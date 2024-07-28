export function redirect(url: string){
    window.location.href = url;
}

export function RemoveUrlParameters(paramsToRemove: string[]) {
    const url = new URL(window.location.href);
    paramsToRemove.forEach(param => url.searchParams.delete(param));
    window.history.replaceState({ path: url.href }, '', url.href);
}
