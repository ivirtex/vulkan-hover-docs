# VkPipelineRasterizationDepthClipStateCreateInfoEXT(3) Manual Page

## Name

VkPipelineRasterizationDepthClipStateCreateInfoEXT - Structure
specifying depth clipping state



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)
includes a `VkPipelineRasterizationDepthClipStateCreateInfoEXT`
structure, then that structure controls whether depth clipping is
enabled or disabled.

The `VkPipelineRasterizationDepthClipStateCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_depth_clip_enable
typedef struct VkPipelineRasterizationDepthClipStateCreateInfoEXT {
    VkStructureType                                        sType;
    const void*                                            pNext;
    VkPipelineRasterizationDepthClipStateCreateFlagsEXT    flags;
    VkBool32                                               depthClipEnable;
} VkPipelineRasterizationDepthClipStateCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `depthClipEnable` controls whether depth clipping is enabled as
  described in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vertexpostproc-clipping"
  target="_blank" rel="noopener">Primitive Clipping</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_RASTERIZATION_DEPTH_CLIP_STATE_CREATE_INFO_EXT`

- <a
  href="#VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-flags-zerobitmask"
  id="VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-flags-zerobitmask"></a>
  VUID-VkPipelineRasterizationDepthClipStateCreateInfoEXT-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_clip_enable](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_enable.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkPipelineRasterizationDepthClipStateCreateFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationDepthClipStateCreateFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineRasterizationDepthClipStateCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
