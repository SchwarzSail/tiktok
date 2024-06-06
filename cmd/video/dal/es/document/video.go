package document

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/video/dal/es"
	"tiktok/cmd/video/dal/es/model"
	"tiktok/kitex_gen/video"
)

func CreateVideo(ctx context.Context, model model.Video, id string) (err error) {
	_, err = es.EsClient.Index().Index(model.Index()).Id(id).BodyJson(&model).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "es.document.CreateVideo failed")
	}
	return nil
}

func UpdateVideo(ctx context.Context, id, key, value string) (err error) {
	_, err = es.EsClient.Update().Index(model.Video{}.Index()).Id(id).Doc(map[string]any{
		key: value,
	}).Do(ctx)
	if err != nil {
		return errors.Wrap(err, "es.document.UpdateVideo failed")
	}
	return nil
}

func QueryVideoByTime(ctx context.Context, timeStamp int64) (list []string, err error) {
	query := elastic.NewRangeQuery("created_at").Gte(timeStamp)
	res, err := es.EsClient.Search(model.Video{}.Index()).Query(query).Size(10).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "es,QueryByTime failed")
	}
	list = make([]string, 0)
	var v model.Video
	for _, data := range res.Hits.Hits {
		if err = json.Unmarshal(data.Source, &v); err != nil {
			return nil, errors.Wrap(err, "序列化失败")
		}
		list = append(list, strconv.Itoa(int(v.Vid)))
	}
	return list, nil
}

func FilterVideo(ctx context.Context, filter *video.SearchRequest) (list []string, err error) {
	query := elastic.NewBoolQuery()
	//匹配关键词,模糊匹配
	if filter.Keyword != "" {
		query = query.Should(elastic.NewMatchQuery("title", filter.Keyword))
		query = query.Should(elastic.NewMatchQuery("description", filter.Keyword))
	}
	//精确匹配
	if filter.Username != nil {
		query = query.Must(elastic.NewTermQuery("user_name", *filter.Username))
	}
	//时间匹配
	if filter.FromDate != nil || filter.ToDate != nil {
		if filter.FromDate != nil && filter.ToDate != nil {
			query = query.Must((elastic.NewRangeQuery("created_at").Gte(*filter.FromDate)).Lte(*filter.ToDate))
		} else if filter.FromDate != nil {
			query = query.Must(elastic.NewRangeQuery("created_at").Gte(*filter.FromDate))
		} else {
			query = query.Must(elastic.NewRangeQuery("created_at").Lte(*filter.ToDate))
		}
	}
	start := (filter.PageNum - 1) * filter.PageSize
	if start < 0 {
		start = 0
	}
	res, err := es.EsClient.Search(model.Video{}.Index()).Query(query).From(int(start)).Size(int(filter.PageSize)).Do(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "es.document.FilterVideo failed")
	}
	var v model.Video
	list = make([]string, 0)
	for _, data := range res.Hits.Hits {
		err = json.Unmarshal(data.Source, &v)
		if err != nil {
			return nil, errors.Wrap(err, "es.document.FilterVideo 序列化失败")
		}
		list = append(list, strconv.Itoa(int(v.Vid)))
	}
	return list, nil
}
