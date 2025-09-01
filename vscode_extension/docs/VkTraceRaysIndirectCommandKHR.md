# VkTraceRaysIndirectCommandKHR(3) Manual Page

## Name

VkTraceRaysIndirectCommandKHR - Structure specifying the parameters of an indirect ray tracing command



## [](#_c_specification)C Specification

The `VkTraceRaysIndirectCommandKHR` structure is defined as:

```c++
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkTraceRaysIndirectCommandKHR {
    uint32_t    width;
    uint32_t    height;
    uint32_t    depth;
} VkTraceRaysIndirectCommandKHR;
```

## [](#_members)Members

- `width` is the width of the ray trace query dimensions.
- `height` is height of the ray trace query dimensions.
- `depth` is depth of the ray trace query dimensions.

## [](#_description)Description

The members of `VkTraceRaysIndirectCommandKHR` have the same meaning as the similarly named parameters of [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdTraceRaysKHR.html).

Valid Usage

- [](#VUID-VkTraceRaysIndirectCommandKHR-width-03638)VUID-VkTraceRaysIndirectCommandKHR-width-03638  
  `width` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[0]
- [](#VUID-VkTraceRaysIndirectCommandKHR-height-03639)VUID-VkTraceRaysIndirectCommandKHR-height-03639  
  `height` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[1]
- [](#VUID-VkTraceRaysIndirectCommandKHR-depth-03640)VUID-VkTraceRaysIndirectCommandKHR-depth-03640  
  `depth` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2] × `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[2]
- [](#VUID-VkTraceRaysIndirectCommandKHR-width-03641)VUID-VkTraceRaysIndirectCommandKHR-width-03641  
  `width` × `height` × `depth` **must** be less than or equal to `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxRayDispatchInvocationCount`

## [](#_see_also)See Also

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkTraceRaysIndirectCommandKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0