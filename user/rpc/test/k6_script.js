import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';

const client = new grpc.Client();
client.load(['definitions'], '../../user.proto');

export const options = {
    vus: 100,
    duration: "5s",
    maxDuration: "30s"
}

const data = {access_token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQ0MTY5NzEsIlVzZXJJZCI6ImU5OTFhNjYxLWU3NmUtNDE2OS04NTFiLWQyZTliMTZiNmNmZiJ9.t9Sc2khsIwWjEdgXF6jFN6NKdez8J_rW1wN7dExFyu8"}

export default function () {
  client.connect('localhost:8080', {
    plaintext: true,
    timeout: 1000
  })

//   console.log('In here')
  const response = client.invoke('user.User/verifyToken', data);

  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
  });

//   console.log(JSON.stringify(response.message));

  client.close();
//   sleep(1);
}