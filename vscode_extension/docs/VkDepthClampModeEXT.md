# VkDepthClampModeEXT(3) Manual Page

## Name

VkDepthClampModeEXT - Modes that determine the depth clamp range



## [](#_c_specification)C Specification

Possible values of [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html)::`depthClampMode`, specifying which range should be used for depth clamping, are:

```c++
// Provided by VK_EXT_depth_clamp_control
typedef enum VkDepthClampModeEXT {
    VK_DEPTH_CLAMP_MODE_VIEWPORT_RANGE_EXT = 0,
    VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT = 1,
} VkDepthClampModeEXT;
```

## [](#_description)Description

- `VK_DEPTH_CLAMP_MODE_VIEWPORT_RANGE_EXT` specifies that the depth clamp range follows the viewport depth range. The depth clamp range of each viewport will implicitly be set to zmin = min(n,f) and zmax = max(n,f), where n and f are the `minDepth` and `maxDepth` depth range values of the viewport.
- `VK_DEPTH_CLAMP_MODE_USER_DEFINED_RANGE_EXT` specifies that a single user-defined depth clamp range will be used for all viewports. The user-defined depth clamp range is defined by the `minDepthClamp` and `maxDepthClamp` members of [VkDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDepthClampRangeEXT.html).

## [](#_see_also)See Also

[VK\_EXT\_depth\_clamp\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clamp_control.html), [VkPipelineViewportDepthClampControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClampControlCreateInfoEXT.html), [vkCmdSetDepthClampRangeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetDepthClampRangeEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDepthClampModeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0