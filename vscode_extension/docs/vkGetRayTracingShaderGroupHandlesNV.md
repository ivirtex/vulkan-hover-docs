# vkGetRayTracingShaderGroupHandlesKHR(3) Manual Page

## Name

vkGetRayTracingShaderGroupHandlesKHR - Query ray tracing pipeline shader group handles



## [](#_c_specification)C Specification

To query the opaque handles of shaders in the ray tracing pipeline, call:

```c++
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

```c++
// Provided by VK_NV_ray_tracing
VkResult vkGetRayTracingShaderGroupHandlesNV(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    firstGroup,
    uint32_t                                    groupCount,
    size_t                                      dataSize,
    void*                                       pData);
```

## [](#_parameters)Parameters

- `device` is the logical device containing the ray tracing pipeline.
- `pipeline` is the ray tracing pipeline object containing the shaders.
- `firstGroup` is the index of the first group to retrieve a handle for from the [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pGroups` or [VkRayTracingPipelineCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoNV.html)::`pGroups` array.
- `groupCount` is the number of shader handles to retrieve.
- `dataSize` is the size in bytes of the buffer pointed to by `pData`.
- `pData` is a pointer to an application-allocated buffer where the results will be written.

## [](#_description)Description

On success, an array of `groupCount` shader handles will be written to `pData`, with each element being of size [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleSize`.

If `pipeline` was created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and the [pipelineLibraryGroupHandles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineLibraryGroupHandles) feature is enabled applications **can** query group handles from that pipeline, even if the pipeline is a library and is never bound to a command buffer. These group handles remain bitwise identical for any `pipeline` which references the pipeline library. Group indices are assigned as-if the pipeline was created without `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`.

Valid Usage

- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-04619)VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-04619  
  `pipeline` **must** be a ray tracing pipeline
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-04050)VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-04050  
  `firstGroup` **must** be less than the number of shader groups in `pipeline`
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-02419)VUID-vkGetRayTracingShaderGroupHandlesKHR-firstGroup-02419  
  The sum of `firstGroup` and `groupCount` **must** be less than or equal to the number of shader groups in `pipeline`
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-02420)VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-02420  
  `dataSize` **must** be at least [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleSize` × `groupCount`
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-07828)VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-07828  
  If the [pipelineLibraryGroupHandles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineLibraryGroupHandles) feature is not enabled, `pipeline` **must** not have been created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-device-parameter)VUID-vkGetRayTracingShaderGroupHandlesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parameter)VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-pData-parameter)VUID-vkGetRayTracingShaderGroupHandlesKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-arraylength)VUID-vkGetRayTracingShaderGroupHandlesKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parent)VUID-vkGetRayTracingShaderGroupHandlesKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRayTracingShaderGroupHandlesKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0