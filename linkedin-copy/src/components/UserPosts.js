import styled from "styled-components";
import React, { useEffect } from "react";
import { useState } from "react";
import axios from "axios";
import Cookies from "js-cookie";
import { CSSTransition } from "react-transition-group";

const UserPosts = (props) => {
  const [displayComments, setDisplayComments] = useState([false]);
  const [content, setContent] = useState("");
  const [likeChanged, setLikeChanged] = useState(false);
  const [posts, setPosts] = useState([]);
  const [user, setUser] = useState(Object);

  const handleLike = (idPost) => {
    let content = true;
    console.log(idPost);
    axios.put(
      "http://localhost:9090/like",
      JSON.stringify({ idPost, content }),
      { withCredentials: true }
    );
    setLikeChanged(!likeChanged);
  };
  const handleContent = (e) => {
    setContent(e.target.value);
  };
  const handleLikeIcon = (likes) => {
    let tmp = Cookies.get("username");
    if (likes === undefined) {
      return "/images/icons8-like-icon.png";
    }
    for (let i = 0; i < likes.length; i++) {
      if (likes[i].Author === tmp) {
        return "/images/icons8-like-64.png";
      }
    }
    return "/images/icons8-like-icon.png";
  };
  const handlePostComment = (idPost) => {
    console.log(idPost);
    axios.put(
      "http://localhost:9090/comment",
      JSON.stringify({ idPost, content }),
      {
        withCredentials: true,
      }
    );
    setContent("");
    document.getElementById("commentContent").value = "";
    const elems = document.getElementsByClassName("addingComments");
    for (let i = 0; i < elems.length; i++) {
      elems.item(i).value = "";
    }
  };

  const handleDisplayComments = (postKey) => {
    if (displayComments[postKey] === undefined) {
      for (let i = displayComments.length - 1; i < postKey; i++) {
        displayComments.push(false);
        console.log(i);
      }
    }
    console.log(displayComments[postKey]);
    setDisplayComments((prevState) =>
      prevState.map((item, idx) => (idx === postKey ? !item : item))
    );
  };

  useEffect(() => {
    const images = async () => {
      const userRes = await axios.get(
        "http://localhost:9090/users/" + props.userInfo.username
      );
      console.log(userRes.data);
      setUser(userRes.data);
      const request = "http://localhost:9090/posts/" + props.userInfo.Username;
      var res;
      try {
        res = await axios.get(request, { withCredentials: true });
      } catch (err) {
        console.log(err);
      }
      if (res !== undefined) {
        setPosts(res.data);
        console.log(res.data);
        return;
      }
    };
    images();
  }, [content, likeChanged]);

  return (
    <Container>
      <>
        {!props.userInfo.Public && !props.followingUser ? (
          <NoContentPrompt>
            <img src="/images/icons8-private-lock-90.png" alt=""></img>
            <p>This Account is Private</p>
            <h3>Follow this account to see their posts</h3>
          </NoContentPrompt>
        ) : !props.data ? (
          <NoContentPrompt>
            <p>There are no articles</p>
          </NoContentPrompt>
        ) : (
          <Content>
            {posts.length > 0 &&
              posts.map((post, key) => (
                <Article key={key}>
                  <SharedActor>
                    <a>
                      <img src="/images/user.svg" alt="" />
                      <div>
                        <span>{post.Username}</span>
                        <span>{user.Email}</span>
                        <span>27/6/2022</span>
                      </div>
                    </a>
                    <button>
                      <img src="/images/icons8-ellipsis-60.png" alt="" />
                    </button>
                  </SharedActor>
                  <Description>{post.TxtContent}</Description>
                  <SharedImg>
                    <a>
                      <img
                        src={"http://localhost:9090/images/" + post.ImageName}
                        alt=""
                      />
                    </a>
                  </SharedImg>
                  <SocialCounts>
                    <li>
                      <button>
                        <img src="/images/icons8-like-64.png" alt="" />
                        {post.Likes === undefined ? (
                          <span>0</span>
                        ) : (
                          <span>{post.Likes.length}</span>
                        )}
                      </button>
                    </li>
                    <li>
                      <button>
                        <img src="/images/icons8-comment-icon.png" alt="" />
                        {post.Comments === undefined ? (
                          <span>0</span>
                        ) : (
                          <span>{post.Comments.length}</span>
                        )}
                      </button>
                    </li>
                  </SocialCounts>
                  {Cookies.get("username") !== undefined ? (
                    <>
                      <SocialActions>
                        <button
                          onClick={() => {
                            handleLike(post.ID);
                          }}
                        >
                          <img src={handleLikeIcon(post.Likes)} alt="" />
                          <span>Like</span>
                        </button>
                        <button onClick={() => handleDisplayComments(key)}>
                          <img src="/images/icons8-comment-58.png" alt="" />
                          <span>Comments</span>
                        </button>
                        <button>
                          <img src="/images/icons8-share-48.png" alt="" />
                          <span>Share</span>
                        </button>
                      </SocialActions>
                      <AddComment>
                        <img src="/images/user.svg" alt="" />
                        <input
                          id="commentContent"
                          className="addingComments"
                          type="text"
                          placeholder=" Add comment..."
                          onChange={handleContent}
                        ></input>
                        <button onClick={() => handlePostComment(post.ID)}>
                          Post
                        </button>
                      </AddComment>
                    </>
                  ) : (
                    <></>
                  )}
                  <CSSTransition
                    in={displayComments[key]}
                    timeout={300}
                    unmountOnExit
                    classNames="comment"
                  >
                    <Comments>
                      {post.Comments !== undefined &&
                        post.Comments.map((comment, key) => (
                          <Comment key={key}>
                            <UserImg>
                              <img src="/images/user.svg" alt="" />
                              <CommentInfo>
                                <h5>{comment.Author}</h5>
                              </CommentInfo>
                            </UserImg>
                            <CommentContent>{comment.Content}</CommentContent>
                          </Comment>
                        ))}
                    </Comments>
                  </CSSTransition>
                </Article>
              ))}
          </Content>
        )}
      </>
    </Container>
  );
};

