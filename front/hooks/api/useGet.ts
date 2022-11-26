import useApi from "./useApi"

export const useGet = <T, S>(url: string) => {
  return useApi<T, S>('GET', url)
}
