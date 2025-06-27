# VkTraceRaysIndirectCommandKHR(3) Manual Page

## Name

VkTraceRaysIndirectCommandKHR - Structure specifying the parameters of
an indirect ray tracing command



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkTraceRaysIndirectCommandKHR` structure is defined as:

``` c
// Provided by VK_KHR_ray_tracing_pipeline
typedef struct VkTraceRaysIndirectCommandKHR {
    uint32_t    width;
    uint32_t    height;
    uint32_t    depth;
} VkTraceRaysIndirectCommandKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `width` is the width of the ray trace query dimensions.

- `height` is height of the ray trace query dimensions.

- `depth` is depth of the ray trace query dimensions.

## <a href="#_description" class="anchor"></a>Description

The members of `VkTraceRaysIndirectCommandKHR` have the same meaning as
the similarly named parameters of
[vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html).

Valid Usage

- <a href="#VUID-VkTraceRaysIndirectCommandKHR-width-03638"
  id="VUID-VkTraceRaysIndirectCommandKHR-width-03638"></a>
  VUID-VkTraceRaysIndirectCommandKHR-width-03638  
  `width` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[0\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[0\]

- <a href="#VUID-VkTraceRaysIndirectCommandKHR-height-03639"
  id="VUID-VkTraceRaysIndirectCommandKHR-height-03639"></a>
  VUID-VkTraceRaysIndirectCommandKHR-height-03639  
  `height` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[1\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[1\]

- <a href="#VUID-VkTraceRaysIndirectCommandKHR-depth-03640"
  id="VUID-VkTraceRaysIndirectCommandKHR-depth-03640"></a>
  VUID-VkTraceRaysIndirectCommandKHR-depth-03640  
  `depth` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupCount`\[2\] ×
  `VkPhysicalDeviceLimits`::`maxComputeWorkGroupSize`\[2\]

- <a href="#VUID-VkTraceRaysIndirectCommandKHR-width-03641"
  id="VUID-VkTraceRaysIndirectCommandKHR-width-03641"></a>
  VUID-VkTraceRaysIndirectCommandKHR-width-03641  
  `width` × `height` × `depth` **must** be less than or equal to
  `VkPhysicalDeviceRayTracingPipelinePropertiesKHR`::`maxRayDispatchInvocationCount`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkTraceRaysIndirectCommandKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
