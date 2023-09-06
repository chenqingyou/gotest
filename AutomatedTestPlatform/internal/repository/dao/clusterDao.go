package dao

//func (intDao *StructDao) DBSClusterPerInsert(ctx context.Context, EngineResult domain.ClusterEnginePerResult) string {
//	err := intDao.db.WithContext(ctx).Create(&EngineResult).Error
//	if err != nil {
//		return fmt.Sprintf("DBInsert Mysql err[%v]", err)
//	}
//	return ""
//}
//
//func (intDao *StructDao) DBClusterPerDetailInsert(ctx context.Context, EngineResult domain.ClusterEnginePerDetailResult) string {
//	err := intDao.db.WithContext(ctx).Create(&EngineResult).Error
//	if err != nil {
//		return fmt.Sprintf("DBInsert Mysql err[%v]", err)
//	}
//	return ""
//}
//
//func (intDao *StructDao) DBClusterApiInsert(ctx context.Context, EngineResult domain.SearchEngineApiResult) string {
//	err := intDao.db.WithContext(ctx).Create(&EngineResult).Error
//	if err != nil {
//		return fmt.Sprintf("DBInsert Mysql err[%v]", err)
//	}
//	return ""
//}
//
//func (intDao *StructDao) DBClusterAlgoInsert(ctx context.Context, EngineResult domain.SearchEngineResultReload) string {
//	// 存毫秒数
//	err := intDao.db.WithContext(ctx).Create(&EngineResult).Error
//	if err != nil {
//		return fmt.Sprintf("DBInsert Mysql err[%v]", err)
//	}
//	return ""
//}
