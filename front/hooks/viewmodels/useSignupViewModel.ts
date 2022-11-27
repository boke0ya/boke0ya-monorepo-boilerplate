import { useRouter } from "next/router"
import { setCookie } from "nookies"
import { useState } from "react"
import { usePost } from "../../hooks/api"
import useS3Upload from "../../hooks/api/useS3"
import useGlobalViewModel from "../../hooks/viewmodels/useGlobalViewModel"
import { SignupRequest, SignupResponse } from "../../types/api/user"

const useSignupViewModel = () => {
  const globalViewModel = useGlobalViewModel()
  const router = useRouter()
  const [icon, setIcon] = useState<File | Blob>(null)
  const [name, setName] = useState('')
  const [screenName, setScreenName] = useState('')
  const [password, setPassword] = useState('')

  const signupApi = usePost<SignupRequest, SignupResponse>(`/api/signup`)
  const [isLoading, setIsLoading] = useState(false)
  const uploadIconApi = useS3Upload(`/api/users/profile/icon`)

  const error = signupApi.error ?? uploadIconApi.error

  const signup = async () => {
    try{
      setIsLoading(true)
      let emailVerificationToken: string = '';
      if(router.query.token instanceof Array){
        emailVerificationToken = router.query.token[0]
      }else if(typeof emailVerificationToken === 'string'){
        emailVerificationToken = router.query.token
      }
      const { token } = await signupApi.mutate({
        token: emailVerificationToken,
        name,
        screenName,
        password
      })
      setCookie(null, 'AUTH_TOKEN', token, {
        maxAge: 14 * 24 * 3600,
      })
      await uploadIconApi.upload(icon)
      await globalViewModel.loadSessionUser()

    }catch(_){ }
    setIsLoading(false)
  }

  const validateName = () => {
    return name.trim() !== ''
  }
  const validateScreenName = () => {
    return screenName.trim() !== '' && /^[a-zA-Z0-9._-]+$/.test(screenName)
  }
  const validatePassword = () => {
    return password.trim() !== '' && /^\S+$/.test(password)
  }
  const validate = () => {
    return validateName() &&
      validateScreenName() && 
      validatePassword()
  }

  return {
    icon, setIcon,
    name, setName,
    screenName, setScreenName,
    password, setPassword,
    isLoading,
    error,
    signup,
    validate,
  }
}

export default useSignupViewModel
