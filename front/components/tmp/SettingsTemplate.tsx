import { useContext } from "react";
import { GlobalContext } from "@/hooks/viewmodels/useGlobalViewModel";
import useSettingsViewModel from "@/hooks/viewmodels/useSettingsViewModel";
import Button from "@atm/Button";
import ErrorMessage from "@atm/ErrorMessage";
import NormalLayout from "@layout/NormalLayout";
import FormFooter from "@mol/FormFooter";
import IconInput from "@mol/IconInput";
import TextInput from "@mol/TextInput";
import Box from "@utils/Box";

interface SettingsTemplateProps {
  viewmodel: ReturnType<typeof useSettingsViewModel>;
}

const SettingsTemplate = ({ viewmodel }: SettingsTemplateProps) => {
  const globalViewModel = useContext(GlobalContext);
  return (
    <NormalLayout>
      <h1>Settings</h1>
      <Box>
        <h2>Profile</h2>
        <IconInput
          defaultUrl={globalViewModel.sessionUser?.iconUrl}
          value={viewmodel.icon}
          onChange={(icon) => viewmodel.setIcon(icon)}
        />
        <TextInput
          label="ID"
          placeholder="Only A to Z and a to z and some symbols(._-) allowed"
          pattern="^[a-zA-Z0-9._-]+$"
          value={viewmodel.screenName}
          onChange={(e) => viewmodel.setScreenName(e.target.value)}
        />
        <TextInput
          label="Name"
          placeholder="Ameno Uzume"
          value={viewmodel.name}
          onChange={(e) => viewmodel.setName(e.target.value)}
        />
        {viewmodel.profileError ? (
          <ErrorMessage>{viewmodel.profileError.getMessage()}</ErrorMessage>
        ) : (
          <></>
        )}
        <FormFooter>
          <span />
          <Button onClick={() => viewmodel.updateProfile()}>Save Profile</Button>
        </FormFooter>
      </Box>
      <Box>
        <h2>E-mail</h2>
        <TextInput
          label="E-mail"
          placeholder="ameno.uzume@example.com"
          type="email"
          value={viewmodel.email}
          onChange={(e) => viewmodel.setEmail(e.target.value)}
        />
        {viewmodel.emailError ? (
          <ErrorMessage>{viewmodel.emailError.getMessage()}</ErrorMessage>
        ) : (
          <></>
        )}
        <FormFooter>
          <span />
          <Button onClick={() => viewmodel.updateEmail()}>Send confirmation mail</Button>
        </FormFooter>
      </Box>
      <Box>
        <h2>Password</h2>
        <TextInput
          label="Current Password"
          type="password"
          value={viewmodel.password}
          onChange={(e) => viewmodel.setPassword(e.target.value)}
        />
        <TextInput
          label="New Password"
          type="password"
          value={viewmodel.password}
          onChange={(e) => viewmodel.setPassword(e.target.value)}
        />
        {viewmodel.passwordError ? (
          <ErrorMessage>{viewmodel.passwordError.getMessage()}</ErrorMessage>
        ) : (
          <></>
        )}
        <FormFooter>
          <span />
          <Button onClick={() => viewmodel.updatePassword()}>Change Password</Button>
        </FormFooter>
      </Box>
    </NormalLayout>
  );
};

export default SettingsTemplate;
