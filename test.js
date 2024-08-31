import hello from 'k6/x/hello';

export default function () {
  console.log(`${hello.helloWorld()}`);
}
