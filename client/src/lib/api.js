const API_BASE = "/api";

async function request(endpoint, options = {}) {
  const response = await fetch(`${API_BASE}${endpoint}`, {
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
