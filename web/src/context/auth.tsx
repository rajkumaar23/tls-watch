import { API } from "@/lib/utils";
import { useState, ReactNode, useEffect, createContext } from "react";
import { Navigate, Outlet, useLocation, useNavigate } from "react-router-dom";
import Cookies from "js-cookie";
import { AUTH_COOKIE } from "@/lib/constants";
import { User } from "@/lib/types";

type AuthProviderProps = {
  children: ReactNode;
};

export type AuthContextType = {
  user: User | null;
  setUser: React.Dispatch<React.SetStateAction<User | null>>;
};

export const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState<User | null>(null);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await API.get("/auth/me");
        setUser(response.data.profile);
      } catch (error) {
        if (location.pathname != "/login") {
          Cookies.remove(AUTH_COOKIE, {
            domain: window.location.hostname,
            path: "/",
            secure: true,
          });
          navigate("/login");
        }
        console.error(error);
      }
    };

    fetchUserData();
  }, [navigate, location.pathname]);

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
