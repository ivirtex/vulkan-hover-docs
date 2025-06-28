# VkPipelineRasterizationConservativeStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationConservativeStateCreateInfoEXT - Structure specifying conservative raster state



## [](#_c_specification)C Specification

If the `pNext` chain of [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationStateCreateInfo.html) includes a `VkPipelineRasterizationConservativeStateCreateInfoEXT` structure, then that structure includes parameters controlling conservative rasterization.

`VkPipelineRasterizationConservativeStateCreateInfoEXT` is defined as:

```c++
// Provided by VK_EXT_conservative_rasterization
typedef struct VkPipelineRasterizationConservativeStateCreateInfoEXT {
    VkStructureType                                           sType;
    const void*                                               pNext;
    VkPipelineRasterizationConservativeStateCreateFlagsEXT    flags;
    VkConservativeRasterizationModeEXT                        conservativeRasterizationMode;
    float                                                     extraPrimitiveOverestimationSize;
} VkPipelineRasterizationConservativeStateCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `flags` is reserved for future use.
- `conservativeRasterizationMode` is the conservative rasterization mode to use.
- `extraPrimitiveOverestimationSize` is the extra size in pixels to increase the generating primitive during conservative rasterization at each of its edges in `X` and `Y` equally in screen space beyond the base overestimation specified in `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`primitiveOverestimationSize`. If `conservativeRasterizationMode` is not `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT`, this value is ignored.

## [](#_description)Description

If this structure is not included in the `pNext` chain, `conservativeRasterizationMode` is considered to be `VK_CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT`, and conservative rasterization is disabled.

Polygon rasterization **can** be made conservative by setting `conservativeRasterizationMode` to `VK_CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT` or `VK_CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT` in `VkPipelineRasterizationConservativeStateCreateInfoEXT`.

Note

If [`conservativePointAndLineRasterization`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-conservativePointAndLineRasterization) is supported, conservative rasterization can be applied to line and point primitives, otherwise it must be disabled.

Valid Usage

- [](#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-extraPrimitiveOverestimationSize-01769)VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-extraPrimitiveOverestimationSize-01769  
  `extraPrimitiveOverestimationSize` **must** be in the range of `0.0` to `VkPhysicalDeviceConservativeRasterizationPropertiesEXT`::`maxExtraPrimitiveOverestimationSize` inclusive

Valid Usage (Implicit)

- [](#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-sType-sType)VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_CONSERVATIVE_STATE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-flags-zerobitmask)VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`
- [](#VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-conservativeRasterizationMode-parameter)VUID-VkPipelineRasterizationConservativeStateCreateInfoEXT-conservativeRasterizationMode-parameter  
  `conservativeRasterizationMode` **must** be a valid [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConservativeRasterizationModeEXT.html) value

## [](#_see_also)See Also

[VK\_EXT\_conservative\_rasterization](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_conservative_rasterization.html), [VkConservativeRasterizationModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConservativeRasterizationModeEXT.html), [VkPipelineRasterizationConservativeStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineRasterizationConservativeStateCreateFlagsEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineRasterizationConservativeStateCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0