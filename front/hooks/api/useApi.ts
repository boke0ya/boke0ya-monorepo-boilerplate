import { useState } from "react"
import api from "../../libs/fetch/api"
import { ApiError } from "../../types/errors"

const useApi = <T, S>(method: 'POST' | 'GET' | 'PUT' | 'DELETE', url: string) => {
  const [result, setResult] = useState<{
    data?: S,
    error?: ApiError,
  }>({})
  const mutate = async (data: T) => {
    setResult({})
    try{
      const result = await api<T, S>(method, url, {
        ...data
      })
      setResult({
        data: result
      })
    }catch(e){
      if(e instanceof ApiError){
        setResult({
          error: e
        })
      }
    }
  }
  return {
    ...result,
    isLoading: !result.data && !result.error,
    mutate,
  }
}

export default useApi
