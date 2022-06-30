import React from "react";
import { useContext } from "react";
import AuthApi from "./AuthApi";
import { Outlet } from "react-router";
import { Navigate } from "react-router-dom";
import Cookies from "js-cookie";

const AuthedRoute = () => {
  const Auth = useContext(AuthApi);
  const username = Cookies.get("username");
  if (username !== undefined) {
    return <Navigate to="/home" />;
  } else {
    return <Outlet />;
  }
};

export default AuthedRoute;
