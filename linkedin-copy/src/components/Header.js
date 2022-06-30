import styled from "styled-components";
import React, { useEffect } from "react";
import Cookies from "js-cookie";
import { useNavigate, Link, createSearchParams } from "react-router-dom";
import axios from "axios";
import { useState } from "react";
import { CSSTransition } from "react-transition-group";
import "./addConnection.css";

const Header = (props) => {
  const [displayNotifs, setDisplayNotifs] = useState(false);
  const [postNotifs, setPostNotifs] = useState([]);
  const [postsLen, setPostsLen] = useState(0);
  const [displayRequests, setDisplayRequests] = useState(false);
  const [requests, setRequests] = useState([]);
  const [requestsLen, setRequestsLen] = useState(0);
  const [search, setSearch] = useState(Object);
  const [displaySearchRes, setDisplaySearchRes] = useState(false);
  const [resolvedRequest, setResolvedRequest] = useState(false);
  let navigate = useNavigate();

  const handleSignOut = (e) => {
    if (Cookies.get("username") !== undefined) {
      Cookies.remove("username");
      navigate("/");
      return;
    }
    Cookies.remove("guest");
    navigate("/");
  };
  const handleMyProfile = (e) => {
    if (Cookies.get("username") !== undefined) {
      navigate("/myProfile");
    }
  };
  const showNotifications = (e) => {
    setDisplayNotifs(!displayNotifs);
    setDisplayRequests(false);
    setDisplaySearchRes(false);
  };
  const showRequests = (e) => {
    setDisplayRequests(!displayRequests);
    setDisplayNotifs(false);
    setDisplaySearchRes(false);
  };
  const handleSearch = async (e) => {
    const res = await axios.get(
      "http://localhost:9090/users/" + e.target.value
    );
    if (res.status === 200) {
      console.log(res.data);
      setSearch(res.data);
      setDisplaySearchRes(true);
      return;
    }
    setSearch(Object);
    setDisplaySearchRes(false);
  };
  useEffect(() => {
    const getPostNotifs = async () => {
      const postRes = await axios.get("http://localhost:9090/notifications", {
        withCredentials: true,
      });
      console.log(postRes.status);
      console.log(postRes.data);
      setPostNotifs(postRes.data);
      setPostsLen(postRes.data.length);
    };
    const getRequests = async () => {
      let requestRes;
      requestRes = await axios.get("http://localhost:9090/requests", {
        withCredentials: true,
      });
      console.log(requestRes.status)
      setRequests(requestRes.data);
      setRequestsLen(requestRes.data.length);
    };
    if (Cookies.get("username") !== undefined) {
      getPostNotifs();
      getRequests();
    }
  }, [displaySearchRes, resolvedRequest]);
  const handleNavToUser = (e) => {
    if(search.Username === Cookies.get("username")){
      navigate("/home");
      return;
    }
    const params = { username: search.Username };
    setDisplaySearchRes(false);
    navigate({
      pathname: "/user/profile",
      search: `?${createSearchParams(params)}`,
    });
    window.location.reload();
  };

  const handleAccept = async (e) => {
    await axios.post(
      "http://localhost:9090/follow-accept/" + e,
      {},
      { withCredentials: true }
    );
    console.log(e);
    setResolvedRequest(true);
  };
  const handleDecline = async (e) => {
    await axios.post(
      "http://localhost:9090/follow-decline" + e,
      {},
      { withCredentials: true }
    );
    setResolvedRequest(true);
  };

  return (
    <Container>
      <Content>
        <Logo>
          <a href="/Home">
            <img src="/images/home-logo-dislinkt.png" alt="" />
          </a>
        </Logo>
        <Search>
          <div>
            <input type="text" placeholder="Search" onChange={handleSearch} />
          </div>
          <SearchIcon>
            <img src="/images/search-icon.svg" alt="" />
          </SearchIcon>
        </Search>
        <CSSTransition
          in={displaySearchRes}
          timeout={200}
          unmountOnExit
          classNames="search"
        >
          <WrapResult>
            <SearchResult onClick={handleNavToUser}>
              <ResTxt>{search.Username}</ResTxt>
            </SearchResult>
          </WrapResult>
        </CSSTransition>
        <Nav>
          <NavListWrap>
            {Cookies.get("username") !== undefined ? (
              <>
                <NavList className="">
                  <a href="/home">
                    <img src="/images/nav-home.svg" alt="" />
                    <span>Home</span>
                  </a>
                </NavList>

                <NavList>
                  <a href="/home">
                    <img src="/images/nav-network.svg" alt="" />
                    <span>My Network</span>
                  </a>
                </NavList>

                <NavList>
                  <a href="/home">
                    <img src="/images/nav-jobs.svg" alt="" />
                    <span>Jobs</span>
                  </a>
                </NavList>

                <NavList>
                  <a href="/home">
                    <img src="/images/nav-messaging.svg" alt="" />
                    <span>Messaging</span>
                  </a>
                </NavList>
                <NavList>
                  <a onClick={showRequests}>
                    <img src="/images/icons8-user-groups-100.png" alt="" />
                    <span>Requests</span>
                    {requestsLen === 0 ? (
                      <></>
                    ) : (
                      <NotificationCount>{requestsLen}</NotificationCount>
                    )}
                  </a>
                </NavList>
                <CSSTransition
                  in={displayRequests}
                  timeout={300}
                  unmountOnExit
                  classNames="box"
                >
                  <Requests>
                    <h1>Requests</h1>
                    <>
                      {!requests ? (
                        <p
                          style={{
                            marginLeft: "40px",
                            marginTop: "10px",
                            paddingBottom: "15px",
                          }}
                        >
                          There are no requests
                        </p>
                      ) : (
                        <NotifBox>
                          {requests.length > 0 &&
                            requests.map((request, key) => (
                              <Request key={key}>
                                <Requester>{request.Requester}</Requester>
                                <Accept
                                  onClick={() => {
                                    handleAccept(request.Requester);
                                  }}
                                >
                                  Accept
                                </Accept>
                                <Decline
                                  onClick={() => {
                                    handleDecline(request.Requester);
                                  }}
                                  value={request.Requester}
                                >
                                  Decline
                                </Decline>
                              </Request>
                            ))}
                        </NotifBox>
                      )}
                    </>
                  </Requests>
                </CSSTransition>

                <NavList>
                  <a onClick={showNotifications}>
                    <img src="/images/nav-notifications.svg" alt="" />
                    <span>Notifications</span>
                    {postsLen === 0 ? (
                      <></>
                    ) : (
                      <NotificationCount>{postsLen}</NotificationCount>
                    )}
                  </a>
                </NavList>
                <CSSTransition
                  in={displayNotifs}
                  timeout={300}
                  unmountOnExit
                  classNames="box"
                >
                  <Notifications>
                    <h1>Notifications</h1>
                    <>
                      {!postNotifs ? (
                        /* <p style={{ marginLeft: "30px", marginTop: "10px", paddingBottom: "20px" }}>
                      There are no notifications
                    </p>*/
                        <NotifBox>
                          <Notification>
                            <FirstDiv>sibin</FirstDiv>
                            <h3>ja volim pero</h3>
                            <img src="/images/home-logo.svg" alt=""></img>
                          </Notification>
                        </NotifBox>
                      ) : (
                        <NotifBox>
                          {postNotifs.length > 0 &&
                            postNotifs.map((post, key) => (
                              <Notification key={key}>
                                <FirstDiv>{post.Username}</FirstDiv>
                                <div>
                                  <h3>{post.TxtContent}</h3>
                                </div>
                                <div>
                                  <img
                                    src={
                                      "http://localhost:9090/images/" +
                                      post.ImageName
                                    }
                                    alt=""
                                  ></img>
                                </div>
                              </Notification>
                            ))}
                        </NotifBox>
                      )}
                    </>
                  </Notifications>
                </CSSTransition>
              </>
            ) : (
              <></>
            )}
            <User>
              <a href="/Header">
                <img src="/images/user.svg" alt="" />
                <span>Me</span>
                <img src="/images/down-icon.svg" alt="" />
              </a>
              <MyProfile>
                <a onClick={handleMyProfile}>My profile</a>
              </MyProfile>
              <SignOut>
                <a onClick={handleSignOut}>Sign Out</a>
              </SignOut>
            </User>
            {Cookies.get("username") !== undefined ? (
              <Work>
                <a href="/Header">
                  <img src="/images/nav-work.svg" alt="" />
                  <span>
                    Work
                    <img src="/images/down-icon.svg" alt="" />
                  </span>
                </a>
              </Work>
            ) : (
              <></>
            )}
          </NavListWrap>
        </Nav>
      </Content>
    </Container>
  );
};
const Requester = styled.div`
  font-weight: 600;
  padding: 2px 2px 2px 2px;
  width: 80px;
  word-break: break-all;
`;
const Accept = styled.div`
  padding: 2px 2px 2px 2px;
  border: 2px solid #0a66c2;
  border-radius: 5px 5px 5px 5px;
  color: #0a66c2;
  &:hover {
    border-color: rgba(10, 102, 194, 0.4);
    color: rgba(10, 102, 194, 0.4);
    cursor: pointer;
  }

  margin-bottom: 4px;
`;
const Decline = styled.div`
  padding: 2px 2px 2px 2px;
  border: 2px solid rgba(177, 0, 30, 1);
  border-radius: 5px 5px 5px 5px;
  margin-right: 10px;
  color: rgba(177, 0, 30, 1);
  &:hover {
    border-color: rgba(177, 0, 30, 0.4);
    color: rgba(177, 0, 30, 0.4);
    cursor: pointer;
  }

  margin-bottom: 4px;
  //margin-right: 5px;
`;
const Container = styled.div`
  background-color: white;
  border-bottom: 1px solid rgba(0, 0, 0, 0.08);
  left: 0;
  padding: 0 24px;
  position: fixed;
  top: 0;
  width: 100vw;
  z-index: 100;
`;

