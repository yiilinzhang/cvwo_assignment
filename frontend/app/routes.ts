import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("pages/HomeScreen.tsx"),
  route("posts/:id", "pages/TopicPostsScreen.$id.tsx"),
  route("postcomments/:id", "pages/PostCommentsScreen.$id.tsx"),
  route("addposts", "pages/AddPostScreen.tsx"),
  route("addtopics", "pages/AddTopicScreen.tsx"),
  route("sign-in", "pages/LoginScreen.tsx"),
  route("sign-up", "pages/SignupScreen.tsx"),
  route("editpost/:id", "pages/EditPostScreen.tsx"),
] satisfies RouteConfig;
