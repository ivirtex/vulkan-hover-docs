# VkMultiviewPerViewAttributesInfoNVX(3) Manual Page

## Name

VkMultiviewPerViewAttributesInfoNVX - Structure specifying the multiview per-attribute properties



## [](#_c_specification)C Specification

The `VkMultiviewPerViewAttributesInfoNVX` structure is defined as:

```c++
// Provided by VK_NVX_multiview_per_view_attributes with VK_VERSION_1_3 or VK_KHR_dynamic_rendering
typedef struct VkMultiviewPerViewAttributesInfoNVX {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           perViewAttributes;
    VkBool32           perViewAttributesPositionXOnly;
} VkMultiviewPerViewAttributesInfoNVX;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `perViewAttributes` specifies that shaders compiled for this pipeline write the attributes for all views in a single invocation of each vertex processing stage. All pipelines executed within a render pass instance that includes this bit **must** write per-view attributes to the `*PerViewNV[]` shader outputs, in addition to the non-per-view (e.g. `Position`) outputs.
- `perViewAttributesPositionXOnly` specifies that shaders compiled for this pipeline use per-view positions which only differ in value in the x component. Per-view viewport mask **can** also be used.

## [](#_description)Description

When dynamic render pass instances are being used, instead of specifying `VK_SUBPASS_DESCRIPTION_PER_VIEW_ATTRIBUTES_BIT_NVX` or `VK_SUBPASS_DESCRIPTION_PER_VIEW_POSITION_X_ONLY_BIT_NVX` in the subpass description flags, the per-attribute properties of the render pass instance **must** be specified by the `VkMultiviewPerViewAttributesInfoNVX` structure Include the `VkMultiviewPerViewAttributesInfoNVX` structure in the `pNext` chain of [VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGraphicsPipelineCreateInfo.html) when creating a graphics pipeline for dynamic rendering, [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) when starting a dynamic render pass instance, and [VkCommandBufferInheritanceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceInfo.html) when specifying the dynamic render pass instance parameters for secondary command buffers.

Valid Usage (Implicit)

- [](#VUID-VkMultiviewPerViewAttributesInfoNVX-sType-sType)VUID-VkMultiviewPerViewAttributesInfoNVX-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MULTIVIEW_PER_VIEW_ATTRIBUTES_INFO_NVX`

## [](#_see_also)See Also

[VK\_KHR\_dynamic\_rendering](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_dynamic_rendering.html), [VK\_NVX\_multiview\_per\_view\_attributes](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NVX_multiview_per_view_attributes.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMultiviewPerViewAttributesInfoNVX)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0