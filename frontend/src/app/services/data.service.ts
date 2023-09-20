import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http'
import { lastValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  constructor(private http: HttpClient) { }

  private api = 'to-be-determined';

  async get_test(api: string) {
    return await lastValueFrom(this.http.get<any>(api, { headers: new HttpHeaders({}) }));
  }

  async post_test(api: string) {
    return await lastValueFrom(this.http.post<any>(api, { headers: new HttpHeaders({}), message: 14 }));
  }
}
