import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { Chart } from 'chart.js/auto';
import { map } from 'rxjs';
interface Regions {
  value: string;
  viewValue: string;
}

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent {
  constructor(private service: DataService) {}

  popularity: { chart: any; active: boolean } = {
    chart: Chart,
    active: false,
  };

  api: string = localStorage.getItem('api') as string;

  startYear = 1900;
  endYear = 2021;

  selectedValue1: string = 'africa';
  selectedValue2: string = 'asia';

  regions: Regions[] = [
    { value: 'africa', viewValue: 'Africa' },
    { value: 'americas', viewValue: 'The Americas' },
    { value: 'europe', viewValue: 'Europe' },
    { value: 'oceania', viewValue: 'Oceania' },
    { value: 'asia', viewValue: 'Asia' },
  ];

  saveAPI() {
    localStorage.setItem('api', this.api);
  }

  submit() {
    this.service
      .get_popularity(
        localStorage.getItem('api') as string,
        this.startYear,
        this.endYear
      )
      .then((res) => {
        if (res.years && res.popularities) {
          if (!this.popularity.active) {
            this.popularity.active = true;

            this.popularity.chart = new Chart('Popularity over Time', {
              type: 'bar',
              data: {
                labels: [],
                datasets: [],
              },
              options: {
                aspectRatio: 2.5,
              },
            });
          }

          this.popularity.chart.data = {
            labels: res.years,
            datasets: [
              {
                label: 'Popularity',
                data: res.popularities,
              },
            ],
          };

          this.popularity.chart.update();
        } else {
          ('Get request failed, read console for more information.');
          console.log(res);
        }
      })
      .catch((res) => {
        console.log(res);
      });
  }
}
