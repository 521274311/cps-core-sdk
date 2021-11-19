package r

type UcrowdRankItems struct {
	Commirate uint32 `json:"commirate,omitempty"`
	Price     string `json:"price,omitempty"`
	ItemId    string `json:"item_id,omitempty"`
}

type TbkDgMaterialOptionalRequest struct {
	IsOverseas        bool              `json:"is_overseas,omitempty"`
	IsTmall           bool              `json:"is_tmall,omitempty"`
	NeedFreeShipment  bool              `json:"need_free_shipment,omitempty"`
	NeedPrepay        bool              `json:"need_prepay,omitempty"`
	IncludePayRate30  bool              `json:"include_pay_rate_30,omitempty"`
	IncludeGoodsRate  bool              `json:"include_goods_rate,omitempty"`
	IncludeRfdRate    bool              `json:"include_rfd_rate,omitempty"`
	HasCoupon         bool              `json:"has_coupon,omitempty"`
	GetTopnRate       uint8             `json:"get_topn_rate,omitempty"`
	NpxLevel          uint8             `json:"npx_level,omitempty"`
	PageSize          uint8             `json:"page_size,omitempty"`
	Platform          uint8             `json:"platform,omitempty"`
	StartDsr          uint16            `json:"start_dsr,omitempty"`
	EndTkRate         uint16            `json:"end_tk_rate,omitempty"`
	StartTkRate       uint16            `json:"start_tk_rate,omitempty"`
	EndKaTkRate       uint16            `json:"end_ka_tk_rate,omitempty"`
	StartKaTkRate     uint16            `json:"start_ka_tk_rate,omitempty"`
	PageNo            uint32            `json:"page_no,omitempty"`
	EndPrice          uint32            `json:"end_price,omitempty"`
	StartPrice        uint32            `json:"start_price,omitempty"`
	MaterialId        uint32            `json:"material_id,omitempty"`
	AdzoneId          uint64            `json:"adzone_id,omitempty" require:"true"`
	LockRateEndTime   uint64            `json:"lock_rate_end_time,omitempty"`
	LockRateStartTime uint64            `json:"lock_rate_start_time,omitempty"`
	Sort              string            `json:"sort,omitempty"`
	Itemloc           string            `json:"itemloc,omitempty"`
	Cat               string            `json:"cat,omitempty"`
	Q                 string            `json:"q,omitempty"`
	Ip                string            `json:"ip,omitempty"`
	DeviceEncrypt     string            `json:"device_encrypt,omitempty"`
	DeviceValue       string            `json:"device_value,omitempty"`
	DeviceType        string            `json:"device_type,omitempty"`
	Longitude         string            `json:"longitude,omitempty"`
	Latitude          string            `json:"latitude,omitempty"`
	CityCode          string            `json:"city_code,omitempty"`
	SellerIds         string            `json:"seller_ids,omitempty"`
	SpecialId         string            `json:"special_id,omitempty"`
	RelationId        string            `json:"relation_id,omitempty"`
	PageResultKey     string            `json:"page_result_key,omitempty"`
	UcrowdId          uint32            `json:"ucrowd_id,omitempty"`
	UcrowdRankItems   []UcrowdRankItems `json:"ucrowd_rank_items,omitempty"`
}

