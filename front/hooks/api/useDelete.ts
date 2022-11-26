import useApi from "./useApi"

export const useDelete = <T, S>(url: string) => {
  return useApi<T, S>('DELETE', url)
}
