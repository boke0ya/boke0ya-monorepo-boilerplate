import useSignupViewModel from "../../hooks/viewmodels/useSignupViewModel"
import Button from "../atm/Button"
import ErrorMessage from "../atm/ErrorMessage"
import IconInput from "../mol/IconInput"
import TextInput from "../mol/TextInput"

interface SignupTemplateProps {
  viewmodel: ReturnType<typeof useSignupViewModel>;
}

const SignupTemplate = ({
  viewmodel,
}: SignupTemplateProps) => {
  return (
    <>
      <IconInput
        value={viewmodel.icon}
        defaultUrl={viewmodel.defaultIconUrl}
        onChange={(file) => viewmodel.setIcon(file)}
      />
      <TextInput
        label='名前'
        placeholder='紫式部'
        value={viewmodel.name}
        onChange={(e) => viewmodel.setName(e.target.value)}
      />
      <TextInput
        label='ID'
        placeholder='半角英数字、一部の記号(._-)のみ'
        pattern='^[a-zA-Z0-9._-]+$'
        value={viewmodel.screenName}
        onChange={(e) => viewmodel.setScreenName(e.target.value)}
      />
      <TextInput
        type='password'
        label='パスワード'
        placeholder='パスワード'
        value={viewmodel.password}
        onChange={(e) => viewmodel.setPassword(e.target.value)}
      />
      {
        viewmodel.error ? (
          <ErrorMessage>{viewmodel.error.getMessage()}</ErrorMessage>
        ) : (<></>)
      }
      <Button disabled={!viewmodel.validate()} onClick={() => viewmodel.signup()} isLoading={viewmodel.isLoading}>新規登録</Button>
    </>
  )
}

export default SignupTemplate
