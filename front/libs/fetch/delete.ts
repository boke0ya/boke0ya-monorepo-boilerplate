import api from "./api"

const deleteApi = async <T, S>(url: string, data: T = null): Promise<S> => {
  return await api<T, S>('DELETE', url, data)
}

export default deleteApi
