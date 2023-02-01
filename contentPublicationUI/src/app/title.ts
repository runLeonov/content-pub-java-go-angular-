import {User} from "./user";

export class Title {
  public id: number;
  public titleName: string;
  public originalAuthor: string;
  public creationDate: Date = new Date();
  public description: string = '';
  public titleImg: string = '';
  public content: TitleContent;
  public categories: Category[];
  public tags: Tag[];
  public serials: Serial[];


  constructor(id: number, titleName: string, originalAuthor: string, creationDate: Date, description: string, titleImgBase64: string, content: TitleContent, categories: Category[], tags: Tag[], serials: Serial[]) {
    this.id = id;
    this.titleName = titleName;
    this.originalAuthor = originalAuthor;
    this.creationDate = creationDate;
    this.description = description;
    this.titleImg = titleImgBase64;
    this.content = content;
    this.categories = categories;
    this.tags = tags;
    this.serials = serials;
  }

  static getEmptyTitle(): Title {
    return new Title(0, "", "", new Date(), "", "", TitleContent.getEmptyContent(), [], [], []);
  }
}

export class Like {
  public titleContentId: number
  public userId: number


  constructor(titleContentId: number, userId: number) {
    this.titleContentId = titleContentId;
    this.userId = userId;
  }

  static getEmptyLike(): Like {
    return new Like(0, 0);
  }
}

export class PossibleContent {
  public categories: Category[];
  public tags: Tag[];
  public serials: Serial[];

  constructor(categories: Category[], tags: Tag[], serials: Serial[]) {
    this.categories = categories;
    this.tags = tags;
    this.serials = serials;
  }
}

export class TitleContent {
  public id: number;
  public titleId: number;
  public likesCount: number = 0;
  public likes: Like[] = [];
  public views: number;
  public images: Image[];


  constructor(id: number, titleId: number, likes: Like[], views: number, images: Image[]) {
    this.id = id;
    this.titleId = titleId;
    this.likes = likes;
    this.views = views;
    this.images = images;
  }

  static getEmptyContent(): TitleContent {
    return new TitleContent(0, 0, [], 0, []);
  }
}

export class Image {
  public id: number;
  public img64: string;

  constructor(id: number, img64: string) {
    this.id = id;
    this.img64 = img64;
  }

  static getEmptyImg(): Image {
    return new Image(0, "");
  }
}


export class Category {
  public id: number;
  public genre: string;
  public titles: Title[];


  constructor(id: number, genre: string, titles: Title[]) {
    this.id = id;
    this.genre = genre;
    this.titles = titles;
  }

  static getEmptyCategory(): Category {
    return new Category(0, "", []);
  }
}

export class Tag {
  public id: number;
  public tag: string;
  public titles: Title[];


  constructor(id: number, tagName: string, titles: Title[]) {
    this.id = id;
    this.tag = tagName;
    this.titles = titles;
  }

  static getEmptyTag(): Tag {
    return new Tag(0, "", []);
  }
}

export class Serial {
  public id: number;
  public serialName: string;
  public titles: Title[];

  constructor(id: number, serialName: string, titles: Title[]) {
    this.id = id;
    this.serialName = serialName;
    this.titles = titles;
  }

  static getEmptySerial(): Serial {
    return new Serial(0, "", []);
  }
}
