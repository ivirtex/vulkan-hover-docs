# VkPipelineViewportDepthClampControlCreateInfoEXT(3) Manual Page

## Name

VkPipelineViewportDepthClampControlCreateInfoEXT - Structure specifying parameters of a newly created pipeline depth clamp control state



## [](#_c_specification)C Specification

The `VkPipelineViewportDepthClampControlCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_clamp_control
typedef struct VkPipelineViewportDepthClampControlCreateInfoEXT {
    VkStructureType                sType;
    const void*                    pNext;
    VkDepthClampModeEXT            depthClampMode;
    const VkDepthClampRangeEXT*    pDepthClampRange;
} VkPipelineViewportDepthClampControlCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `depthClampMode` determines how the clamp range is determined for each viewport.
- `pDepthClampRange` sets the depth clamp range for all viewports if `depthClampMode` is `VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT`.

## [](#_description)Description

This structure extends `VkPipelineViewportStateCreateInfo` and specifies the depth clamp range used in the pipeline. If this structure is not provided in the next chain then `depthClampMode` defaults to `VK_DEPTH_CLAMP_MODE_VIEWPORT_RANGE_EXT`.

Valid Usage

- [](#VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-pDepthClampRange-09646)VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-pDepthClampRange-09646  
  If `depthClampMode` is `VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT`, and the pipeline is not created with `VK_DYNAMIC_STATE_DEPTH_CLAMP_RANGE_EXT`, then `pDepthClampRange` **must** be a valid pointer to a valid `VkDepthClampRangeEXT` structure

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-sType-sType)VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLAMP_CONTROL_CREATE_INFO_EXT`
- [](#VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-depthClampMode-parameter)VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-depthClampMode-parameter  
  `depthClampMode` **must** be a valid [VkDepthClampModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampModeEXT.html) value
- [](#VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-pDepthClampRange-parameter)VUID-VkPipelineViewportDepthClampControlCreateInfoEXT-pDepthClampRange-parameter  
  If `pDepthClampRange` is not `NULL`, `pDepthClampRange` **must** be a valid pointer to a valid [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_depth\_clamp\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_control.html), [VkDepthClampModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampModeEXT.html), [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportDepthClampControlCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0