import { useState } from "react"
import { CreatePutUserIconUrlResponse } from "../../types/api/user"
import { ApiError, ApiErrorCode } from "../../types/errors"
import { usePost } from "./usePost"

const useS3Upload = (url: string) => {
  const createUrlApi = usePost<{}, CreatePutUserIconUrlResponse>(url)
  const [error, setError] = useState<ApiError>(null)
  const upload = async (file: File | Blob) => {
    setError(null)
    const { url } = await createUrlApi.mutate({})
    return await fetch(url, {
      method: 'PUT',
      body: file,
      headers: {
        'Content-Type': file.type,
      }
    }).then(res => {
      if(res.ok){
        return
      }else{
        const err = new ApiError({
          code: ApiErrorCode.FailedToPersistBucketObject,
          message: 'Failed to persist bucket object'
        })
        setError(err)
        throw err
      }
    })
  }
  return {
    ...createUrlApi,
    error: error ?? createUrlApi.error,
    upload
  }
}

export default useS3Upload
