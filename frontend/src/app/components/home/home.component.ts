import { Component } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { TabsComponent } from 'src/app/tabs/tabs.component';
import { EChartsOption } from 'echarts';

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})
export class HomeComponent {
  constructor(private service: DataService) {}

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
  // Charts
  chartOption: EChartsOption = {
    xAxis: {
      type: 'category',
      data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
    },
    yAxis: {
      type: 'value',
    },
    series: [
      {
        data: [820, 932, 901, 934, 1290, 1330, 1320],
        type: 'line',
      },
    ],
  };
}
