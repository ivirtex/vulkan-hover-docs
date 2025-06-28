# vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(3) Manual Page

## Name

vkGetRayTracingCaptureReplayShaderGroupHandlesKHR - Query opaque capture replay data for pipeline shader group handles



## [](#_c_specification)C Specification

To query the opaque capture data of shader groups in a ray tracing pipeline, call:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
VkResult vkGetRayTracingCaptureReplayShaderGroupHandlesKHR(
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
- `firstGroup` is the index of the first group to retrieve a handle for from the [VkRayTracingPipelineCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingPipelineCreateInfoKHR.html)::`pGroups` array.
- `groupCount` is the number of shader handles to retrieve.
- `dataSize` is the size in bytes of the buffer pointed to by `pData`.
- `pData` is a pointer to an application-allocated buffer where the results will be written.

## [](#_description)Description

On success, an array of `groupCount` shader handles will be written to `pData`, with each element being of size [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleCaptureReplaySize`.

Once queried, this opaque data **can** be provided at pipeline creation time (in a subsequent execution), using [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`pShaderGroupCaptureReplayHandle`, as described in [Ray Tracing Capture Replay](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#ray-tracing-capture-replay).

If `pipeline` was created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR` and the [pipelineLibraryGroupHandles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineLibraryGroupHandles) feature is enabled applications **can** query capture replay group handles from that pipeline. The capture replay handle remains bitwise identical for any `pipeline` which references the pipeline library. Group indices are assigned as-if the pipeline was created without `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`.

Valid Usage

- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-04620)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-04620  
  `pipeline` **must** be a ray tracing pipeline
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-04051)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-04051  
  `firstGroup` **must** be less than the number of shader groups in `pipeline`
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-03483)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-firstGroup-03483  
  The sum of `firstGroup` and `groupCount` **must** be less than or equal to the number of shader groups in `pipeline`
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-03484)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-03484  
  `dataSize` **must** be at least [VkPhysicalDeviceRayTracingPipelinePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPipelinePropertiesKHR.html)::`shaderGroupHandleCaptureReplaySize` × `groupCount`
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03606)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-rayTracingPipelineShaderGroupHandleCaptureReplay-03606  
  `VkPhysicalDeviceRayTracingPipelineFeaturesKHR`::`rayTracingPipelineShaderGroupHandleCaptureReplay` **must** be enabled to call this function
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-03607)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-03607  
  `pipeline` **must** have been created with a `flags` that included `VK_PIPELINE_CREATE_RAY_TRACING_SHADER_GROUP_HANDLE_CAPTURE_REPLAY_BIT_KHR`
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-07829)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-07829  
  If the [pipelineLibraryGroupHandles](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineLibraryGroupHandles) feature is not enabled, `pipeline` **must** not have been created with `VK_PIPELINE_CREATE_LIBRARY_BIT_KHR`

Valid Usage (Implicit)

- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-device-parameter)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parameter)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pData-parameter)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pData-parameter  
  `pData` **must** be a valid pointer to an array of `dataSize` bytes
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-arraylength)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-dataSize-arraylength  
  `dataSize` **must** be greater than `0`
- [](#VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parent)VUID-vkGetRayTracingCaptureReplayShaderGroupHandlesKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from `device`

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRayTracingCaptureReplayShaderGroupHandlesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0