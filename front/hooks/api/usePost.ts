import useApi from "./useApi"

export const usePost = <T, S>(url: string) => {
  return useApi<T, S>('POST', url)
}
