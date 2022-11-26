import api from "./api"

const putApi = async <T, S>(url: string, data: T | {} = {}): Promise<S> => {
  return await api<T, S>('PUT', url, data)
}

export default putApi
