# VkViewport(3) Manual Page

## Name

VkViewport - Structure specifying a viewport



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkViewport` structure is defined as:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `x` and `y` are the viewport’s upper left corner (x,y).

- `width` and `height` are the viewport’s width and height,
  respectively.

- `minDepth` and `maxDepth` are the depth range for the viewport.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Despite their names, <code>minDepth</code> <strong>can</strong> be
less than, equal to, or greater than <code>maxDepth</code>.</p></td>
</tr>
</tbody>
</table>

The framebuffer depth coordinate `z`<sub>f</sub> **may** be represented
using either a fixed-point or floating-point representation. However, a
floating-point representation **must** be used if the depth/stencil
attachment has a floating-point depth component. If an m-bit fixed-point
representation is used, we assume that it represents each value 2m−1k​,
where k ∈ { 0, 1, …​, 2<sup>m</sup>-1 }, as k (e.g. 1.0 is represented in
binary as a string of all ones).

The viewport parameters shown in the above equations are found from
these values as

  
o<sub>x</sub> = `x` + `width` / 2

  
o<sub>y</sub> = `y` + `height` / 2

  
o<sub>z</sub> = `minDepth` (or (`maxDepth` + `minDepth`) / 2 if
[VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)::`negativeOneToOne`
is `VK_TRUE`)

  
p<sub>x</sub> = `width`

  
p<sub>y</sub> = `height`

  
p<sub>z</sub> = `maxDepth` - `minDepth` (or (`maxDepth` - `minDepth`) /
2 if
[VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html)::`negativeOneToOne`
is `VK_TRUE`)

If a render pass transform is enabled, the values
(p<sub>x</sub>,p<sub>y</sub>) and (o<sub>x</sub>, o<sub>y</sub>)
defining the viewport are transformed as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-renderpass-transform"
target="_blank" rel="noopener">render pass transform</a> before
participating in the viewport transform.

The application **can** specify a negative term for `height`, which has
the effect of negating the y coordinate in clip space before performing
the transform. When using a negative `height`, the application
**should** also adjust the `y` value to point to the lower left corner
of the viewport instead of the upper left corner. Using the negative
`height` allows the application to avoid having to negate the y
component of the `Position` output from the last <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-graphics-subsets-pre-rasterization"
target="_blank" rel="noopener">pre-rasterization shader stage</a>.

The width and height of the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxViewportDimensions"
target="_blank" rel="noopener">implementation-dependent maximum viewport
dimensions</a> **must** be greater than or equal to the width and height
of the largest image which **can** be created and attached to a
framebuffer.

The floating-point viewport bounds are represented with an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-viewportSubPixelBits"
target="_blank" rel="noopener">implementation-dependent precision</a>.

Valid Usage

- <a href="#VUID-VkViewport-width-01770"
  id="VUID-VkViewport-width-01770"></a> VUID-VkViewport-width-01770  
  `width` **must** be greater than `0.0`

- <a href="#VUID-VkViewport-width-01771"
  id="VUID-VkViewport-width-01771"></a> VUID-VkViewport-width-01771  
  `width` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewportDimensions`\[0\]

- <a href="#VUID-VkViewport-apiVersion-07917"
  id="VUID-VkViewport-apiVersion-07917"></a>
  VUID-VkViewport-apiVersion-07917  
  If the [VK_KHR_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance1.html) extension is
  not enabled, the
  [VK_AMD_negative_viewport_height](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_negative_viewport_height.html)
  extension is not enabled, and
  [VkPhysicalDeviceProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties.html)::`apiVersion`
  is less than Vulkan 1.1, `height` **must** be greater than `0.0`

- <a href="#VUID-VkViewport-height-01773"
  id="VUID-VkViewport-height-01773"></a> VUID-VkViewport-height-01773  
  The absolute value of `height` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewportDimensions`\[1\]

- <a href="#VUID-VkViewport-x-01774" id="VUID-VkViewport-x-01774"></a>
  VUID-VkViewport-x-01774  
  `x` **must** be greater than or equal to `viewportBoundsRange`\[0\]

- <a href="#VUID-VkViewport-x-01232" id="VUID-VkViewport-x-01232"></a>
  VUID-VkViewport-x-01232  
  (`x` + `width`) **must** be less than or equal to
  `viewportBoundsRange`\[1\]

- <a href="#VUID-VkViewport-y-01775" id="VUID-VkViewport-y-01775"></a>
  VUID-VkViewport-y-01775  
  `y` **must** be greater than or equal to `viewportBoundsRange`\[0\]

- <a href="#VUID-VkViewport-y-01776" id="VUID-VkViewport-y-01776"></a>
  VUID-VkViewport-y-01776  
  `y` **must** be less than or equal to `viewportBoundsRange`\[1\]

- <a href="#VUID-VkViewport-y-01777" id="VUID-VkViewport-y-01777"></a>
  VUID-VkViewport-y-01777  
  (`y` + `height`) **must** be greater than or equal to
  `viewportBoundsRange`\[0\]

- <a href="#VUID-VkViewport-y-01233" id="VUID-VkViewport-y-01233"></a>
  VUID-VkViewport-y-01233  
  (`y` + `height`) **must** be less than or equal to
  `viewportBoundsRange`\[1\]

- <a href="#VUID-VkViewport-minDepth-01234"
  id="VUID-VkViewport-minDepth-01234"></a>
  VUID-VkViewport-minDepth-01234  
  If the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is not enabled, `minDepth` **must** be between `0.0` and
  `1.0`, inclusive

- <a href="#VUID-VkViewport-maxDepth-01235"
  id="VUID-VkViewport-maxDepth-01235"></a>
  VUID-VkViewport-maxDepth-01235  
  If the
  [`VK_EXT_depth_range_unrestricted`](VK_EXT_depth_range_unrestricted.html)
  extension is not enabled, `maxDepth` **must** be between `0.0` and
  `1.0`, inclusive

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html),
[VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html),
[vkCmdSetViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewport.html),
[vkCmdSetViewportWithCount](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCount.html),
[vkCmdSetViewportWithCountEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetViewportWithCountEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkViewport"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
