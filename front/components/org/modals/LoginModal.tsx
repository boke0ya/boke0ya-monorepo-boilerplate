import { ComponentProps, useState } from "react";
import useGlobalViewModel from "../../../hooks/viewmodels/useGlobalViewModel";
import { ApiError } from "../../../types/errors";
import Button from "../../atm/Button";
import ErrorMessage from "../../atm/ErrorMessage";
import TextInput from "../../mol/TextInput";
import Modal, { ModalFooter } from "./Modal";

interface LoginModalProps {
  isOpen: boolean;
  onClose?(): void;
}

const LoginModal = ({
  ...props
}: LoginModalProps) => {
  const globalViewModel = useGlobalViewModel()
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
      <ModalFooter>
        <span />
        <Button isLoading={isLoading} onClick={login}>ログイン</Button>
      </ModalFooter>
    </Modal>
  )
}

export default LoginModal;
