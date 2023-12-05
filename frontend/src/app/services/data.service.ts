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
  notify(message: string, action: string = 'Close', duration: number = 2000) {
    this.snackBar.open(message, action, { duration: duration });
  }

  // API Link
  saveAPI(api: string) {
    if (api[api.length - 1] == '/') {
      api = api.substring(0, api.length - 1)
    }

    localStorage.setItem('api', api);
  }

  getAPI() {
    return 'api' in localStorage ? localStorage.getItem('api') : '';
  }

  // Backend Param Updaters
  async update_years(years: {
    min: number;
    max: number;
    selected_min: number;
    selected_max: number;
  }) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetYearRange', {
        headers: new HttpHeaders({}),
      })
    )
      .then((res) => {
        years.min = res.start_year;
        years.max = res.end_year;

        if (years.selected_min < years.min) {
          years.selected_min = years.min;
        }
        if (years.selected_max < years.max) {
          years.selected_max = years.max;
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }

  async update_regions(regions: {
    list: string[];
    selected_1: string;
    selected_2: string;
  }) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetRegions', {
        headers: new HttpHeaders({}),
      })
    )
      .then((res) => {
        regions.list.length = 0;

        (res.regions as string[]).forEach((region) => {
          regions.list.push(region);
        });

        if (regions.list.length > 0) {
          regions.selected_1 = regions.list[0];
          regions.selected_2 = regions.list[0];
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }

  async update_subregions(subregions: {
    list: string[];
    selected: string;
  }) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetSubregions', {
        headers: new HttpHeaders({}),
      })
    )
      .then((res) => {
        console.log(res)
        subregions.list.length = 0;

        (res.subregions as string[]).forEach((subregion) => {
          subregions.list.push(subregion);
        });

        if (subregions.list.length > 0) {
          subregions.selected = subregions.list[0];
        }
      })
      .catch((error) => {
        console.log(error);
      });
  }

  async update_genres(genres: string[]) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetGenres', {
        headers: new HttpHeaders({}),
      })
    )
      .then((res) => {
        genres.length = 0;

        (res.genres as string[]).forEach((genre) => {
          genres.push(genre);
        });
      })
      .catch((error) => {
        console.log(error);
      });
  }

  // Backend Graph Updaters
  async update_popularity(startYear: number, endYear: number, attribute: string, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetPopularity', {
        headers: new HttpHeaders({}),
        params: new HttpParams()
          .set('start_year', startYear)
          .set('end_year', endYear)
          .set('attribute', attribute)
      })
    )
      .then((res) => {
        if (res.years && res.popularities) {
          chart.data = {
            labels: res.years,
            datasets: [{ label: 'Popularity', data: res.popularities }],
          };
          chart.update();
        } else {
          this.notify('Request failed, read console.');
          console.log(res);
        }
      })
      .catch((res) => {
        this.notify('Request failed, read console.');
        console.log(res);
      });
  }

  async update_explicit(startYear: number, endYear: number, subregion: string, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetExplicit', {
        headers: new HttpHeaders({}),
        params: new HttpParams()
          .set('start_year', startYear)
          .set('end_year', endYear)
          .set('subregion', subregion)
      })
    )
      .then((res) => {
        if (res.years && res.explicit) {
          chart.data = {
            labels: res.years,
            datasets: [{ label: 'Explicit', data: res.explicit }],
          };
          chart.update();
        } else {
          this.notify('Request failed, read console.');
          console.log(res);
        }
      })
      .catch((res) => {
        this.notify('Request failed, read console.');
        console.log(res);
      });
  }

  async update_attribute(startYear: number, endYear: number, attribute1: string, attribute2: string, genre: string, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetAttributeComparison', {
        headers: new HttpHeaders({}),
        params: new HttpParams()
          .set('start_year', startYear)
          .set('end_year', endYear)
          .set('attribute_1', attribute1)
          .set('attribute_2', attribute2)
          .set('genre', genre)
      })
    )
      .then((res) => {
        console.log(res);
        if (res.years && res.attribute_1 && res.attribute_2) {
          chart.data = { labels: res.years, datasets: [{ label: attribute1, data: res.attribute_1 }, { label: attribute2, data: res.attribute_2 }] };
          chart.update()
        } else {
          this.notify('Request failed, read console.');
          console.log(res);
        }
      })
      .catch((res) => {
        this.notify('Request failed, read console.');
        console.log(res);
      });
  }

  async update_genre(startYear: number, endYear: number, genre_1: string, genre_2: string, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetGenrePopularity', {
        headers: new HttpHeaders({}),
        params: new HttpParams().set('start_year', startYear).set('end_year', endYear).set('genre_1', genre_1).set('genre_2', genre_2)
      })
    ).then((res) => {
      if (res.years && res.popularity_1 && res.popularity_2) {
        chart.data = { labels: res.years, datasets: [{ label: genre_1, data: res.popularity_1 }, { label: genre_2, data: res.popularity_2 }] };
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

  async update_title_length(startYear: number, endYear: number, region_1: string, region_2: string, chart: Chart) {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/GetTitleLength', {
        headers: new HttpHeaders({}),
        params: new HttpParams().set('start_year', startYear).set('end_year', endYear).set('region_1', region_1).set('region_2', region_2)
      })
    ).then((res) => {
      if (res.years && res.title_1 && res.title_2) {
        chart.data = { labels: res.years, datasets: [{ label: region_1, data: res.title_1 }, { label: region_2, data: res.title_2 }] };
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
  async update_tuples() {
    lastValueFrom(
      this.http.get<any>(this.getAPI() + '/api/v0/CountTuples', {
        headers: new HttpHeaders({}),
      })
    )
      .then((res) => {
        this.notify('Group 01 has ' + res + ' total tuples!', "Close", 0);
      })
      .catch((res) => {
        this.notify('Request failed, read console.');
      });
  }
}
