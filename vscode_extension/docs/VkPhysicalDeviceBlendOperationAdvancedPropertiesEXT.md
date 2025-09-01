# VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT(3) Manual Page

## Name

VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT - Structure describing advanced blending limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT` structure is defined as:

```c++
// Provided by VK_EXT_blend_operation_advanced
typedef struct VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           advancedBlendMaxColorAttachments;
    VkBool32           advancedBlendIndependentBlend;
    VkBool32           advancedBlendNonPremultipliedSrcColor;
    VkBool32           advancedBlendNonPremultipliedDstColor;
    VkBool32           advancedBlendCorrelatedOverlap;
    VkBool32           advancedBlendAllOperations;
} VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`advancedBlendMaxColorAttachments` is one greater than the highest color attachment index that **can** be used in a render pass instance, for a pipeline that uses an [advanced blend operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#framebuffer-blend-advanced).
- []()`advancedBlendIndependentBlend` specifies whether advanced blend operations **can** vary per-attachment.
- []()`advancedBlendNonPremultipliedSrcColor` specifies whether the source color **can** be treated as non-premultiplied. If this is `VK_FALSE`, then [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`srcPremultiplied` **must** be `VK_TRUE`.
- []()`advancedBlendNonPremultipliedDstColor` specifies whether the destination color **can** be treated as non-premultiplied. If this is `VK_FALSE`, then [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`dstPremultiplied` **must** be `VK_TRUE`.
- []()`advancedBlendCorrelatedOverlap` specifies whether the overlap mode **can** be treated as correlated. If this is `VK_FALSE`, then [VkPipelineColorBlendAdvancedStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAdvancedStateCreateInfoEXT.html)::`blendOverlap` **must** be `VK_BLEND_OVERLAP_UNCORRELATED_EXT`.
- []()`advancedBlendAllOperations` specifies whether all advanced blend operation enums are supported. See the valid usage of [VkPipelineColorBlendAttachmentState](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineColorBlendAttachmentState.html).

## [](#_description)Description

If the `VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT-sType-sType)VUID-VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_BLEND_OPERATION_ADVANCED_PROPERTIES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_blend\_operation\_advanced](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_blend_operation_advanced.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceBlendOperationAdvancedPropertiesEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0