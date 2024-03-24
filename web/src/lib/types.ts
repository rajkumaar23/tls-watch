export type Domain = {
  id: number;
  user_id: number;
  domain: string;
  created_at: string;
  updated_at: string;
};

export type User = {
    id: number;
    oidc_subject: string;
    name: string;
    picture: string;
    created_at: string;
    updated_at: string;
  };