const Content = styled.div`
  display: flex;
  align-items: center;
  margin: 0 auto;
  min-height: 100%;
  max-width: 1128px;
`;

const Logo = styled.span`
  margin-right: 8px;
  font-size: 0px;
`;

const Search = styled.div`
  opacity: 1;
  flex-grow: 1;
  position: relative;
  & > div {
    max-width: 280px;
    input {
      border: none;
      box-shadow: none;
      background-color: #eef3f8;
      border-radius: 2px;
      color: rgba(0, 0, 0, 0.9);
      width: 218px;
      padding: 0 8px 0 40px;
      line-height: 1.75;
      font-weight: 400;
      font-size: 14px;
      height: 34px;
      border-color: #dce6f1;
      vertical-align: text-top;
    }
  }
`;

const SearchIcon = styled.div`
  width: 40px;
  position: absolute;
  z-index: 1;
  top: 10px;
  left: 2px;
  border-radius: 0 2px 2px 0;
  margin: 0;
  pointer-events: none;
  display: flex;
  justify-content: center;
  align-items: center;
`;

const Nav = styled.nav`
  margin-left: auto;
  display: block;
  @media (max-width: 768px) {
    position: fixed;
    left: 0;
    bottom: 0;
    background: white;
    width: 100%;
  }
`;

