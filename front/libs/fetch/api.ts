import { parseCookies } from "nookies"
import { ApiError } from "../../types/errors"

const api = async <T, S>(method: 'POST' | 'PUT' | 'GET' | 'DELETE', url: string, data: T = null): Promise<S> => {
  const headers: { [key: string]: string } = {
    'Content-Type': 'application/json'
  }
  const { AUTH_TOKEN: token } = parseCookies()
  if(token){
    headers['Authorization'] = `Bearer ${token}`
  }
  return await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}${url}`, {
    method,
    headers,
    body: data ? JSON.stringify(data) : null,
  }).then(async res => {
    if(!res.ok){
      throw new ApiError(await res.json())
    }
    return res.json()
  }).then((res: S) => {
    return res
  })
}

export default api
