let BASE_URL = "http://localhost:8080";
if (process.env.NODE_ENV === "production") {
  BASE_URL = "http://gupwd.com:8080";
  console.log("Running using production host!"); // eslint-disable-line no-console
}

const USER_ROUTE = "/users";
const REGISTER_ROUTE = "/register";
const PAGE_ROUTE = "/pages";
const PAGE_LIST_ROUTE = "/pages";
const LOGIN_ROUTE = "/login";
const UPLOAD_ROUTE = "/pages/new"
const USER_PAGE_ROUTE = "/users"
export {
  BASE_URL,
  USER_ROUTE,
  PAGE_ROUTE,
  REGISTER_ROUTE,
  PAGE_LIST_ROUTE,
  LOGIN_ROUTE,
  UPLOAD_ROUTE,
  USER_PAGE_ROUTE
};