const NavListWrap = styled.ul`
  display: flex;
  flex-wrap: nowrap;
  list-style-type: none;
  .active {
    span:after {
      content: "";
      transform: scaleX(1);
      border-bottom: 2px solid var(--white, #fff);
      bottom: 0;
      left: 0;
      position: absolute;
      transition: transform 0.2s ease-in-out;
      width: 100%;
      border-color: rgba(0, 0, 0, 0.9);
    }
  }
`;

const NavList = styled.li`
  display: flex;
  align-items: center;
  a {
    align-items: center;
    background: transparent;
    display: flex;
    flex-direction: column;
    font-size: 12px;
    font-weight: 400;
    justify-content: center;
    line-height: 1.5;
    min-height: 52px;
    min-width: 80px;
    position: relative;
    text-decoration: none;
    img {
      width: 30px;
      height: 30px;
    }
    &:hover {
      cursor: pointer;
    }
    span {
      color: rgba(0, 0, 0, 0.6);
      display: flex;
      align-items: center;
    }
    @media (max-width: 768px) {
      min-width: 70px;
    }
  }
  &:hover,
  &:active {
    a {
      span {
        color: rgba(0, 0, 0, 0.9);
      }
    }
  }
`;

const SignOut = styled.div`
  position: absolute;
  top: 107px;
  background: white;
  border-radius: 5px 5px 5px 5px;
  width: 100px;
  height: 40px;
  a {
    font-size: 16px;
    font-weight: 500;
  }
  transition-duration: 167ms;
  text-align: center;
  display: none;
  &:hover {
    cursor: pointer;
    background-color: rgba(0, 0, 0, 0.1);
  }
`;

