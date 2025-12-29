async function request(endpoint, options = {}) {
  const response = await fetch(endpoint, {
    headers: {
      "Content-Type": "application/json",
      ...options.headers,
    },
    ...options,
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error || "Request failed");
  }

  return data;
}

export async function register(email, password, firstName, lastName) {
  return request("/register", {
    method: "POST",
    body: JSON.stringify({
      email,
      password,
      first_name: firstName,
      last_name: lastName,
    }),
  });
}

export async function login(email, password) {
  return request("/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
}

function getAuthHeader() {
  const token = typeof localStorage !== 'undefined'
    ? localStorage.getItem('auth_token')
    : null;
  return token ? { Authorization: `Bearer ${token}` } : {};
}

export async function getFriends() {
  return request("/api/friends", {
    method: "GET",
    headers: getAuthHeader(),
  });
}

export async function getIncomingRequests() {
  return request("/api/friends/incoming", {
    method: "GET",
    headers: getAuthHeader(),
  });
}

export async function getOutgoingRequests() {
  return request("/api/friends/outgoing", {
    method: "GET",
    headers: getAuthHeader(),
  });
}

export async function sendFriendRequest(email) {
  return request("/api/friends/create", {
    method: "POST",
    headers: getAuthHeader(),
    body: JSON.stringify({ email }),
  });
}

export async function acceptFriendRequest(friendId) {
  return request("/api/friends/accept", {
    method: "PUT",
    headers: getAuthHeader(),
    body: JSON.stringify({ friend_id: friendId }),
  });
}
