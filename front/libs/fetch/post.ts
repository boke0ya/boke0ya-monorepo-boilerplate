import api from "./api"

const postApi = async <T, S>(url: string, data: T = null): Promise<S> => {
  return await api<T, S>('POST', url, data)
}

export default postApi
