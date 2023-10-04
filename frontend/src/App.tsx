import { Container } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { useEffect, useState } from "react";
import { Post } from "./types/user";
import React from "react";


function App() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const onSubmitted = (post: Post) => {
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
    <>
      <GlobalStyles
        styles={{
          body: {
            margin: 0,
          },
        }}
      />
      <Container
        maxWidth="md"
        sx={{
          py: 3,
        }}
      >
        <Form onSubmitted={onSubmitted} />
        <Timeline posts={posts} isLoading={isLoading} fetchPosts={fetchPosts} />
      </Container>
    </>
  );
}

export default App;
