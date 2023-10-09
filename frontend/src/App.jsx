
import { Container, CssBaseline } from "@mui/material";
import { GlobalStyles } from "@mui/material";
import { Form } from "./components/Form";
import { Timeline } from "./components/Timeline";
import { Pagination } from '@mui/material';
import { useEffect, useState } from "react";
import { ColorModeProvider } from "./components/theme/ColorModeProvider.jsx";
import { ToggleTheme } from "./components/theme/ToggleTheme.jsx";

function App() {

  const [posts, setPosts] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const onSubmitted = (post) => {
    setPosts([post, ...posts]);
  };

  
  const fetchPosts = async (page) => {
    setIsLoading(true);
    const index = page -1;
    const res = await fetch(`api/posts?page=${index}`);
    const data = await res.json();
    if (!res.ok) {
      console.error(data);
      return;
    }
    setPosts(data);
    setIsLoading(false);
  };
  
  
  // ページが変更されたときに呼び出される関数
  const handlePageChange = (event, page) => {
    
    setCurrentPage(page);
  };

  useEffect(() => {
    fetchPosts(currentPage);
  }, [currentPage]);


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
          <Pagination 
            count={10}//ページ数 
            color="primary"//色
            page={currentPage} // 現在のページ数
            onChange={handlePageChange} // ページが変更されたときに呼び出される関数
          />
        </Container>
      </ColorModeProvider>
    </>
    
  );
}

export default App;
