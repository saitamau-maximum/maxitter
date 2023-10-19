import { Container, CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { CountPost } from "./components/CountPost";
import { useEffect, useState } from "react";
import { ColorModeProvider } from "./components/theme/ColorModeProvider.jsx";
import { ToggleTheme } from "./components/theme/ToggleTheme.jsx";

function App() {
  const [posts, setPosts] = useState([]);
  const [amount, setAmount] = useState(0);
  const [isLoading, setIsLoading] = useState(false);
  const onSubmitted = (post) => {
    setPosts([post, ...posts]);
    //setAmount(amount+1);
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
    setPosts(data);
    setIsLoading(false);
  };


  const fetchAmountPost = async () => {
    const res = await fetch("/api/posts/count");
    const data =await res.json();
    if (!res.ok) {
      console.error(data);
      return;
    }
    console.log(data);
    setAmount(data.count);
  };

  useEffect(() => {
    fetchPosts();
    fetchAmountPost();
  }, []);

  return (
    <>
      <ColorModeProvider>
        <CssBaseline />
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
