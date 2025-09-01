# vkGetRayTracingShaderGroupStackSizeKHR(3) Manual Page

## Name

vkGetRayTracingShaderGroupStackSizeKHR - Query ray tracing pipeline shader group shader stack size



## [](#_c_specification)C Specification

To query the pipeline stack size of shaders in a shader group in the ray tracing pipeline, call:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
VkDeviceSize vkGetRayTracingShaderGroupStackSizeKHR(
    VkDevice                                    device,
    VkPipeline                                  pipeline,
    uint32_t                                    group,
    VkShaderGroupShaderKHR                      groupShader);
```

## [](#_parameters)Parameters

- `device` is the logical device containing the ray tracing pipeline.
- `pipeline` is the ray tracing pipeline object containing the shaders groups.
- `group` is the index of the shader group to query.
- `groupShader` is the type of shader from the group to query.

## [](#_description)Description

The return value is the ray tracing pipeline stack size in bytes for the specified shader as called from the specified shader group.

Valid Usage

- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-04622)VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-04622  
  `pipeline` **must** be a ray tracing pipeline
- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-group-03608)VUID-vkGetRayTracingShaderGroupStackSizeKHR-group-03608  
  The value of `group` **must** be less than the number of shader groups in `pipeline`
- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-03609)VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-03609  
  The shader identified by `groupShader` in `group` **must** not be `VK_SHADER_UNUSED_KHR`

Valid Usage (Implicit)

- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-device-parameter)VUID-vkGetRayTracingShaderGroupStackSizeKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parameter)VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parameter  
  `pipeline` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-parameter)VUID-vkGetRayTracingShaderGroupStackSizeKHR-groupShader-parameter  
  `groupShader` **must** be a valid [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderGroupShaderKHR.html) value
- [](#VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parent)VUID-vkGetRayTracingShaderGroupStackSizeKHR-pipeline-parent  
  `pipeline` **must** have been created, allocated, or retrieved from `device`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html), [VkShaderGroupShaderKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderGroupShaderKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetRayTracingShaderGroupStackSizeKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0