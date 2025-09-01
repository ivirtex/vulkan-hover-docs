# VkPipelineCache(3) Manual Page

## Name

VkPipelineCache - Opaque handle to a pipeline cache object



## [](#_c_specification)C Specification

Pipeline cache objects allow the result of pipeline construction to be reused between pipelines and between runs of an application. Reuse between pipelines is achieved by passing the same pipeline cache object when creating multiple related pipelines. Reuse across runs of an application is achieved by retrieving pipeline cache contents in one run of an application, saving the contents, and using them to preinitialize a pipeline cache on a subsequent run. The contents of the pipeline cache objects are managed by the implementation. Applications **can** manage the host memory consumed by a pipeline cache object and control the amount of data retrieved from a pipeline cache object.

Pipeline cache objects are represented by `VkPipelineCache` handles:

```c++
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineCache)
```

## [](#_see_also)See Also

[VK\_DEFINE\_NON\_DISPATCHABLE\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_DEFINE_NON_DISPATCHABLE_HANDLE.html), [VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateComputePipelines.html), [vkCreateDataGraphPipelinesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateDataGraphPipelinesARM.html), [vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateExecutionGraphPipelinesAMDX.html), [vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateGraphicsPipelines.html), [vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreatePipelineCache.html), [vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesKHR.html), [vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRayTracingPipelinesNV.html), [vkDestroyPipelineCache](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyPipelineCache.html), [vkGetPipelineCacheData](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineCacheData.html), [vkMergePipelineCaches](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMergePipelineCaches.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineCache).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0