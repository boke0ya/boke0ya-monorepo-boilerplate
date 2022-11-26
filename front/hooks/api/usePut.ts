import useApi from "./useApi"

export const usePut = <T, S>(url: string) => {
  return useApi<T, S>('PUT', url)
}
