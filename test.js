import minecraft from 'k6/x/minecraft';

const mc = new minecraft.Client();

export default function () {
  mc.JoinServer("localhost:25565");
}
