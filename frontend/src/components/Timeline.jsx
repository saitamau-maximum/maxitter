import { Box, Button, Card, CardContent, Typography } from "@mui/material";

export const Timeline = ({ posts, isLoading, fetchPosts }) => {
  return (
    <>
      <Box textAlign="center" m={3}>
        <Button
          variant="contained"
          color="primary"
          disabled={isLoading}
          onClick={fetchPosts}
        >
          {isLoading ? "読み込み中" : "更新する"}
        </Button>
      </Box>
      {posts.map((post) => (
        <Card key={post.id} sx={{ my: 2 }}>
          <CardContent>
            <Typography variant="body1">{post.body} {post.created_at}</Typography>
          </CardContent>
        </Card>
      ))}
    </>
  );
};
