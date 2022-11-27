import { useState } from "react"
import api from "../../libs/fetch/api"
import { ApiError } from "../../types/errors"

const useApi = <T, S>(method: 'POST' | 'GET' | 'PUT' | 'DELETE', url: string) => {
  const [result, setResult] = useState<{
    isLoading: boolean;
    data?: S;
    error?: ApiError;
  }>({
    isLoading: false,
  })
  const mutate = async (data: T) => {
    setResult({
      isLoading: true,
    })
    try{
      const result = await api<T, S>(method, url, {
        ...data
      })
      setResult({
        isLoading: false,
        data: result
      })
      return result
    }catch(e){
      if(e instanceof ApiError){
        setResult({
          isLoading: false,
          error: e
        })
      }
      throw e
    }
  }
  return {
    ...result,
    mutate,
  }
}

export default useApi
