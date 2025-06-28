# VkViewport(3) Manual Page

## Name

VkViewport - Structure specifying a viewport



## [](#_c_specification)C Specification

The `VkViewport` structure is defined as:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkViewport {
    float    x;
    float    y;
    float    width;
    float    height;
    float    minDepth;
    float    maxDepth;
} VkViewport;
```

## [](#_members)Members

- `x` and `y` are the viewport’s upper left corner (x,y).
- `width` and `height` are the viewport’s width and height, respectively.
- `minDepth` and `maxDepth` are the depth range for the viewport.

## [](#_description)Description

Note

Despite their names, `minDepth` **can** be less than, equal to, or greater than `maxDepth`.

The framebuffer depth coordinate `z`f **may** be represented using either a fixed-point or floating-point representation. However, a floating-point representation **must** be used if the depth/stencil attachment has a floating-point depth component. If an m-bit fixed-point representation is used, we assume that it represents each value 2m−1k​, where k ∈ { 0, 1, …​, 2m-1 }, as k (e.g. 1.0 is represented in binary as a string of all ones).

The viewport parameters shown in the above equations are found from these values as

ox = `x` + `width` / 2

oy = `y` + `height` / 2

oz = `minDepth` (or (`maxDepth` + `minDepth`) / 2 if [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)::`negativeOneToOne` is `VK_TRUE`)

px = `width`

py = `height`

pz = `maxDepth` - `minDepth` (or (`maxDepth` - `minDepth`) / 2 if [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)::`negativeOneToOne` is `VK_TRUE`)

If a render pass transform is enabled, the values (px,py) and (ox, oy) defining the viewport are transformed as described in [render pass transform](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vertexpostproc-renderpass-transform) before participating in the viewport transform.

The application **can** specify a negative term for `height`, which has the effect of negating the y coordinate in clip space before performing the transform. When using a negative `height`, the application **should** also adjust the `y` value to point to the lower left corner of the viewport instead of the upper left corner. Using the negative `height` allows the application to avoid having to negate the y component of the `Position` output from the last [pre-rasterization shader stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization).

The width and height of the [implementation-dependent maximum viewport dimensions](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxViewportDimensions) **must** be greater than or equal to the width and height of the largest image which **can** be created and attached to a framebuffer.

The floating-point viewport bounds are represented with an [implementation-dependent precision](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-viewportSubPixelBits).

Valid Usage

- [](#VUID-VkViewport-width-01770)VUID-VkViewport-width-01770  
  `width` **must** be greater than `0.0`
- [](#VUID-VkViewport-width-01771)VUID-VkViewport-width-01771  
  `width` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewportDimensions`\[0]
- [](#VUID-VkViewport-apiVersion-07917)VUID-VkViewport-apiVersion-07917  
  If the [VK\_KHR\_maintenance1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance1.html) extension is not enabled, the [VK\_AMD\_negative\_viewport\_height](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_negative_viewport_height.html) extension is not enabled, and [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties.html)::`apiVersion` is less than Vulkan 1.1, `height` **must** be greater than `0.0`
- [](#VUID-VkViewport-height-01773)VUID-VkViewport-height-01773  
  The absolute value of `height` **must** be less than or equal to `VkPhysicalDeviceLimits`::`maxViewportDimensions`\[1]
- [](#VUID-VkViewport-x-01774)VUID-VkViewport-x-01774  
  `x` **must** be greater than or equal to `viewportBoundsRange`\[0]
- [](#VUID-VkViewport-x-01232)VUID-VkViewport-x-01232  
  (`x` + `width`) **must** be less than or equal to `viewportBoundsRange`\[1]
- [](#VUID-VkViewport-y-01775)VUID-VkViewport-y-01775  
  `y` **must** be greater than or equal to `viewportBoundsRange`\[0]
- [](#VUID-VkViewport-y-01776)VUID-VkViewport-y-01776  
  `y` **must** be less than or equal to `viewportBoundsRange`\[1]
- [](#VUID-VkViewport-y-01777)VUID-VkViewport-y-01777  
  (`y` + `height`) **must** be greater than or equal to `viewportBoundsRange`\[0]
- [](#VUID-VkViewport-y-01233)VUID-VkViewport-y-01233  
  (`y` + `height`) **must** be less than or equal to `viewportBoundsRange`\[1]
- [](#VUID-VkViewport-minDepth-01234)VUID-VkViewport-minDepth-01234  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled, `minDepth` **must** be between `0.0` and `1.0`, inclusive
- [](#VUID-VkViewport-maxDepth-01235)VUID-VkViewport-maxDepth-01235  
  If the `VK_EXT_depth_range_unrestricted` extension is not enabled, `maxDepth` **must** be between `0.0` and `1.0`, inclusive

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html), [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineViewportStateCreateInfo.html), [vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewport.html), [vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWithCount.html), [vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdSetViewportWithCountEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkViewport)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0