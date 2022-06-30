import React from "react";
import { useContext } from "react";
import AuthApi from "./AuthApi";
import { Outlet } from "react-router";
import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";

const ProtectedRoutes = () => {
  const Auth = useContext(AuthApi);
  const username = Cookies.get("username");
  const guest = Cookies.get("guest")
  if (username !== undefined || guest !== undefined) {
    return <Outlet />;
  } else {
    return <Navigate to="/signin" />;
  }
};


export default ProtectedRoutes;