const MyProfile = styled.div`
  position: absolute;
  top: 67px;
  background: white;
  border-radius: 5px 5px 5px 5px;
  width: 100px;
  height: 40px;
  a {
    font-size: 16px;
    font-weight: 500;
  }
  transition-duration: 167ms;
  text-align: center;
  display: none;
  &:hover {
    cursor: pointer;
    background-color: rgba(0, 0, 0, 0.1);
  }
`;

const User = styled(NavList)`
  a > svg {
    width: 24px;
    border-radius: 50%;
  }
  a > img {
    width: 24px;
    height: 24px;
    border-radius: 50%;
  }
  span {
    display: flex;
    align-items: center;
  }
  &:hover {
    ${SignOut} {
      align-items: center;
      display: flex;
      justify-content: center;
    }
    ${MyProfile} {
      align-items: center;
      display: flex;
      justify-content: center;
    }
  }
`;

const Work = styled(User)`
  border-left: 1px solid rgba(0, 0, 0, 0.08);
`;
const Requests = styled.div`
  border-radius: 5px 5px 5px 5px;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  position: absolute;
  z-index: 10;
  top: 70px;
  right: 380px;
  width: 250px;
  min-height: 30px;
  background-color: white;

  h1 {
    padding-left: 10px;
    padding-top: 10px;
    padding-bottom: 5px;
    font-size: 12px;
    font-weight: 550;
  }
`;
const Notifications = styled.div`
  border-radius: 5px 5px 5px 5px;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  position: absolute;
  z-index: 10;
  top: 70px;
  right: 310px;
  width: 250px;
  background-color: white;

  h1 {
    padding-left: 10px;
    padding-top: 10px;
    padding-bottom: 5px;
    font-size: 12px;
    font-weight: 550;
  }
`;

const NotifBox = styled.div`
  display: flex;
  flex-direction: column;
  width: 250px;
`;
const Request = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  border-top: 2px solid rgba(0, 0, 0, 0.1);
  padding: 3px 3px 3px 3px;
  padding-top: 10px;
  margin-left: 4px;
  margin-right: 4px;
  margin-top: 3px;
  height: 30px;
`;
const Notification = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  border-top: 2px solid rgba(0, 0, 0, 0.1);
  padding: 3px 3px 3px 3px;
  margin-left: 4px;
  margin-right: 4px;
  height: auto;
  p {
    font-size: 9px;
    font-weight: 300;
  }
  h3 {
    font-size: 12px;
    font-weight: 400;
    margin-top: 4px;
  }
  img {
    width: 30px;
    height: 30px;
  }

  &:hover {
    cursor: pointer;
    background-color: rgba(0, 0, 0, 0.08);
  }
`;

const FirstDiv = styled.div`
  align-self: baseline;
  font-size: 16px;
  margin-top: 5px;
`;
const NotificationCount = styled.label`
  position: absolute;
  right: 24px;
  top: -5px;
  font-family: "sans-serif";
  font-weight: 500;
  padding-left: 4px;
  padding-right: 4px;
  z-index: 50;
  border-radius: 50%;
  background-color: rgba(245, 56, 75, 1);
  color: white;
`;
const SearchResult = styled.div`
  background-color: white;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
  border-radius: 0 0 5px 5px;
  width: 266px;
  height: 40px;
  display: flex;
  flex-direction: row;
  transition-duration: 167ms;
  &:hover {
    cursor: pointer;
    background-color: rgba(54, 56, 60, 0.6);
    color: white;
  }
`;
const ResTxt = styled.h1`
  margin-left: 40px;
  align-self: center;
  font-weight: 500;
  text-decoration: none;
`;
const WrapResult = styled.div`
  position: absolute;
  top: 50px;
  left: 286px;
`;
export default Header;
