import { Footer } from "@/components/footer";
import { Button } from "@/components/ui/button";
import { API_URL } from "@/lib/constants";
import { Link } from "react-router-dom";

export default function Login() {
  return (
    <>
      <div className="container mt-32">
        <div className="md:w-2/4 m-auto border-none">
          <div className="flex flex-col items-center space-y-1.5 p-6">
            <h3 className="font-semibold tracking-tight text-2xl lg:text-6xl">
              tls watch{" "}
              <span className="text-xs tracking-normal lg:text-xl">
                as a service
              </span>
            </h3>
            <p className="text-sm text-muted-foreground lg:text-xl">
              monitor your tls certificates without hassle
            </p>
          </div>
          <div className="flex p-6 pt-0">
            <Button className="m-auto lg:text-xl" asChild>
              <Link to={`${API_URL}/auth/login`}>login to start &#8594;</Link>
            </Button>
          </div>
        </div>
      </div>
      <Footer className="border-t fixed bottom-0 left-0 z-20 w-full p-1 md:flex md:items-center md:justify-between md:p-2" />
    </>
  );
}
