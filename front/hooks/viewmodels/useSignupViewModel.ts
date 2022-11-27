import { useRouter } from "next/router"
import { setCookie } from "nookies"
import { useEffect, useState } from "react"
import { usePost } from "../../hooks/api"
import useS3Upload from "../../hooks/api/useS3"
import useGlobalViewModel from "../../hooks/viewmodels/useGlobalViewModel"
import { SignupRequest, SignupResponse } from "../../types/api/user"
import useIconGenerator from "../useIconGenerateor"

const useSignupViewModel = () => {
  const globalViewModel = useGlobalViewModel()
  const router = useRouter()
  const [icon, setIcon] = useState<File | Blob>(null)
  const [defaultIconUrl, setDefaultIconUrl] = useState<string>(null)
  const [name, setName] = useState('')
  const [screenName, setScreenName] = useState('')
  const [password, setPassword] = useState('')

  const iconGenerator = useIconGenerator()
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
        path: '/'
      })
      await uploadIconApi.upload(icon ?? iconGenerator.icon)
      await globalViewModel.loadSessionUser()
      router.push('/')
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

  useEffect(() => {
    if(iconGenerator.icon){
      setDefaultIconUrl(URL.createObjectURL(iconGenerator.icon))
    }
    return () => {
      if(defaultIconUrl){
        URL.revokeObjectURL(defaultIconUrl)
      }
    }
  }, [iconGenerator.icon])

  return {
    icon, setIcon,
    defaultIconUrl,
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
