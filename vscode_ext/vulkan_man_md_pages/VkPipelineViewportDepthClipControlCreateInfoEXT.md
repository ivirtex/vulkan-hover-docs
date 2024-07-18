# VkPipelineViewportDepthClipControlCreateInfoEXT(3) Manual Page

## Name

VkPipelineViewportDepthClipControlCreateInfoEXT - Structure specifying
parameters of a newly created pipeline depth clip control state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineViewportDepthClipControlCreateInfoEXT` structure is
defined as:

``` c
// Provided by VK_EXT_depth_clip_control
typedef struct VkPipelineViewportDepthClipControlCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           negativeOneToOne;
} VkPipelineViewportDepthClipControlCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `negativeOneToOne` sets the z<sub>m</sub> in the *view volume* to
  -w<sub>c</sub>

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-negativeOneToOne-06470"
  id="VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-negativeOneToOne-06470"></a>
  VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-negativeOneToOne-06470  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthClipControl"
  target="_blank" rel="noopener"><code>depthClipControl</code></a> is
  not enabled, `negativeOneToOne` **must** be `VK_FALSE`

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-sType-sType"
  id="VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-sType-sType"></a>
  VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLIP_CONTROL_CREATE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_depth_clip_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_depth_clip_control.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportDepthClipControlCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
