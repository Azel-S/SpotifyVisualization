import { Component } from '@angular/core';
import { DataService } from 'src/app/services/data.service';
import { Chart } from 'chart.js/auto';
import { MatTabChangeEvent } from '@angular/material/tabs';
import { FormControl } from '@angular/forms';
import { Observable } from 'rxjs';
import { map, startWith } from 'rxjs/operators';
import { MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';

@Component({
  selector: 'app-root',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css'],
})

export class HomeComponent {
  constructor(public service: DataService) {
    this.service.update_years(this.years);
    this.service.update_regions(this.regions);
    this.service.update_subregions(this.subregions);
    this.service.update_genres(this.genre_1.list)
    this.service.update_genres(this.genre_2.list)
    this.genre_1.filter = this.genre_1.control.valueChanges.pipe(startWith(''), map(value => this.filterGenre1(value!)));
    this.genre_2.filter = this.genre_2.control.valueChanges.pipe(startWith(''), map(value => this.filterGenre2(value!)));
  }

  // VARIABLES
  years: { min: number, max: number, selected_min: number, selected_max: number } = { min: 0, max: 0, selected_min: 0, selected_max: 0 }
  attributes: { list: string[], selected_1: string, selected_2: string } = { list: ["acousticness", "danceability", "energy", "speechiness", "loudness", "instrumentalness", "liveness", "tempo", "duration", "valence"], selected_1: "", selected_2: "" };
  regions: { list: string[], selected_1: string, selected_2: string } = { list: [], selected_1: "", selected_2: "" };
  subregions: { list: string[], selected: string } = { list: [], selected: "" };
  genre_1: { list: string[], selected: string, control: FormControl, filter: Observable<string[]> } = { list: [], selected: "pop", control: new FormControl(''), filter: new Observable<string[]>() }
  genre_2: { list: string[], selected: string, control: FormControl, filter: Observable<string[]> } = { list: [], selected: "rock", control: new FormControl(''), filter: new Observable<string[]>() }
  charts: { popularity: Boolean, explicit: Boolean, attribute: Boolean, genre: Boolean, title: Boolean } = { popularity: false, explicit: false, attribute: false, genre: false, title: false };

  tabIndex: number = 0;

  // FUNCTIONS
  updateTabIndex(event: MatTabChangeEvent) {
    this.tabIndex = event.index;

    switch (this.tabIndex) {
      case 0: {
        if (!this.charts.popularity) {
          new Chart('popularity_chart', {
            type: 'bar', data: { labels: [], datasets: [] },
            options: { scales: { x: { title: { display: true, text: 'Time' } }, y: { title: { display: true, text: 'Ratio of Popularity to Attribute' } } } }
          });
          this.charts.popularity = true;
        }
        break;
      }
      case 1: {
        if (!this.charts.explicit) {
          new Chart('explicit_chart', {
            type: 'line', data: { labels: [], datasets: [] },
            options: { scales: { x: { title: { display: true, text: 'Time' } }, y: { title: { display: true, text: 'Number of Explicit Songs' } } } }
          });
          this.charts.explicit = true;
        }
        break;
      }
      case 2: {
        if (!this.charts.attribute) {
          new Chart('attribute_chart',
            {
              type: 'line', data: { labels: [], datasets: [] },
              options: { scales: { x: { title: { display: true, text: 'Time' } }, y: { title: { display: true, text: 'Average Value of Attribute' } } } }
            });
          this.charts.attribute = true;
        }
        break;
      }
      case 3: {
        if (!this.charts.genre) {
          new Chart('genre_chart',
            {
              type: 'line', data: { labels: [], datasets: [] },
              options: { scales: { x: { title: { display: true, text: 'Time' } }, y: { title: { display: true, text: 'Popularity of Genre' } } } }
            });
          this.charts.genre = true;
        }
        break;
      }
      case 4: {
        if (!this.charts.title) {
          new Chart('title_chart',
            {
              type: 'line',
              data: { labels: [], datasets: [] },
              options: { scales: { x: { title: { display: true, text: 'Time' } }, y: { title: { display: true, text: 'Length of Characters' } } } }
            });
          this.charts.title = true;
        }
        break;
      }
      default: {
        break;
      }
    }
  }

  filterGenre1(value: string): string[] {
    return this.genre_1.list.filter(option => option.toLowerCase().startsWith(value.toLowerCase()));
  }

  filterGenre2(value: string): string[] {
    return this.genre_2.list.filter(option => option.toLowerCase().startsWith(value.toLowerCase()));
  }

  updateGenre1(event: MatAutocompleteSelectedEvent) {
    this.genre_1.selected = event.option.value;
  }

  updateGenre2(event: MatAutocompleteSelectedEvent) {
    this.genre_2.selected = event.option.value;
  }

  submit() {
    if (this.tabIndex == 0) {
      this.service.update_popularity(
        this.years.selected_min,
        this.years.selected_max,
        this.attributes.selected_1,
        Chart.getChart('popularity_chart')!
      );
    } else if (this.tabIndex == 1) {
      this.service.update_explicit(this.years.selected_min, this.years.selected_max, this.subregions.selected, Chart.getChart('explicit_chart')!);
    } else if (this.tabIndex == 2) {
      this.service.update_attribute(
        this.years.selected_min,
        this.years.selected_max,
        this.attributes.selected_1,
        this.attributes.selected_2,
        this.genre_1.selected,
        Chart.getChart('attribute_chart')!
      )
    }
    else if (this.tabIndex == 3) {
      this.service.update_genre(this.years.selected_min, this.years.selected_max, this.genre_1.selected, this.genre_2.selected, Chart.getChart('genre_chart')!);
    } else if (this.tabIndex == 4) {
      this.service.update_title_length(this.years.selected_min, this.years.selected_max, this.regions.selected_1, this.regions.selected_2, Chart.getChart('title_chart')!);
    } else {
      this.service.notify('Unexpected tab, please investigate.');
    }
  }

  test() {
    this.service.notify('Current Tab: ' + this.tabIndex.toString());
  }
  count() {
    this.service.update_tuples();
  }

  reload() {
    window.location.reload()
  }
}
