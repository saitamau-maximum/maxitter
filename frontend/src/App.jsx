import { Container, CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { CountPost } from "./components/CountPost";
import { Header } from "./components/Header"
import { useEffect, useState } from "react";
import { ColorModeProvider } from "./components/theme/ColorModeProvider.jsx";
import { ToggleTheme } from "./components/theme/ToggleTheme.jsx";

function App() {
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
    fetchAmountPost();
  }, []);

  return (
    <>
      <ColorModeProvider>
        <CssBaseline />
        <Header/>
        <ToggleTheme />
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
          <Timeline
            posts={posts}
            isLoading={isLoading}
            fetchPosts={fetchPosts}
          />
        </Container>
        <CountPost 
          amount={amount}
        />
      </ColorModeProvider>
    </>
  );
}

export default App;
