import { useState } from "react";
import { IndexPage } from "./generated";
import GoLogo from "../../public/go.png";
import ReactLogo from "../../public/react.png";
import Counter from "./components/Counter";
import "./Home.css";

function Home({ number }: IndexPage) {
  const [count, setCount] = useState(number);
  return (
    <div className="home">
      <div className="img-container">
        <img src={GoLogo} alt="Go logo" height={70} width={170} />
        <img src={ReactLogo} alt="React logo" height={80} width={90} />
      </div>
      <h1>Go + React</h1>
      <div className="counter-container">
        <Counter count={count} increment={() => setCount(count + 1)} />
      </div>
      <a href="https://github.com/natewong1313/go-react-ssr" target="_blank">
        View project on GitHub
      </a>
    </div>
  );
}

export default Home;
