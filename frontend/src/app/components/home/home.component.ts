import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { Chart } from 'chart.js/auto';

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})

export class HomeComponent {
  constructor(private service: DataService) { }

  public chart: any;
  result: string = 'Awaiting button press.';

  get(api: string) {
    this.service
      .get_test(api)
      .then((res) => {
        if (res.fact) {
          this.result = res.fact;
        } else {
          this.result =
            'Get request failed, read console for more information.';
          console.log(res);
        }
      })
      .catch((res) => {
        this.result = 'Get request failed, read console for more information.';
        console.log(res);
      });
  }

  post(api: string) {
    this.service
      .post_test(api)
      .then((res) => {
        if (res.message) {
          this.result = res.message;
        } else {
          this.result =
            'Post request failed, read console for more information.';
          console.log(res);
        }
      })
      .catch((res) => {
        this.result = 'Post request failed, read console for more information.';
        console.log(res);
      });
  }

  createChart() {
    this.chart = new Chart("Loudness over Time", {
      type: 'bar', //this denotes tha type of chart

      data: {// values on X-Axis
        labels: ['2022-05-10', '2022-05-11', '2022-05-12', '2022-05-13',
          '2022-05-14', '2022-05-15', '2022-05-16', '2022-05-17',],
        datasets: [
          {
            label: "Loudness",
            data: ['467', '576', '572', '79', '92',
              '574', '573', '576'],
            backgroundColor: 'blue'
          },
        ]
      },
      options: {
        aspectRatio: 2.5
      }
    });
  }
}
