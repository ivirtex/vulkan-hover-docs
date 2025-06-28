# VkDepthClampRangeEXT(3) Manual Page

## Name

VkDepthClampRangeEXT - Structure specifying a depth clamp range



## [](#_c_specification)C Specification

The `VkDepthClampRangeEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_clamp_control
typedef struct VkDepthClampRangeEXT {
    float    minDepthClamp;
    float    maxDepthClamp;
} VkDepthClampRangeEXT;
```

## [](#_members)Members

- `minDepthClamp` sets zmin in the depth clamp range of the viewport.
- `maxDepthClamp` sets zmax in the depth clamp range of the viewport.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDepthClampRangeEXT-pDepthClampRange-00999)VUID-VkDepthClampRangeEXT-pDepthClampRange-00999  
  `minDepthClamp` **must** be less than or equal to `maxDepthClamp`
- [](#VUID-VkDepthClampRangeEXT-pDepthClampRange-09648)VUID-VkDepthClampRangeEXT-pDepthClampRange-09648  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled, `minDepthClamp` **must** be greater than or equal to `0.0`
- [](#VUID-VkDepthClampRangeEXT-pDepthClampRange-09649)VUID-VkDepthClampRangeEXT-pDepthClampRange-09649  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled, `maxDepthClamp` **must** be less than or equal to `1.0`

## [](#_see_also)See Also

[VK\_EXT\_depth\_clamp\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_control.html), [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html), [vkCmdSetDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampRangeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDepthClampRangeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0