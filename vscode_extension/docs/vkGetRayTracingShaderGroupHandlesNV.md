# vkGetRayTracingShaderGroupHandlesKHR(3) Manual Page

## Name

vkGetRayTracingShaderGroupHandlesKHR - Query ray tracing pipeline shader
group handles



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the opaque handles of shaders in the ray tracing pipeline,
call:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
VkResult vkGetRayTracingShaderGroupHandlesKHR(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    firstGroup,
    uint32_t                                    groupCount,
    size_t                                      dataSize,
    void*                                       pData);
```

or the equivalent command

``` c
// Provided by VK_NV_ray_tracing
VkResult vkGetRayTracingShaderGroupHandlesNV(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    firstGroup,
    uint32_t                                    groupCount,
    size_t                                      dataSize,
    void*                                       pData);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device containing the ray tracing pipeline.

- `pipeline` is the ray tracing pipeline object containing the shaders.

- `firstGroup` is the index of the first group to retrieve a handle for
  from the
  [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pGroups`
  or
  [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pGroups`
  array.

- `groupCount` is the number of shader handles to retrieve.

- `dataSize` is the size in bytes of the buffer pointed to by `pData`.

- `pData` is a pointer to an application-allocated buffer where the
  results will be written.

## <a href="#_description" class="anchor"></a>Description

On success, an array of `groupCount` shader handles will be written to
`pData`, with each element being of size
[VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleSize`.

If `pipeline` was created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineLibraryGroupHandles"
target="_blank" rel="noopener">pipelineLibraryGroupHandles</a> feature
is enabled applications **can** query group handles from that pipeline,
even if the pipeline is a library and is never bound to a command
buffer. These group handles remain bitwise identical for any `pipeline`
which references the pipeline library. Group indices are assigned as-if
the pipeline was created without `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`.

Valid Usage

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-04619"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-04619"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-04619  
  `pipeline` **must** be a ray tracing pipeline

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-04050"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-04050"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-04050  
  `firstGroup` **must** be less than the number of shader groups in
  `pipeline`

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-02419"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-02419"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-02419  
  The sum of `firstGroup` and `groupCount` **must** be less than or
  equal to the number of shader groups in `pipeline`

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-02420"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-02420"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-02420  
  `dataSize` **must** be at least
  [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleSize`
  × `groupCount`

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-07828"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-07828"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-07828  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineLibraryGroupHandles"
  target="_blank" rel="noopener">pipelineLibraryGroupHandles</a> feature
  is not enabled, `pipeline` **must** not have been created with
  `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

Valid Usage (Implicit)

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-device-parameter"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-device-parameter"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parameter"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parameter"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html) handle

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-pData-parameter"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-pData-parameter"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes

- <a
  href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-arraylength"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-arraylength"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`

- <a href="#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parent"
  id="VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parent"></a>
  VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from
  `device`

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetRayTracingShaderGroupHandlesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
