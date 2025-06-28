# VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT(3) Manual Page

## Name

VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT - Structure describing whether rasterization order attachment access can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT` structure is defined as:

```c++
// Provided by VK_EXT_rasterization_order_attachment_access
typedef struct VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           rasterizationOrderColorAttachmentAccess;
    VkBool32           rasterizationOrderDepthAttachmentAccess;
    VkBool32           rasterizationOrderStencilAttachmentAccess;
} VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT;
```

or the equivalent

```c++
// Provided by VK_ARM_rasterization_order_attachment_access
typedef VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesARM;
```

## [](#_members)Members

This structure describes the following features:

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`rasterizationOrderColorAttachmentAccess` indicates that rasterization order access to color and input attachments is supported by the implementation.
- []()`rasterizationOrderDepthAttachmentAccess` indicates that rasterization order access to the depth aspect of depth/stencil and input attachments is supported by the implementation.
- []()`rasterizationOrderStencilAttachmentAccess` indicates that rasterization order access to the stencil aspect of depth/stencil and input attachments is supported by the implementation.

## [](#_description)Description

If the `VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT` structure is included in the `pNext` chain of the [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html) structure passed to [vkGetPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceFeatures2.html), it is filled in to indicate whether each corresponding feature is supported. If the application wishes to use a [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) with any features described by `VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT`, it **must** add an instance of the structure, with the desired feature members set to `VK_TRUE`, to the `pNext` chain of [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html) when creating the [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html).

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT-sType-sType)VUID-VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_RASTERIZATION_ORDER_ATTACHMENT_ACCESS_FEATURES_EXT`

## [](#_see_also)See Also

[VK\_EXT\_rasterization\_order\_attachment\_access](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_rasterization_order_attachment_access.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceRasterizationOrderAttachmentAccessFeaturesEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0