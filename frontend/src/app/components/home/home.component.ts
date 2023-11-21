import { Component } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { Chart } from 'chart.js/auto';
import { MatTabChangeEvent } from '@angular/material/tabs';
import { FormControl } from '@angular/forms';
import { Observable } from 'rxjs';
import { map, startWith } from 'rxjs/operators';

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})

export class HomeComponent {
  constructor(public service: DataService) {
    this.service.update_years(this.years);
    this.service.update_regions(this.regions);
    this.service.update_genres(this.genres.list)
    this.genres.filter = this.genres.control.valueChanges.pipe(startWith(''), map(value => this.filterGenre(value!)));
  }

  // VARIABLES
  years: { min: number, max: number, selected_min: number, selected_max: number } = { min: 0, max: 0, selected_min: 0, selected_max: 0 }
  regions: { list: string[], selected_1: string, selected_2: string } = { list: [], selected_1: "", selected_2: "" };
  genres: { list: string[], control: FormControl, filter: Observable<string[]> } = { list: [], control: new FormControl(''), filter: new Observable<string[]>() }
  charts: { popularity: Boolean, explicit: Boolean } = { popularity: false, explicit: false };

  tabIndex: number = 0;

  // FUNCTIONS
  updateTabIndex(event: MatTabChangeEvent) {
    this.tabIndex = event.index;
  }

  filterGenre(value: string): string[] {
    return this.genres.list.filter(option => option.toLowerCase().includes(value.toLowerCase()));
  }

  submit() {
    if (this.tabIndex == 0) {
      if (!this.charts.popularity) {
        new Chart('popularity_chart', { type: 'bar', data: { labels: [], datasets: [] } });
        this.charts.popularity = true;
      }

      this.service.update_popularity(this.years.selected_min, this.years.selected_max, Chart.getChart('popularity_chart')!);
    } else if (this.tabIndex == 1) {
      if (!this.charts.explicit) {
        new Chart('explicit_chart', { type: 'bar', data: { labels: [], datasets: [] } });
        this.charts.explicit = true;
      }

      this.service.update_explicit(this.years.selected_min, this.years.selected_max, Chart.getChart('explicit_chart')!);
    } else {
      this.service.notify("Unexpected tab, please investigate.");
    }
  }

  test() {
    this.service.notify("Current Tab: " + this.tabIndex.toString());
  }
}