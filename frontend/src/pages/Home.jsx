import { Container } from "@mui/material";
import { useEffect, useState } from "react";
import { Form } from "../components/Form";
import { Timeline } from "../components/Timeline";
import { Pagination } from "@mui/material";

export const Home = () => {
  const [posts, setPosts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const onSubmitted = (post) => {
    setPosts([post, ...posts]);
  };

  const fetchPosts = async () => {
    setIsLoading(true);
    const res = await fetch("/api/posts");
    const data = await res.json();
    if (!res.ok) {
      console.error(data);
      return;
    }
    setPosts(data);
    setIsLoading(false);
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
      <Pagination count={10} size="large" />
    </Container>
  );
};
