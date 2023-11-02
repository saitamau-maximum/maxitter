import { Container } from "@mui/material";
import { useEffect, useState } from "react";
import { Form } from "../components/Form";
import { Timeline } from "../components/Timeline";
import { CountPost } from "../components/CountPost";

export const Home = () => {
  const [posts, setPosts] = useState([]);
  const [postCount, setpostCount] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const onSubmitted = (post) => {
    setPosts([post, ...posts]);
    fetchAmountPost();
  };

  const fetchPosts = async () => {
    setIsLoading(true);
    const res = await fetch("/api/posts");
    const data = await res.json();
    if (!res.ok) {
      console.error(data);
      return;
    }
    setPosts(data.posts); // 投稿データをセット
    setpostCount(data.count); // 投稿件数をセット
    setIsLoading(false);
};

const fetchAmountPost = async () => {
  const res = await fetch("/api/posts/count");
  const data =await res.json();
  if (!res.ok) {
    console.error(data);
    return;
  }
  setpostCount(data.count);
};

  useEffect(() => {
    fetchPosts();
  }, []);

  return (
    <Container
      maxWidth="md"
      sx={{
        py: 3,
      }}
    >
      <Form onSubmitted={onSubmitted} />
      <Timeline posts={posts} isLoading={isLoading} fetchPosts={fetchPosts} />
      <CountPost 
          amount={postCount}
        />
    </Container>
  );
};
