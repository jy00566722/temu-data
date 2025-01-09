const getOrderList = (body)=>{
    body = JSON.parse(body)
    if(body?.result?.list?.length > 0){
        let orderList = []
        for (let index = 0; index < body.result.list.length; index++) {
            const element = body.result.list[index];

        const data = {
            deliveryOrderSn :element.deliveryOrderSn, //发货单号
            subPurchaseOrderSn :element.subPurchaseOrderSn, //备货单号
            deliveryOrderSn :element.deliveryOrderSn, //发货批次号
            expressCompany :element.expressCompany, //物流公司
            expressDeliverySn :element.expressDeliverySn, //物流单号
            subWarehouseName :element.subWarehouseName, //收货仓名称
            supplierId :element.supplierId, //店铺ID
            productSkcId :element.productSkcId, //商品SKC
            productName :element.subPurchaseOrderBasicVO.productName, //商品名称
            productSkcPicture :element.subPurchaseOrderBasicVO.productSkcPicture, //商品图片
            skcExtCode :element.skcExtCode, //商品货号
            deliverPackageNum :element.deliverPackageNum, //发货包裹数量
            deliverSkcNum :element.deliverSkcNum, //发货商品总数量
            packageDetailList :element.packageDetailList, //发货SKU明细
            status :element.status, //发货单状态
    
        }
        orderList.push(data)
    }
    console.log(orderList)
}else{
    console.log('无数据')
}
}
export {getOrderList}