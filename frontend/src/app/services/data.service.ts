import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { MatSnackBar, MatSnackBarConfig } from '@angular/material/snack-bar';
import { lastValueFrom } from 'rxjs';
import { Chart } from 'chart.js';

@Injectable({
  providedIn: 'root',
})

export class DataService {
  constructor(private http: HttpClient, private snackBar: MatSnackBar) { }

  // Notify
  notify(message: string, action: string = "Close", duration: number = 2000) {
    this.snackBar.open(message, action, { duration: duration });
  }

  // API Link
  saveAPI(api: string) {
    localStorage.setItem('api', api);
  }

  getAPI() {
    return "api" in localStorage ? localStorage.getItem('api') : '';
  }

  // Backend Param Updaters
  async update_years(years: { min: number, max: number, selected_min: number, selected_max: number }) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetYearRange', {
        headers: new HttpHeaders({})
      })
    ).then((res) => {
      years.min = res.start_year;
      years.max = res.end_year;

      if (years.selected_min < years.min) {
        years.selected_min = years.min;
      }
      if (years.selected_max < years.max) {
        years.selected_max = years.max;
      }
    }).catch((error) => {
      console.log(error);
    });
  }

  async update_regions(regions: { list: string[], selected_1: string, selected_2: string }) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetRegions', {
        headers: new HttpHeaders({})
      })
    ).then((res) => {
      regions.list.length = 0;

      (res.regions as string[]).forEach((region) => {
        regions.list.push(region);
      })

      if (regions.list.length > 0) {
        regions.selected_1 = regions.list[0]
        regions.selected_2 = regions.list[0]
      }
    }).catch((error) => {
      console.log(error);
    });
  }

  async update_genres(genres: string[]) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetGenres', {
        headers: new HttpHeaders({})
      })
    ).then((res) => {
      genres.length = 0;

      (res.genres as string[]).forEach((genre) => {
        genres.push(genre);
      })
    }).catch((error) => {
      console.log(error);
    });
  }

  // Backend Graph Updaters
  async update_popularity(startYear: number, endYear: number, chart: Chart) {
    lastValueFrom(this.http.get<any>(this.getAPI() + '/api/v0/GetPopularity', {
      headers: new HttpHeaders({}),
      params: new HttpParams().set('start_year', startYear).set('end_year', endYear),
    })).then((res) => {
      if (res.years && res.popularities) {
        chart.data = { labels: res.years, datasets: [{ label: 'Popularity', data: res.popularities }] };
        chart.update()
      } else {
        this.notify('Request failed, read console.');
        console.log(res);
      }
    }).catch((res) => {
      this.notify('Request failed, read console.');
      console.log(res);
    });
  }

  async update_explicit(startYear: number, endYear: number, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetExplicit', {
        headers: new HttpHeaders({}),
        params: new HttpParams().set('start_year', startYear).set('end_year', endYear)
      })
    ).then((res) => {
      if (res.years && res.explicit) {
        chart.data = { labels: res.years, datasets: [{ label: 'Explicit', data: res.explicit }] };
        chart.update()
      } else {
        this.notify('Request failed, read console.');
        console.log(res);
      }
    })
      .catch((res) => {
        this.notify('Request failed, read console.');
        console.log(res);
      });;
  }
  async update_genre(startYear: number, endYear: number, genre_1: string, genre_2: string,  chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetGenreFollowers', {
        headers: new HttpHeaders({}),
        params: new HttpParams().set('start_year', startYear).set('end_year', endYear).set('genre_1', genre_1).set('genre_2', genre_2)
      })
    ).then((res) => {
      if (res.years && res.followers_1 && res.followers_2) {
        chart.data = { labels: res.years, datasets: [{ label: genre_1, data: res.followers_1 }, { label:  genre_2, data: res.followers_2 }] };
        chart.update()
      } else {
        this.notify('Request failed, read console.');
        console.log(res);
      }
    })
      .catch((res) => {
        this.notify('Request failed, read console.');
        console.log(res);
      });;
  }
}
