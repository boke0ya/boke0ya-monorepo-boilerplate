import { parseCookies } from "nookies"
import { ApiError } from "../../types/errors"

const api = async <T, S>(method: 'POST' | 'PUT' | 'GET' | 'DELETE', url: string, data: T | {} = {}): Promise<S> => {
  const headers: { [key: string]: string } = {
    'Content-Type': 'application/json'
  }
  const { BEATHUB_SESSION_ID: token } = parseCookies()
  if(token){
    headers['Authorization'] = `Bearer ${token}`
  }
  return await fetch(`${process.env.NEXT_PUBLIC_BASE_URL}${url}`, {
    method,
    headers,
    body: JSON.stringify(data)
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
