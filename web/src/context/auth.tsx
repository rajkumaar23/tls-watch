import { API } from "@/lib/utils";
import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
} from "react";
import { Navigate, Outlet } from "react-router-dom";
import Cookies from "js-cookie";
import { AUTH_COOKIE } from "@/lib/constants";

export type User = {
  id: number;
  oidc_subject: string;
  name: string;
  picture: string;
  created_at: string;
  updated_at: string;
};

type AuthContextType = {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
};

const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

type AuthProviderProps = {
  children: ReactNode;
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await API.get("/auth/me");
        setUser(response.data.profile);
      } catch (error) {
        console.error("error fetching user data:", error);
      }
    };

    if (!user) {
      fetchUserData();
    }
  });

  return (
    <AuthContext.Provider value={{ user, setUser }}>
      {children}
    </AuthContext.Provider>
  );
};

export const RequireAuth = () => {
  if (!Cookies.get(AUTH_COOKIE)) {
    return <Navigate to={{ pathname: "/login" }} />;
  }
  return <Outlet />;
};
