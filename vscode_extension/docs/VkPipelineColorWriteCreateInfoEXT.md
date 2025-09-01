# VkPipelineColorWriteCreateInfoEXT(3) Manual Page

## Name

VkPipelineColorWriteCreateInfoEXT - Structure specifying color write state of a newly created pipeline



## [](#_c_specification)C Specification

The `VkPipelineColorWriteCreateInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_color_write_enable
typedef struct VkPipelineColorWriteCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           attachmentCount;
    const VkBool32*    pColorWriteEnables;
} VkPipelineColorWriteCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `attachmentCount` is the number of [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) elements in `pColorWriteEnables`.
- `pColorWriteEnables` is a pointer to an array of per target attachment boolean values specifying whether color writes are enabled for the given attachment.

## [](#_description)Description

When this structure is included in the `pNext` chain of [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html), it defines per-attachment color write state. If this structure is not included in the `pNext` chain, it is equivalent to specifying this structure with `attachmentCount` equal to the `attachmentCount` member of [VkPipelineColorBlendStateCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendStateCreateInfo.html), and `pColorWriteEnables` pointing to an array of as many `VK_TRUE` values.

If the [`colorWriteEnable`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-colorWriteEnable) feature is not enabled, all [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) elements in the `pColorWriteEnables` array **must** be `VK_TRUE`.

Color Write Enable interacts with the [Color Write Mask](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-color-write-mask) as follows:

- If `colorWriteEnable` is `VK_TRUE`, writes to the attachment are determined by the `colorWriteMask`.
- If `colorWriteEnable` is `VK_FALSE`, the `colorWriteMask` is ignored and writes to all components of the attachment are disabled. This is equivalent to specifying a `colorWriteMask` of 0.

Valid Usage

- [](#VUID-VkPipelineColorWriteCreateInfoEXT-pAttachments-04801)VUID-VkPipelineColorWriteCreateInfoEXT-pAttachments-04801  
  If the [`colorWriteEnable`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-colorWriteEnable) feature is not enabled, all elements of `pColorWriteEnables` **must** be `VK_TRUE`
- [](#VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-07608)VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-07608  
  If the pipeline is being created with `VK_DYNAMIC_STATE_COLOR_BLEND_ADVANCED_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_ENABLE_EXT`, `VK_DYNAMIC_STATE_COLOR_BLEND_EQUATION_EXT`, or `VK_DYNAMIC_STATE_COLOR_WRITE_MASK_EXT` dynamic states not set, `attachmentCount` **must** be equal to the `attachmentCount` member of the `VkPipelineColorBlendStateCreateInfo` structure specified during pipeline creation
- [](#VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-06655)VUID-VkPipelineColorWriteCreateInfoEXT-attachmentCount-06655  
  `attachmentCount` **must** be less than or equal to the `maxColorAttachments` member of `VkPhysicalDeviceLimits`

Valid Usage (Implicit)

- [](#VUID-VkPipelineColorWriteCreateInfoEXT-sType-sType)VUID-VkPipelineColorWriteCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_COLOR_WRITE_CREATE_INFO_EXT`
- [](#VUID-VkPipelineColorWriteCreateInfoEXT-pColorWriteEnables-parameter)VUID-VkPipelineColorWriteCreateInfoEXT-pColorWriteEnables-parameter  
  If `attachmentCount` is not `0`, `pColorWriteEnables` **must** be a valid pointer to an array of `attachmentCount` [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html) values

## [](#_see_also)See Also

[VK\_EXT\_color\_write\_enable](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_color_write_enable.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineColorWriteCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0