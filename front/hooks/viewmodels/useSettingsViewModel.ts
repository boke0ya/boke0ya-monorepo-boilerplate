import { useContext, useEffect, useState } from "react";
import {
  UpdateEmailRequest,
  UpdateEmailResponse,
  UpdatePasswordRequest,
  UpdatePasswordResponse,
  UpdateProfileRequest,
  UpdateProfileResponse,
} from "@/types/api/user";
import { usePost, usePut } from "../api";
import useS3Upload from "../api/useS3";
import { GlobalContext } from "./useGlobalViewModel";

const useSettingsViewModel = () => {
  const globalViewModel = useContext(GlobalContext);
  const updateEmailApi = usePost<UpdateEmailRequest, UpdateEmailResponse>(`/api/users/email`);
  const [email, setEmail] = useState("");
  const updateEmail = async () => {
    await updateEmailApi.mutate({
      email,
    });
  };

  const updatePasswordApi = usePut<UpdatePasswordRequest, UpdatePasswordResponse>(
    `/api/users/password`
  );
  const [currentPassword, setCurrentPassword] = useState("");
  const [password, setPassword] = useState("");
  const updatePassword = async () => {
    await updatePasswordApi.mutate({
      currentPassword,
      password,
    });
  };

  const updateProfileApi = usePut<UpdateProfileRequest, UpdateProfileResponse>(
    `/api/users/profile`
  );
  const updateIconApi = useS3Upload(`/api/users/profile/icon`);
  const [screenName, setScreenName] = useState("");
  const [name, setName] = useState("");
  const [icon, setIcon] = useState<File | Blob | null>(null);
  const updateProfile = async () => {
    const updateProfilePromise = updateProfileApi.mutate({
      screenName,
      name,
    });
    const updateIconPromise = (async () => {
      if (icon) {
        updateIconApi.upload(icon);
      }
    })();
    await Promise.all([updateProfilePromise, updateIconPromise]);
    await globalViewModel.loadSessionUser();
  };
  useEffect(() => {
    if (globalViewModel.sessionUser) {
      setScreenName(globalViewModel.sessionUser.screenName);
      setName(globalViewModel.sessionUser.name);
    }
  }, [globalViewModel.sessionUser]);
  return {
    isLoadingUpdateEmail: updateEmailApi.isLoading,
    isLoadingUpdatePassword: updatePasswordApi.isLoading,
    isLoadingUpdateProfile: updateProfileApi.isLoading,
    updateEmail,
    updatePassword,
    updateProfile,
    emailError: updateEmailApi.error,
    passwordError: updatePasswordApi.error,
    profileError: updateProfileApi.error,
    email,
    setEmail,
    password,
    setPassword,
    currentPassword,
    setCurrentPassword,
    name,
    setName,
    screenName,
    setScreenName,
    icon,
    setIcon,
  };
};

export default useSettingsViewModel;
