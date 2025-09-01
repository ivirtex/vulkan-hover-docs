# VkShaderGroupShaderKHR(3) Manual Page

## Name

VkShaderGroupShaderKHR - Shader group shaders



## [](#_c_specification)C Specification

Possible values of `groupShader` in [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html) are:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef enum VkShaderGroupShaderKHR {
    VK_SHADER_GROUP_SHADER_GENERAL_KHR = 0,
    VK_SHADER_GROUP_SHADER_CLOSEST_HIT_KHR = 1,
    VK_SHADER_GROUP_SHADER_ANY_HIT_KHR = 2,
    VK_SHADER_GROUP_SHADER_INTERSECTION_KHR = 3,
} VkShaderGroupShaderKHR;
```

## [](#_description)Description

- `VK_SHADER_GROUP_SHADER_GENERAL_KHR` uses the shader specified in the group with [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`generalShader`
- `VK_SHADER_GROUP_SHADER_CLOSEST_HIT_KHR` uses the shader specified in the group with [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`closestHitShader`
- `VK_SHADER_GROUP_SHADER_ANY_HIT_KHR` uses the shader specified in the group with [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`anyHitShader`
- `VK_SHADER_GROUP_SHADER_INTERSECTION_KHR` uses the shader specified in the group with [VkRayTracingShaderGroupCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRayTracingShaderGroupCreateInfoKHR.html)::`intersectionShader`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html), [vkGetRayTracingShaderGroupStackSizeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetRayTracingShaderGroupStackSizeKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkShaderGroupShaderKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0