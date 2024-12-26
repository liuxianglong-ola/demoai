// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"context"
	"demogogo/internal/dao/internal"
	"demogogo/internal/model/entity"
)

// internalChannelsDao is internal type for wrapping internal DAO implements.
type internalChannelsDao = *internal.ChannelsDao

// channelsDao is the data access object for table channels.
// You can define custom methods on it to extend its functionality as you wish.
type channelsDao struct {
	internalChannelsDao
}

var (
	// Channels is globally public accessible object for table channels operations.
	Channels = channelsDao{
		internal.NewChannelsDao(),
	}
)

// Fill with you ideas below.

func (m *channelsDao) GetRecordById(ctx context.Context, id int64) (e *entity.Channels, err error) {
	err = m.Ctx(ctx).Where("id=?", id).Scan(&e)
	if err != nil {
		return nil, err
	}
	return e, nil
}