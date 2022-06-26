package model

// Account 账户，虚拟管理员和若干客户账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

// type RealEstate struct {
// 	RealEstateID string `json:"realEstateId"` //车辆ID
// 	Proprietor   string `json:"proprietor"`   //所有者(客户)(客户AccountId)
// 	Encumbrance  bool   `json:"encumbrance"`  //是否作为担保
// 	Carmodel     string `json:"carmodel"`
// 	Car_scz      string `json:"car_scz"`
// 	Car_scd      string `json:"car_scd"`
// 	Car_scsj     string `json:"car_scsj"`
// 	Car_lbjph    string `json:"car_lbjph"`
// 	Car_lbjscz   string `json:"Car_lbjscz"`
// 	Car_lbjscd   string `json:"car_lbjscd"`
// 	Car_lbjscsj  string `json:"car_lbjscsj"`
// }

// RealEstate 车辆作为担保出售、捐赠或质押时Encumbrance为true，默认状态false。
// 仅当Encumbrance为false时，才可发起出售、捐赠或质押
// Proprietor和RealEstateID一起作为复合键,保证可以通过Proprietor查询到名下所有的商品信息
type RealEstate struct {
	RealEstateID string `json:"realEstateId"` //车辆ID
	Proprietor   string `json:"proprietor"`   //所有者(客户)(客户AccountId)
	Encumbrance  bool   `json:"encumbrance"`  //是否作为担保
	Carmodel     string `json:"carmodel"`
	Car_scz      string `json:"car_scz"`
	Car_scd      string `json:"car_scd"`
	Car_scsj     string `json:"car_scsj"`
	Car_lbjph    string `json:"car_lbjph"`
	Car_lbjscz   string `json:"car_lbjscz"`
	Car_lbjscd   string `json:"car_lbjscd"`
	Car_lbjscsj  string `json:"car_lbjscsj"`
}

// Selling 销售要约
// 需要确定ObjectOfSale是否属于Seller
// 买家初始为空
// Seller和ObjectOfSale一起作为复合键,保证可以通过seller查询到名下所有发起的销售
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //销售对象(正在出售的车辆RealEstateID)
	Seller        string  `json:"seller"`        //发起销售人、卖家(卖家AccountId)
	Buyer         string  `json:"buyer"`         //参与销售人、买家(买家AccountId)
	Price         float64 `json:"price"`         //价格
	CreateTime    string  `json:"createTime"`    //创建时间
	SalePeriod    int     `json:"salePeriod"`    //智能合约的有效期(单位为天)
	SellingStatus string  `json:"sellingStatus"` //销售状态
}

// SellingStatusConstant 销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// SellingBuy 买家参与销售
// 销售对象不能是买家发起的
// Buyer和CreateTime作为复合键,保证可以通过buyer查询到名下所有参与的销售
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //参与销售人、买家(买家AccountId)
	CreateTime string  `json:"createTime"` //创建时间
	Selling    Selling `json:"selling"`    //销售对象
}

// Donating 捐赠要约
// 需要确定ObjectOfDonating是否属于Donor
// 需要指定受赠人Grantee，并等待受赠人同意接收
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //捐赠对象(正在捐赠的车辆RealEstateID)
	Donor            string `json:"donor"`            //捐赠人(捐赠人AccountId)
	Grantee          string `json:"grantee"`          //受赠人(受赠人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //捐赠状态
}

// DonatingStatusConstant 捐赠状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "捐赠中", //捐赠人发起捐赠合约，等待受赠人确认受赠
		"cancelled":     "已取消", //捐赠人在受赠人确认受赠之前取消捐赠或受赠人取消接收受赠
		"done":          "完成",  //受赠人确认接收，交易完成
	}
}

// DonatingGrantee 供受赠人查询的
type DonatingGrantee struct {
	Grantee    string   `json:"grantee"`    //受赠人(受赠人AccountId)
	CreateTime string   `json:"createTime"` //创建时间
	Donating   Donating `json:"donating"`   //捐赠对象
}

const (
	AccountKey         = "account-key"
	RealEstateKey      = "real-estate-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
	DonatingKey        = "donating-key"
	DonatingGranteeKey = "donating-grantee-key"
)
