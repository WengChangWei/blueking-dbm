# -*- coding: utf-8 -*-
"""
TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-DB管理系统(BlueKing-BK-DBM) available.
Copyright (C) 2017-2023 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at https://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""

from django.utils.translation import gettext_lazy as _

from ..base import BaseApi
from ..domains import KUBERNETES_APIGW_DOMAIN


class _KubernetesApi(BaseApi):
    MODULE = _("k8s 服务")
    BASE = KUBERNETES_APIGW_DOMAIN

    def __init__(self):
        self.create_cluster = self.generate_data_api(
            method="POST",
            url="/v4/dbs/cluster/create",
            description=_("创建集群"),
        )
        self.apply_clb = self.generate_data_api(
            method="POST",
            url="/v4/dbs/clb/create",
            description=_("创建集群clb"),
        )
        self.get_clb = self.generate_data_api(
            method="POST",
            url="/v4/dbs/clb/get",
            description=_("获取集群clb"),
        )
        self.expose_ports = self.generate_data_api(
            method="POST",
            url="/v4/dbs/oopsRequest/expose",
            description=_("暴露端口"),
        )
        self.get_regions = self.generate_data_api(
            method="GET",
            url="/v4/dbs/metadata/k8s_cluster_config/regions",
            description=_("获取区域列表"),
        )


KubernetesApi = _KubernetesApi()
