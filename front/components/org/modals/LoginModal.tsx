import { useContext, useState } from "react";
import { GlobalContext } from "../../../hooks/viewmodels/useGlobalViewModel";
import { ApiError } from "../../../types/errors";
import Button from "../../atm/Button";
import ErrorMessage from "../../atm/ErrorMessage";
import FormFooter from "../../mol/FormFooter";
import TextInput from "../../mol/TextInput";
import Modal from "./Modal";

interface LoginModalProps {
  isOpen: boolean;
  onClose?(): void;
}

const LoginModal = ({
  ...props
}: LoginModalProps) => {
  const globalViewModel = useContext(GlobalContext)
  const [error, setError] = useState<ApiError>(null)
  const [id, setId] = useState<string>('')
  const [password, setPassword] = useState<string>('')
  const [isLoading, setIsLoading] = useState(false)
  const login = async () => {
    try{
      setIsLoading(true)
      await globalViewModel.login(id, password)
      onClose()
    }catch(e){
      if(e instanceof ApiError){
        setError(e)
      }
    }
    setIsLoading(false)
  }
  const onClose = () => {
    setId('')
    setPassword('')
    setError(null)
    props.onClose()
  }
  return (
    <Modal {...props} onClose={onClose} title='ログイン'>
      <TextInput
        type='text'
        placeholder="ID または メールアドレス"
        label="ID または メールアドレス"
        value={id}
        onChange={(e) => setId(e.target.value)}
      />
      <TextInput
        type='password'
        placeholder="パスワード"
        label="パスワード"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      {
        error ? (
          <ErrorMessage>{error.getMessage()}</ErrorMessage>
        ) : (<></>)
      }
      <FormFooter>
        <Button isLoading={isLoading} onClick={() => {
          globalViewModel.closeLoginModal()
          globalViewModel.openLoginModal()
        }}>新規登録</Button>
        <Button isLoading={isLoading} onClick={login}>ログイン</Button>
      </FormFooter>
    </Modal>
  )
}

export default LoginModal;
