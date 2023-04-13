import http from 'k6/http';
import { check } from 'k6';
import { Rate } from 'k6/metrics';

export const errorRate = new Rate('errors');
export const options = {
  // k6 run --vus 10 --duration 30s script.js
  stages: [
    { duration: '30s', target: 100 }, // simulate ramp-up of traffic from 1 to 60 users over 5 minutes.
    { duration: '1m', target: 100 }, // stay at 60 users for 10 minutes
    { duration: '30s', target: 200 }, // ramp-up to 100 users over 3 minutes (peak hour starts)
    { duration: '1m', target: 200 }, // stay at 100 users for short amount of time (peak hour)
    { duration: '30s', target: 100 }, // ramp-down to 60 users over 3 minutes (peak hour ends)
    { duration: '1m', target: 100 }, // continue at 60 for additional 10 minutes
    { duration: '30s', target: 0 }, // ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
};
export default function () {
  const url = 'http://103.13.207.248/v1/merchants/14/users';
  
  const params = {
    headers: {
      'Content-Type': 'application/json',
      'X-API-KEY':'ilYhqFdiMaghE2V67zO1',
      'API-KEY-DEV':'xiycYpedf2HfhwVn6EwI65m8O63oBoKXN43BUrOVw04=',
      'Authorization':'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzkzMDQwMzIsImlkIjoyLCJpc19tZXJjaGFudCI6dHJ1ZSwibWVyY2hhbnRfaWQiOjE0LCJyb2xlX2lkIjoxfQ.3u9OLDFOG-qyLk0drC3o8w29TSW7OoQCE07wnAoCYPY'
    },
  };
  check(http.get(url, params), {
    'status is 200': (r) => r.status == 200,
  }) || errorRate.add(1);
}