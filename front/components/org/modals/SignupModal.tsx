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
  ...props
}: SignupModalProps) => {
  const [email, setEmail] = useState('')
  const [isSucceed, setIsSucceed] = useState(false)
  const signupApi = usePost<SignupEmailVerificationRequest, null>(`/api/email-verification/signup`)
  const signup = async () => {
    try{
      await signupApi.mutate({
        email
      })
      setIsSucceed(true)
    }catch(_){}
  }
  const onClose = () => {
    setEmail('')
    props.onClose()
  }
  return (
    <Modal isOpen={isOpen} onClose={onClose} title='新規登録'>
      {(() => {
        if(!isSucceed){
          return (
            <>
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
                <Button disabled={email.length === 0} onClick={signup} isLoading={signupApi.isLoading}>確認メールを送信</Button>
              </ModalFooter>
            </>
          )
        }else{
          return (
            <>
              <h2>確認メールを送信しました！</h2>
              <p>受信ボックスを確認し、メールに記載されたURLから本登録を行って下さい</p>
              <ModalFooter>
                <span />
                <Button onClick={onClose}>閉じる</Button>
              </ModalFooter>
            </>
          )
        }
      })()}
    </Modal>
  )
}

export default SignupModal
