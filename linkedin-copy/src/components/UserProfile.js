import styled from "styled-components";
import Rightside from "./Rightside";
import React, { useEffect } from "react";
import UserInformation from "./UserInformation";
import UserPosts from "./UserPosts";
import { useLocation } from "react-router-dom";
import { useState } from "react";
import axios from "axios";
import { useSearchParams } from "react-router-dom";
import Cookies from "js-cookie";

const UserProfile = () => {
  const [searchParams] = useSearchParams();
  const [user, setUser] = useState(Object);
  const [posts, setPosts] = useState([]);
  const [following, setFollowing] = useState(false);
  const [loading, setLoading] = useState(true);
  const [guest, setGuest] = useState(false);

  useEffect(() => {
    const user = async () => {
      const userRes = await axios.get(
        "http://localhost:9090/users/" + searchParams.get("username")
      );
      setUser(userRes.data);
      console.log(userRes.data);
      const request =
        "http://localhost:9090/posts/" + searchParams.get("username");
      var res;
      try {
        res = await axios.get(request, { withCredentials: true });
      } catch (err) {
        console.log(err);
      }
      if (res !== undefined) {
        setPosts(res.data);
      }

      if (Cookies.get("username") !== undefined) {
        const followRes = await axios.get(
          "http://localhost:9090/follows/" + searchParams.get("username"),
          { withCredentials: true }
        );
        if (followRes.status === 204) {
          setFollowing(false);
        } else {
          setFollowing(true);
        }
      } else {
        setFollowing(false);
      }
      setLoading(false);
    };
    user();
  }, [searchParams]);
  return (
    <Container>
      <Section></Section>
      <Layout>
        {!loading ? (
          <>
          <UserInformation data={user} followingUser={following} />
          <UserPosts data={posts} userInfo={user} followingUser={following} />
          </>
        ) : (
          <></>
        )}
        <Rightside />
      </Layout>
    </Container>
  );
};

const Container = styled.div`
  padding-top: 52px;
  max-width: 100%;
  height: 100vh;
`;

const Content = styled.div`
  max-width: 1128px;
  margin-left: auto;
  margin-right: auto;
`;

const Section = styled.section`
  min-height: 50px;
  padding: 16px 0;
  box-sizing: content-box;
  text-align: center;
  text-decoration: underline;
  display: flex;
  justify-content: center;
  h5 {
    color: #0a66c2;
    font-size: 14px;
    a {
      font-weight: 700;
    }
  }
  p {
    font-size: 14px;
    color: #434649;
    font-weight: 600;
  }
  @media (max-width: 768px) {
    flex-direction: column;
    padding: 0 5px;
  }
`;

const Layout = styled.div`
  display: grid;
  grid-template-areas: "leftside main rightside";
  grid-template-columns: minmax(0, 5fr) minmax(0, 12fr) minmax(300px, 7fr);
  column-gap: 25px;
  row-gap: 25px;
  /* grid-template-row: auto; */
  margin: 25px 0;
  @media (max-width: 768px) {
    display: flex;
    flex-direction: column;
    padding: 0 5px;
  }
`;

export default UserProfile;