const Container = styled.div`
  grid-area: main;
`;

const CommonCard = styled.div`
  text-align: center;
  overflow: hidden;
  margin-bottom: 8px;
  background-color: #fff;
  border-radius: 5px;
  position: relative;
  border: none;
  box-shadow: 0 0 0 1px rgb(0 0 0 / 15%), 0 0 0 rgb(0 0 0 / 20%);
`;
const Article = styled(CommonCard)`
  padding: 0;
  margin: 0 0 8px;
  overflow: visible;
`;
const ShareBox = styled(CommonCard)`
  display: flex;
  flex-direction: column;
  color: #958b7b;
  margin: 0 0 8px;
  background: white;
  div {
    button {
      outline: none;
      color: rgba(0, 0, 0, 0.6);
      font-size: 14px;
      line-height: 1.5;
      min-height: 48px;
      background: transparent;
      border: none;
      display: flex;
      align-items: center;
      font-weight: 600;
      &:hover {
        cursor: pointer;
      }

      img {
        width: 30px;
        height: 30px;
      }
    }

    &:first-child {
      display: flex;
      align-items: center;
      padding: 8px 16px 0px 16px;
      img {
        width: 48px;
        border-radius: 50%;
        margin-right: 8px;
      }
      button {
        margin: 4px 0;
        flex-grow: 1;
        border-radius: 35px;
        padding-left: 16px;
        border: 1px solid rgba(0, 0, 0, 0.15);
        background-color: white;
        text-align: left;
      }
    }
    &:nth-child(2) {
      display: flex;
      flex-wrap: wrap;
      justify-content: space-around;
      padding-bottom: 4px;

      button {
        img {
          margin: 0 4px 0 -2px;
        }
        span {
          color: #70b5f9;
        }
      }
    }
  }
`;

