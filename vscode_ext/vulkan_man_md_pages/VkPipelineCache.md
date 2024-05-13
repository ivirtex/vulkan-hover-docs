# VkPipelineCache(3) Manual Page

## Name

VkPipelineCache - Opaque handle to a pipeline cache object



## <a href="#_c_specification" class="anchor"></a>C Specification

Pipeline cache objects allow the result of pipeline construction to be
reused between pipelines and between runs of an application. Reuse
between pipelines is achieved by passing the same pipeline cache object
when creating multiple related pipelines. Reuse across runs of an
application is achieved by retrieving pipeline cache contents in one run
of an application, saving the contents, and using them to preinitialize
a pipeline cache on a subsequent run. The contents of the pipeline cache
objects are managed by the implementation. Applications **can** manage
the host memory consumed by a pipeline cache object and control the
amount of data retrieved from a pipeline cache object.

Pipeline cache objects are represented by `VkPipelineCache` handles:

``` c
// Provided by VK_VERSION_1_0
VK_DEFINE_NON_DISPATCHABLE_HANDLE(VkPipelineCache)
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html),
[vkCreateExecutionGraphPipelinesAMDX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateExecutionGraphPipelinesAMDX.html),
[vkCreateGraphicsPipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateGraphicsPipelines.html),
[vkCreatePipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreatePipelineCache.html),
[vkCreateRayTracingPipelinesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesKHR.html),
[vkCreateRayTracingPipelinesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRayTracingPipelinesNV.html),
[vkDestroyPipelineCache](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyPipelineCache.html),
[vkGetPipelineCacheData](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineCacheData.html),
[vkMergePipelineCaches](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMergePipelineCaches.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineCache"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
