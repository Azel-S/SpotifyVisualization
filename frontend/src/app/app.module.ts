import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';

import { MatTabsModule } from '@angular/material/tabs';

// Components
// NOTE (Abbas): It seems GitHub pages does not supports routing, so avoid it if possible.
import { HomeComponent } from './components/home/home.component';
import { TabsComponent } from './tabs/tabs.component';

@NgModule({
  declarations: [HomeComponent, TabsComponent],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatTabsModule,
  ],
  providers: [],
  bootstrap: [HomeComponent],
})
export class AppModule {}