type TbkDgMaterialOptionalResponse struct {
	TbkDgMaterialOptionalResponse struct {
		TotalResults uint64 `json:"total_results"`
		ResultList   struct {
			MapData []struct {
				UserType               uint8  `json:"user_type"`
				RewardInfo             uint8  `json:"reward_info"`
				ShopDsr                uint16 `json:"shop_dsr"`
				JddNum                 uint16 `json:"jdd_num"`
				Volume                 uint32 `json:"volume"`
				CouponTotalCount       uint32 `json:"coupon_total_count"`
				CouponRemainCount      uint32 `json:"coupon_remain_count"`
				UvSumPreSale           uint32 `json:"uv_sum_pre_sale"`
				TotalStock             uint32 `json:"total_stock"`
				SellNum                uint32 `json:"sell_num"`
				Stock                  uint32 `json:"stock"`
				SellerId               uint64 `json:"seller_id"`
				LevelOneCategoryId     uint64 `json:"level_one_category_id"`
				CategoryId             uint64 `json:"category_id"`
				ItemId                 uint64 `json:"item_id"`
				LockRateEndTime        uint64 `json:"lock_rate_end_time"`
				LockRateStartTime      uint64 `json:"lock_rate_start_time"`
				PresaleTailEndTime     uint64 `json:"presale_tail_end_time"`
				PresaleTailStartTime   uint64 `json:"presale_tail_start_time"`
				PresaleEndTime         uint64 `json:"presale_end_time"`
				PresaleStartTime       uint64 `json:"presale_start_time"`
				NumIid                 uint64 `json:"num_iid"`
				CouponStartTime        string `json:"coupon_start_time"`
				CouponEndTime          string `json:"coupon_end_time"`
				InfoDxjh               string `json:"info_dxjh"`
				TkTotalSales           string `json:"tk_total_sales"`
				TkTotalCommi           string `json:"tk_total_commi"`
				CouponId               string `json:"coupon_id"`
				Title                  string `json:"title"`
				PictUrl                string `json:"pict_url"`
				ReservePrice           string `json:"reserve_price"`
				ZkFinalPrice           string `json:"zk_final_price"`
				Provcity               string `json:"provcity"`
				ItemUrl                string `json:"item_url"`
				IncludeMkt             string `json:"include_mkt"`
				IncludeDxjh            string `json:"include_dxjh"`
				CommissionRate         string `json:"commission_rate"`
				CouponInfo             string `json:"coupon_info"`
				CommissionType         string `json:"commission_type"`
				ShopTitle              string `json:"shop_title"`
				CouponShareUrl         string `json:"coupon_share_url"`
				Url                    string `json:"url"`
				LevelOneCategoryName   string `json:"level_one_category_name"`
				CategoryName           string `json:"category_name"`
				ShortTitle             string `json:"short_title"`
				WhiteImage             string `json:"white_image"`
				Oetime                 string `json:"oetime"`
				Ostime                 string `json:"ostime"`
				JddPrice               string `json:"jdd_price"`
				XId                    string `json:"x_id"`
				CouponStartFee         string `json:"coupon_start_fee"`
				CouponAmount           string `json:"coupon_amount"`
				ItemDescription        string `json:"item_description"`
				Nick                   string `json:"nick"`
				OrigPrice              string `json:"orig_price"`
				TmallPlayActivityInfo  string `json:"tmall_play_activity_info"`
				RealPostFee            string `json:"real_post_fee"`
				LockRate               string `json:"lock_rate"`
				PresaleDiscountFeeText string `json:"presale_discount_fee_text"`
				PresaleDeposit         string `json:"presale_deposit"`
				YsylTljSendTime        string `json:"ysyl_tlj_send_time"`
				YsylCommissionRate     string `json:"ysyl_commission_rate"`
				YsylTljFace            string `json:"ysyl_tlj_face"`
				YsylClickUrl           string `json:"ysyl_click_url"`
				YsylTljUseEndTime      string `json:"ysyl_tlj_use_end_time"`
				YsylTljUseStartTime    string `json:"ysyl_tlj_use_start_time"`
				SaleBeginTime          string `json:"sale_begin_time"`
				SaleEndTime            string `json:"sale_end_time"`
				Distance               string `json:"distance"`
				UsableShopId           string `json:"usable_shop_id"`
				UsableShopName         string `json:"usable_shop_name"`
				SalePrice              string `json:"sale_price"`
				KuadianPromotionInfo   string `json:"kuadian_promotion_info"`
				SuperiorBrand          string `json:"superior_brand"`
				IsBrandFlashSale       string `json:"is_brand_flash_sale"`
				MatchScore             string `json:"match_score"`
				CommiScore             string `json:"commi_score"`
				TtSoldCount            string `json:"tt_sold_count"`
				HotFlag                string `json:"hot_flag"`
				SmallImages            struct {
					String []string `json:"string"`
				} `json:"small_images"`
				LocalizationExtend struct {
					OrderLeadTime       string `json:"order_lead_time"`
					UserRating          string `json:"user_rating"`
					DeliveryMinPrice    string `json:"delivery_min_price"`
					DeliveryFee         string `json:"delivery_fee"`
					OriginalDeliveryFee string `json:"original_delivery_fee"`
					DeliveryType        string `json:"delivery_type"`
					RecommendReasons    struct {
						String []string `json:"string"`
					} `json:"recommend_reasons"`
					SaleTags struct {
						String []string `json:"string"`
					} `json:"sale_tags"`
					FoodItemList struct {
						FoodMapData []struct {
							FoodPic            string `json:"food_pic"`
							FoodTitle          string `json:"food_title"`
							FoodPromotionPrice string `json:"food_promotion_price"`
							FoodReservePrice   string `json:"food_reserve_price"`
						} `json:"food_map_data"`
					} `json:"food_item_list"`
				} `json:"localization_extend"`
				TopnInfo struct {
					TopnQuantity   uint32 `json:"topn_quantity"`
					TopnTotalCount uint32 `json:"topn_total_count"`
					TopnEndTime    string `json:"topn_end_time"`
					TopnStartTime  string `json:"topn_start_time"`
					TopnRate       string `json:"topn_rate"`
				} `json:"topn_info"`
				BybtInfo struct {
					BybtBrandLogo    string `json:"bybt_brand_logo"`
					BybtPicUrl       string `json:"bybt_pic_url"`
					BybtCouponAmount string `json:"bybt_coupon_amount"`
					BybtShowPrice    string `json:"bybt_show_price"`
					BybtLowestPrice  string `json:"bybt_lowest_price"`
					BybtEndTime      string `json:"bybt_end_time"`
					BybtStartTime    string `json:"bybt_start_time"`
					BybtItemTags     struct {
						String []string `json:"string"`
					} `json:"bybt_item_tags"`
				} `json:"bybt_info"`
			} `json:"map_data"`
		} `json:"result_list"`
		PageResultKey string `json:"page_result_key"`
	} `json:"tbk_dg_material_optional_response,omitempty"`
	ErrorResponse struct {
		Code    int    `json:"code"`
		Msg     string `json:"msg"`
		SubMsg  string `json:"sub_msg"`
		SubCode string `json:"sub_code"`
	} `json:"error_response"`
}
