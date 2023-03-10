import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})

export class LoginSignUpService {

  constructor(private client: HttpClient) { }

  AddOnLogin(username: string, password: string): Promise<any>
  {
    const acctInfo = { username: username, password: password };
    return this.client.post("/credentials/login", acctInfo).toPromise();
  }

  AddOnSignUp(username: string, password: string): Promise<any>
  {
    const acctInfo = { username: username, password: password };
    return this.client.post("/credentials/signup", acctInfo).toPromise();
  }
}
