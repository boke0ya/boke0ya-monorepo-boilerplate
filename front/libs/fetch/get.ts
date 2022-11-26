import api from "./api"

const getApi = async <S>(url: string): Promise<S> => {
  return await api<{}, S>('GET', url)
}

export default getApi
