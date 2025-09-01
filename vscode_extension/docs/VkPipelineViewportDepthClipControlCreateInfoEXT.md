# VkPipelineViewportDepthClipControlCreateInfoEXT(3) Manual Page

## Name

VkPipelineViewportDepthClipControlCreateInfoEXT - Structure specifying parameters of a newly created pipeline depth clip control state



## [](#_c_specification)C Specification

The `VkPipelineViewportDepthClipControlCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_depth_clip_control
typedef struct VkPipelineViewportDepthClipControlCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           negativeOneToOne;
} VkPipelineViewportDepthClipControlCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `negativeOneToOne` sets the zm in the *view volume* to -wc

## [](#_description)Description

Valid Usage

- [](#VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-negativeOneToOne-06470)VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-negativeOneToOne-06470  
  If the [`depthClipControl`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-depthClipControl) feature is not enabled, `negativeOneToOne` **must** be `VK_FALSE`

Valid Usage (Implicit)

- [](#VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-sType-sType)VUID-VkPipelineViewportDepthClipControlCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_DEPTH_CLIP_CONTROL_CREATE_INFO_EXT`

## [](#_see_also)See Also

[VK\_EXT\_depth\_clip\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_depth_clip_control.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineViewportDepthClipControlCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0