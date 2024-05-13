# VkMultiviewPerViewAttributesInfoNVX(3) Manual Page

## Name

VkMultiviewPerViewAttributesInfoNVX - Structure specifying the multiview
per-attribute properties



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMultiviewPerViewAttributesInfoNVX` structure is defined as:

``` c
// Provided by VK_KHR_dynamic_rendering with VK_NVX_multiview_per_view_attributes
typedef struct VkMultiviewPerViewAttributesInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           perViewAttributes;
    VkBool32           perViewAttributesPositionXOnly;
} VkMultiviewPerViewAttributesInfoNVX;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `perViewAttributes` specifies that shaders compiled for this pipeline
  write the attributes for all views in a single invocation of each
  vertex processing stage. All pipelines executed within a render pass
  instance that includes this bit **must** write per-view attributes to
  the `*PerViewNV[]` shader outputs, in addition to the non-per-view
  (e.g. `Position`) outputs.

- `perViewAttributesPositionXOnly` specifies that shaders compiled for
  this pipeline use per-view positions which only differ in value in the
  x component. Per-view viewport mask **can** also be used.

## <a href="#_description" class="anchor"></a>Description

When dynamic render pass instances are being used, instead of specifying
`VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX` or
`VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX` in the subpass
description flags, the per-attribute properties of the render pass
instance **must** be specified by the
`VkMultiviewPerViewAttributesInfoNVX` structure Include the
`VkMultiviewPerViewAttributesInfoNVX` structure in the `pNext` chain of
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html) when
creating a graphics pipeline for dynamic rendering,
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) when starting a dynamic render
pass instance, and
[VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceInfo.html)
when specifying the dynamic render pass instance parameters for
secondary command buffers.

Valid Usage (Implicit)

- <a href="#VUID-VkMultiviewPerViewAttributesInfoNVX-sType-sType"
  id="VUID-VkMultiviewPerViewAttributesInfoNVX-sType-sType"></a>
  VUID-VkMultiviewPerViewAttributesInfoNVX-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MULTIVIEW_PER_VIEW_ATTRIBUTES_INFO_NVX`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_dynamic_rendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_dynamic_rendering.html),
[VK_NVX_multiview_per_view_attributes](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NVX_multiview_per_view_attributes.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMultiviewPerViewAttributesInfoNVX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
