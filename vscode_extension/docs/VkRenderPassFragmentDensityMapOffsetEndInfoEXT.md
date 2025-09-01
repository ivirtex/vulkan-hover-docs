# VkRenderPassFragmentDensityMapOffsetEndInfoEXT(3) Manual Page

## Name

VkRenderPassFragmentDensityMapOffsetEndInfoEXT - Structure specifying fragment density map offset subpass end information



## [](#_c_specification)C Specification

If the [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html)::`pNext` chain or [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html)::`pNext` chain includes a `VkRenderPassFragmentDensityMapOffsetEndInfoEXT` structure, then that structure includes an array of fragment density map offsets per layer for the render pass.

The `VkRenderPassFragmentDensityMapOffsetEndInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_fragment_density_map_offset
typedef struct VkRenderPassFragmentDensityMapOffsetEndInfoEXT {
    VkStructureType      sType;
    const void*          pNext;
    uint32_t             fragmentDensityOffsetCount;
    const VkOffset2D*    pFragmentDensityOffsets;
} VkRenderPassFragmentDensityMapOffsetEndInfoEXT;
```

or the equivalent:

```c++
// Provided by VK_QCOM_fragment_density_map_offset
typedef VkRenderPassFragmentDensityMapOffsetEndInfoEXT VkSubpassFragmentDensityMapOffsetEndInfoQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `fragmentDensityOffsetCount` is the number of offsets being specified.
- `pFragmentDensityOffsets` is a pointer to an array of [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html) structs, each of which describes the offset per layer.

## [](#_description)Description

The array elements are given per `layer` as defined by [Fetch Density Value](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymap-fetch-density-value), where index = layer. Each (x,y) offset is in framebuffer pixels and shifts the fetch of the fragment density map by that amount. Offsets can be positive or negative.

If neither the [VkSubpassEndInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassEndInfo.html)::`pNext` chain for the last subpass of a render pass nor the [VkRenderingEndInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingEndInfoEXT.html)::`pNext` chain of a dynamic render pass include `VkRenderPassFragmentDensityMapOffsetEndInfoEXT`, or if `fragmentDensityOffsetCount` is zero, then the offset (0,0) is used for [Fetch Density Value](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#fragmentdensitymap-fetch-density-value).

Valid Usage

- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityMapOffsets-06503)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityMapOffsets-06503  
  If the [`fragmentDensityMapOffset`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-fragmentDensityMapOffset) feature is not enabled or fragment density map is not enabled in the render pass, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityMapAttachment-06504)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityMapAttachment-06504  
  If [VkRenderPassFragmentDensityMapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassFragmentDensityMapCreateInfoEXT.html)::`fragmentDensityMapAttachment` is not `VK_ATTACHMENT_UNUSED` and was not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pDepthStencilAttachment-06505)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pDepthStencilAttachment-06505  
  If the depth or stencil attachments for the render pass are used and were not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pInputAttachments-06506)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pInputAttachments-06506  
  If any used input attachments for the render pass were not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pColorAttachments-06507)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pColorAttachments-06507  
  If any used color attachments for the render pass were not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pResolveAttachments-06508)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pResolveAttachments-06508  
  If any used resolve attachments for the render pass were not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pPreserveAttachments-06509)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pPreserveAttachments-06509  
  If any used preserve attachments for the render pass were not created with `VK_IMAGE_CREATE_FRAGMENT_DENSITY_MAP_OFFSET_BIT_EXT`, `fragmentDensityOffsetCount` **must** equal `0`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityOffsetCount-06510)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityOffsetCount-06510  
  If `fragmentDensityOffsetCount` is not `0` and multiview is enabled for the render pass, `fragmentDensityOffsetCount` **must** equal the `layerCount` that was specified in creating the fragment density map attachment view
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityOffsetCount-06511)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-fragmentDensityOffsetCount-06511  
  If `fragmentDensityOffsetCount` is not `0` and multiview is not enabled for the render pass, `fragmentDensityOffsetCount` **must** equal `1`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-x-06512)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-x-06512  
  The `x` component of each element of `pFragmentDensityOffsets` **must** be an integer multiple of `fragmentDensityOffsetGranularity.width`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-y-06513)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-y-06513  
  The `y` component of each element of `pFragmentDensityOffsets` **must** be an integer multiple of `fragmentDensityOffsetGranularity.height`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pFragmentDensityOffsets-10730)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pFragmentDensityOffsets-10730  
  Each element of `pFragmentDensityOffsets` must be identical for every [vkCmdEndRendering2EXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdEndRendering2EXT.html) call made in a render pass

Valid Usage (Implicit)

- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-sType-sType)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_FRAGMENT_DENSITY_MAP_OFFSET_END_INFO_EXT`
- [](#VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pFragmentDensityOffsets-parameter)VUID-VkRenderPassFragmentDensityMapOffsetEndInfoEXT-pFragmentDensityOffsets-parameter  
  If `fragmentDensityOffsetCount` is not `0`, `pFragmentDensityOffsets` **must** be a valid pointer to an array of `fragmentDensityOffsetCount` [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html) structures

## [](#_see_also)See Also

[VK\_EXT\_fragment\_density\_map\_offset](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_fragment_density_map_offset.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkOffset2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassFragmentDensityMapOffsetEndInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0