import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { lastValueFrom } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class DataService {
  constructor(private http: HttpClient) {}

  async get_popularity(api: string, startYear: number, endYear: number) {
    let params = new HttpParams()
      .set('start_year', startYear)
      .set('end_year', endYear);
    return await lastValueFrom(
      this.http.get<any>(api + '/api/v0/GetPopularity', {
        headers: new HttpHeaders({}),
        params: params,
      })
    );
  }

  async get_explicit(api: string, startYear: number, endYear: number) {
    let params = new HttpParams()
      .set('start_year', startYear)
      .set('end_year', endYear);
    return await lastValueFrom(
      this.http.get<any>(api + '/api/v0/GetExplicit', {
        headers: new HttpHeaders({}),
        params: params,
      })
    );
  }
}
