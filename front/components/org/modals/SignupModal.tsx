import { useState } from "react";
import { usePost } from "../../../hooks/api";
import { SignupEmailVerificationRequest } from "../../../types/api/user";
import Button from "../../atm/Button";
import ErrorMessage from "../../atm/ErrorMessage";
import TextInput from "../../mol/TextInput";
import Modal, { ModalFooter } from "./Modal"

interface SignupModalProps {
  isOpen: boolean;
  onClose?(): void;
}

const SignupModal = ({
  isOpen,
  onClose,
}: SignupModalProps) => {
  const [email, setEmail] = useState('')
  const signupApi = usePost<SignupEmailVerificationRequest, null>(`/api/login`)
  const signup = async () => {
    signupApi.mutate({
      email
    })
  }
  return (
    <Modal isOpen={isOpen} onClose={onClose} title='新規登録'>
      <TextInput
        type='email'
        placeholder='メールアドレス'
        label='メールアドレス'
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />
      {
        signupApi.error ? (
          <ErrorMessage>{signupApi.error.getMessage()}</ErrorMessage>
        ) : (<></>)
      }
      <ModalFooter>
        <span />
        <Button onClick={signup} isLoading={signupApi.isLoading}>確認メールを送信</Button>
      </ModalFooter>
    </Modal>
  )
}

export default SignupModal
