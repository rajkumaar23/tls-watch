import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { API_URL } from "@/lib/constants";
import { Link } from "react-router-dom";

export default function Login() {
  return (
    <div className="container mt-32">
      <Card className="md:w-2/4 m-auto border-none">
        <CardHeader>
          <CardTitle className="text-2xl lg:text-6xl">
            tls watch <span className="text-xs lg:text-xl">as a service</span>
          </CardTitle>
          <CardDescription className="lg:text-xl">
            monitor your tls certificates without hassle
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Button className="m-auto lg:text-xl" asChild>
            <Link to={`${API_URL}/auth/login`}>login to start &#8594;</Link>
          </Button>
        </CardContent>
      </Card>
    </div>
  );
}