const SharedActor = styled.div`
  padding-right: 40px;
  flex-wrap: nowrap;
  padding: 12px 16px 0;
  margin-bottom: 8px;
  align-items: center;
  display: flex;
  a {
    margin-right: 12px;
    flex-grow: 1;
    overflow: hidden;
    display: flex;
    text-decoration: none;

    img {
      width: 48px;
      height: 48px;
    }
    & > div {
      display: flex;
      flex-direction: column;
      flex-grow: 1;
      flex-basis: 0;
      margin-left: 8px;
      overflow: hidden;
      span {
        text-align: left;
        &:first-child {
          font-size: 14px;
          font-weight: 700;
          color: rgba(0, 0, 0, 1);
        }

        &:nth-child(n + 1) {
          font-size: 12px;
          color: rgba(0, 0, 0, 0.6);
        }
      }
    }
  }

  button {
    position: absolute;
    right: 12px;
    top: 0;
    background: transparent;
    border: none;
    outline: none;
    img {
      width: 20px;
      height: 20px;
    }
  }
`;

const Description = styled.div`
  padding: 0 16px;
  overflow: hidden;
  color: rgba(0, 0, 0, 0.9);
  font-size: 14px;
  text-align: left;
`;

const SharedImg = styled.div`
  margin-top: 8px;
  width: 100%;
  display: block;
  position: relative;
  background-color: #f9fafb;
  img {
    object-fit: fill;
    width: 100%;
    height: 400px;
  }
`;

const SocialCounts = styled.ul`
  line-height: 1.3;
  display: flex;
  align-items: flex-start;
  overflow: auto;
  margin: 0 16px;
  padding: 8px 0;
  border-bottom: 1px solid #e9e5df;
  list-style: none;
  li {
    margin-right: 5px;
    font-size: 12px;
    button {
      display: flex;
      border: none;
      background-color: white;
      img {
        width: 15px;
        height: 15px;
        margin-right: 3px;
      }
      span {
        margin-left: 5px;
        font-size: 15px;
        font-weight: 600;
        color: grey;
      }
    }
  }
`;

const SocialActions = styled.div`
  align-items: center;
  display: flex;
  justify-content: flex-start;
  margin: 0;
  min-height: 40px;
  padding: 4px 8px;
  button {
    display: inline-flex;
    align-items: center;
    padding: 8px;
    color: #0a66c2;
    border: none;
    background-color: white;
    @media (min-width: 768px) {
      margin-left: 8px;
    }
    &:hover {
      cursor: pointer;
    }

    img {
      width: 20px;
      height: 20px;
    }
  }
`;
const NoContentPrompt = styled(CommonCard)`
  min-height: 50px;
  p {
    font-size: 25px;
    font-weight: 500;
    margin-top: 10px;
    margin-bottom: 10px;
  }
  h3 {
    font-size: 15px;
    font-weight: 400;
    margin-bottom: 20px;
  }
  img {
    margin-top: 10px;
  }
`;

const Content = styled.div``;

const Comments = styled.div`
  display: flex;
  flex-direction: column;
`;

const Comment = styled.div`
  display: flex;
  flex-direction: column;
  padding-top: 10px;
  padding-left: 10px;
  padding-bottom: 10px;
  border-top: 2px solid rgba(0, 0, 0, 0.1);
`;
const CommentInfo = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
`;
const CommentContent = styled.div`
  padding-top: 10px;
  padding-left: 10px;
  align-self: flex-start;
`;

const UserImg = styled.div`
  display: flex;
  flex-direction: row;
  h5 {
    align-self: center;
    padding-left: 3px;
    font-size: 18px;
  }
  img {
    width: 23px;
    height: 23px;
    border-radius: 50%;
    padding-bottom: 2px;
    margin-right: 6px;
  }
`;

const AddComment = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  flex-wrap: wrap;
  margin-left: 15px;
  padding-bottom: 10px;
  input {
    border-radius: 24px;
    max-width: 300px;
    height: 30px;
    border-color: grey;
    font-size: 15px;
    padding-left: 15px;
  }
  button {
    margin-left: 10px;
    border-radius: 24px;
    font-size: 16px;
    color: white;
    background-color: #0a66c2;
    border: none;
    width: 60px;
    transition-duration: 167ms;
    &:hover {
      cursor: pointer;
      background-color: rgb(10, 102, 160);
    }
  }
  img {
    margin-right: 10px;
    width: 30px;
    height: 30px;
    border-radius: 50%;
  }
`;
export default UserPosts;
