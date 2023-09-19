import { Component } from '@angular/core';
import { DataService } from 'src/app/services/data.service';

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent {
  constructor(private service: DataService) { }

  result: string = 'Awaiting button press.';

  get(api: string) {
    this.service.get_test(api).then(res => {
      if (res.fact) {
        this.result = res.fact;
      }
      else {
        this.result = "Get request failed, read console for more information.";
        console.log(res);
      }
    }).catch(res => {
      this.result = "Get request failed, read console for more information.";
      console.log(res);
    });
  }

  post(api: string) {
    this.service.post_test(api).then(res => {
      if (res.message) {
        this.result = res.message;
      }
      else {
        this.result = "Post request failed, read console for more information.";
        console.log(res);
      }
    }).catch(res => {
      this.result = "Post request failed, read console for more information.";
      console.log(res);
    });
  }
}
