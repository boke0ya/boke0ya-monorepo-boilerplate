import { destroyCookie, setCookie } from "nookies";
import { createContext, useReducer } from "react";
import getApi from "../../libs/fetch/get";
import postApi from "../../libs/fetch/post";
import { LoginRequest, LoginResponse, User } from "../../types/api/user";

interface GlobalState {
  sessionUser?: User;
}

export const GlobalContext = createContext(null)

interface LoadSessionUserAction {
  type: 'LOAD_SESSION_USER';
  sessionUser: User;
}

interface UnloadSessionUserAction {
  type: 'UNLOAD_SESSION_USER';
}

export type GlobalAction = LoadSessionUserAction | 
  UnloadSessionUserAction;

export const GlobalReducer = (state: GlobalState, action: GlobalAction) => {
  switch(action.type){
    case 'LOAD_SESSION_USER':
      return {
        ...state,
        sessionUser: action.sessionUser,
      }
    case 'UNLOAD_SESSION_USER':
      return {
        ...state,
        sessionUser: null,
      }
  }
}

const useGlobalViewModel = () => {
  const [state, dispatch] = useReducer(GlobalReducer, {
    sessionUser: null,
  })
  const login = async (id: string, password: string) => {
    let email: string
    let screenName: string
    if(/^(.+)@(.+)$/.test(id)){
      email = id
    }else{
      screenName = id
    }
    const { token } = await postApi<LoginRequest, LoginResponse>(`/api/login`, {
      email,
      screenName,
      password
    })
    setCookie(null, 'AUTH_TOKEN', token, {
      maxAge: 14 * 24 * 36000,
      path: '/'
    })
  }
  const loadSessionUser = async () => {
    try{
      const sessionUser = await getApi<User>(`/api/session`)
      dispatch({
        type: 'LOAD_SESSION_USER',
        sessionUser,
      })
    }catch(_){ }
  }
  const logout = () => {
    destroyCookie(null, 'AUTH_TOKEN')
    dispatch({
      type: "UNLOAD_SESSION_USER",
    })
  }
  return {
    ...state,
    login,
    logout,
    loadSessionUser,
  }
}

export default useGlobalViewModel;
