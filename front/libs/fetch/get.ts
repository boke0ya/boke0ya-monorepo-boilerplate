import api from "./api"

const getApi = async <S>(url: string): Promise<S> => {
  return await api<null, S>('GET', url)
}

export default getApi
