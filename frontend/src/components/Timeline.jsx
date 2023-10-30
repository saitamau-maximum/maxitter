import { Box, Button, Card, CardContent, Typography , CardActionArea } from "@mui/material";

const formatDateTime = (dateTimeString) => {
  const options = {timeZone: "Asia/Tokyo", year: "numeric", month: "2-digit", day: "2-digit", hour: "2-digit", minute: "2-digit" };
  return new Date(dateTimeString).toLocaleString("ja-JP", options);
};

export const Timeline = ({ posts, isLoading, fetchPosts, LoadFailed}) => {
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
      <Typography textAlign="center"
        variant="h6"
        sx={{ color: 'error.main' }}
        >
          {LoadFailed ? "読み込みに失敗しました": ""}
      </Typography>
      {posts.map((post) => (
        <Card key={post.id} sx={{ my: 2 }}>
          <CardActionArea>
            <CardContent>
              <Typography variant="body1">{post.body}</Typography>
              <Typography variant="subtitle2" color="text.secondary">{formatDateTime(post.created_at)}</Typography>
            </CardContent>
          </CardActionArea>
        </Card>
      ))}
    </>
  );
};
