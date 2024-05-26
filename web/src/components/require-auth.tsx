import { USER_SESSION_COOKIE } from "@/lib/constants";
import Cookies from "js-cookie";
import { Navigate, Outlet } from "react-router-dom";

export const RequireAuth = () => {
  if (!Cookies.get(USER_SESSION_COOKIE)) {
    return <Navigate to={{ pathname: "/login" }} />;
  }
  return <Outlet />;
};
